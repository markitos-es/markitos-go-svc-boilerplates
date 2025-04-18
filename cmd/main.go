// #[.'.]:> Paquete principal que inicia la aplicación
package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"markitos-svc-boilerplates/infrastructure/api"
	"markitos-svc-boilerplates/infrastructure/configuration"
	"markitos-svc-boilerplates/infrastructure/database"
	"markitos-svc-boilerplates/infrastructure/gapi"
	"markitos-svc-boilerplates/internal/domain"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var repository domain.BoilerplateRepository
var config configuration.BoilerplateConfiguration

// #[.'.]:> Función principal que orquesta el inicio y cierre controlado de la aplicación
func main() {
	//#[.'.]:> PASO 1: Mostrar banner de inicio
	//#[.'.]:> Estos logs nos ayudan a identificar claramente el inicio del servicio
	log.Println("['.']:>")
	log.Println("['.']:>--------------------------------------------")
	log.Println("['.']:>--- <starting markitos-svc-boilerplates>  ---")

	//#[.'.]:> PASO 2: Cargar la configuración desde archivos o variables de entorno
	//#[.'.]:> Esta función se encarga de establecer todos los parámetros operativos
	loadConfiguration()
	log.Println("['.']:>------- configuration loaded")

	//#[.'.]:> PASO 3: Inicializar la conexión a la base de datos y el repositorio
	//#[.'.]:> Preparamos el acceso a datos y la estructura de tablas
	loadDatabase()
	log.Println("['.']:>------- database initialized")

	//#[.'.]:> PASO 4: Iniciar los servidores HTTP y gRPC
	//#[.'.]:> Ponemos en marcha los puntos de entrada para clientes REST y gRPC
	startServers()

	//#[.'.]:> PASO 5: Mostrar banner de finalización al terminar
	//#[.'.]:> Estos logs marcan claramente el fin de la ejecución del servicio
	log.Println("['.']:>--------------------------------------------")
	log.Println("['.']:>--- <markitos-svc-boilerplates stopped>  ---")
	log.Println("['.']:>")
}

// #[.'.]:> Esta función carga la configuración del servicio
func loadConfiguration() {
	//#[.'.]:> PASO 1: Intentar cargar la configuración desde archivo o variables de entorno
	//#[.'.]:> Busca "app.env" en el directorio actual, o usa variables de entorno si no existe
	loadedConfig, err := configuration.LoadConfiguration(".")
	if err != nil {
		//#[.'.]:> Si hay error, terminar la aplicación inmediatamente
		//#[.'.]:> No podemos operar sin configuración válida
		log.Fatal("['.']:>------- unable to load configuration: ", err)
	}

	//#[.'.]:> PASO 2: Almacenar la configuración en una variable global
	//#[.'.]:> Esto la hace accesible al resto de funciones del programa
	config = loadedConfig
}

// #[.'.]:> Esta función inicializa la base de datos y el repositorio
func loadDatabase() {
	//#[.'.]:> PASO 1: Establecer conexión con PostgreSQL usando la cadena de conexión
	//#[.'.]:> GORM abstrae los detalles de la conexión y manejo de la base de datos
	db, err := gorm.Open(postgres.Open(config.DatabaseDsn), &gorm.Config{})
	if err != nil {
		//#[.'.]:> Si no podemos conectar a la base de datos, es un error fatal
		log.Fatal("['.']:> error unable to connect to database:", err)
	}

	//#[.'.]:> PASO 2: Ejecutar migraciones automáticas para crear o actualizar tablas
	//#[.'.]:> Esto asegura que la estructura de la base de datos coincida con nuestros modelos
	err = db.AutoMigrate(&domain.Boilerplate{})
	if err != nil {
		//#[.'.]:> Si las migraciones fallan, no podemos continuar
		log.Fatal("['.']:> error unable to migrate database:", err)
	}

	//#[.'.]:> PASO 3: Crear una instancia del repositorio con la conexión a la base de datos
	//#[.'.]:> El repositorio encapsula toda la lógica de acceso a datos
	repo := database.NewBoilerplatePostgresRepository(db)
	repository = &repo
}

// #[.'.]:> Esta función inicia los servidores y maneja su ciclo de vida
func startServers() {
	//#[.'.]:> PASO 1: Crear un contexto cancelable para señalizar el apagado
	//#[.'.]:> Este contexto se propagará a los servidores para gestionar su ciclo de vida
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	//#[.'.]:> PASO 2: Configurar un canal para capturar señales del sistema operativo
	//#[.'.]:> Esto permite detectar Ctrl+C o señales de apagado del sistema
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	//#[.'.]:> PASO 3: Crear un grupo de espera para coordinar el apagado de servidores
	//#[.'.]:> El WaitGroup nos permite esperar a que ambos servidores se detengan completamente
	var wg sync.WaitGroup
	wg.Add(2)

	//#[.'.]:> PASO 4: Iniciar el servidor REST en una goroutine independiente
	//#[.'.]:> Ejecutamos en paralelo para no bloquear la aplicación principal
	go func() {
		defer wg.Done()
		if err := runRESTServer(ctx); err != nil && err != http.ErrServerClosed {
			log.Printf("['.']:> error running REST server: %v", err)
		}
	}()

	//#[.'.]:> PASO 5: Iniciar el servidor gRPC en otra goroutine independiente
	//#[.'.]:> También en paralelo con el servidor REST y el hilo principal
	go func() {
		defer wg.Done()
		if err := runGRPCServer(ctx); err != nil {
			log.Printf("['.']:> error running gRPC server: %v", err)
		}
	}()

	//#[.'.]:> PASO 6: Bloquear hasta recibir una señal de terminación
	//#[.'.]:> La aplicación esperará aquí hasta que se reciba Ctrl+C o SIGTERM
	<-stop
	log.Println("['.']:>------- shutting down servers gracefully...")

	//#[.'.]:> PASO 7: Cancelar el contexto para iniciar el apagado controlado
	//#[.'.]:> Esto enviará la señal de terminación a ambos servidores
	cancel()

	//#[.'.]:> PASO 8: Esperar a que ambos servidores terminen completamente
	//#[.'.]:> No saldremos hasta que ambos servidores hayan completado su apagado
	wg.Wait()
}

// #[.'.]:> Esta función inicia y maneja el ciclo de vida del servidor REST/HTTP
func runRESTServer(ctx context.Context) error {
	//#[.'.]:> PASO 1: Crear una nueva instancia del servidor API
	//#[.'.]:> Este objeto configura todas las rutas y controladores HTTP
	apiServer := api.NewServer(config.HTTPServerAddress, repository)

	//#[.'.]:> PASO 2: Configurar el servidor HTTP estándar de Go
	//#[.'.]:> Usamos la implementación de servidor HTTP del paquete estándar
	server := &http.Server{
		Addr:    config.HTTPServerAddress,
		Handler: apiServer.Router(),
	}

	//#[.'.]:> PASO 3: Configurar el apagado controlado (graceful shutdown)
	//#[.'.]:> Esta goroutine se ejecuta en segundo plano y espera la señal de apagado
	go func() {
		//#[.'.]:> Esperar a que el contexto se cancele (señal de apagado)
		<-ctx.Done()
		log.Println("['.']:> shutting down REST server...")

		//#[.'.]:> Crear un nuevo contexto con timeout para el apagado
		//#[.'.]:> Si el apagado tarda más de 5 segundos, se forzará
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		//#[.'.]:> Iniciar el proceso de apagado controlado
		//#[.'.]:> Esto cierra conexiones existentes de forma ordenada
		if err := server.Shutdown(shutdownCtx); err != nil {
			log.Printf("['.']:> error shutting down REST server: %v", err)
		}
	}()

	//#[.'.]:> PASO 4: Registrar que el servidor está en funcionamiento
	log.Printf("['.']:> REST server running at %s", config.HTTPServerAddress)

	//#[.'.]:> PASO 5: Iniciar el servidor (este método bloquea hasta que ocurra un error)
	//#[.'.]:> El servidor ahora escucha activamente las peticiones entrantes
	return server.ListenAndServe()
}

// #[.'.]:> Esta función inicia y maneja el ciclo de vida del servidor gRPC
func runGRPCServer(ctx context.Context) error {
	//#[.'.]:> PASO 1: Crear un "oído" (listener) en la red
	//#[.'.]:> Este listener escuchará peticiones TCP en la dirección y puerto configurados
	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		return err
	}

	//#[.'.]:> PASO 2: Crear una nueva instancia del servidor gRPC
	//#[.'.]:> Este objeto es el corazón del servidor y manejará todas las peticiones
	grpcServer := grpc.NewServer()

	//#[.'.]:> PASO 3: Crear la implementación de nuestro servicio
	//#[.'.]:> Esta es la parte que contiene la lógica de negocio real
	server := gapi.NewServer(config.GRPCServerAddress, repository)

	//#[.'.]:> PASO 4: Registrar nuestro servicio en el servidor gRPC
	//#[.'.]:> Esto conecta nuestras implementaciones con el sistema gRPC
	gapi.RegisterBoilerplateServiceServer(grpcServer, server)

	//#[.'.]:> PASO 5: Habilitar la reflexión para facilitar pruebas
	//#[.'.]:> La reflexión permite que herramientas como grpcurl descubran nuestros servicios
	reflection.Register(grpcServer)

	//#[.'.]:> PASO 6: Configurar el apagado controlado (graceful shutdown)
	//#[.'.]:> Esta goroutine se ejecuta en segundo plano y espera la señal de apagado
	go func() {
		//#[.'.]:> Bloquea hasta que el contexto se cancele (señal de apagado)
		<-ctx.Done()

		//#[.'.]:> Registra un mensaje indicando que se está apagando el servidor
		log.Println("['.']:> shutting down gRPC server...")

		//#[.'.]:> Realiza un apagado controlado:
		//#[.'.]:> - Deja de aceptar nuevas conexiones
		//#[.'.]:> - Espera a que terminen las llamadas en curso
		//#[.'.]:> - Cierra todas las conexiones limpiamente
		grpcServer.GracefulStop()
	}()

	//#[.'.]:> PASO 7: Registrar que el servidor está en funcionamiento
	log.Printf("['.']:> gRPC server running at %s", config.GRPCServerAddress)

	//#[.'.]:> PASO 8: Iniciar el servidor (este método bloquea hasta que ocurra un error)
	//#[.'.]:> El servidor ahora escucha activamente las peticiones entrantes
	return grpcServer.Serve(listener)
}
