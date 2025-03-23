// #[.'.]:> Main package for starting the application
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

	"github.com/markitos-es/markitos-svc-boilerplates/infrastructure/api"
	"github.com/markitos-es/markitos-svc-boilerplates/infrastructure/configuration"
	"github.com/markitos-es/markitos-svc-boilerplates/infrastructure/database"
	"github.com/markitos-es/markitos-svc-boilerplates/infrastructure/gapi"
	"github.com/markitos-es/markitos-svc-boilerplates/internal/domain"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var repository domain.BoilerplateRepository
var config configuration.BoilerplateConfiguration

func main() {
	log.Println("['.']:>")
	log.Println("['.']:>--------------------------------------------")
	log.Println("['.']:>--- <starting markitos-svc-boilerplates>  ---")

	loadConfiguration()
	log.Println("['.']:>------- configuration loaded")

	loadDatabase()
	log.Println("['.']:>------- database initialized")

	startServers()
	log.Println("['.']:>--------------------------------------------")
	log.Println("['.']:>--- <markitos-svc-boilerplates stopped>  ---")
	log.Println("['.']:>")
}

func loadConfiguration() {
	loadedConfig, err := configuration.LoadConfiguration(".")
	if err != nil {
		log.Fatal("['.']:>------- unable to load configuration: ", err)
	}

	config = loadedConfig
}

func loadDatabase() {
	db, err := gorm.Open(postgres.Open(config.DatabaseDsn), &gorm.Config{})
	if err != nil {
		log.Fatal("['.']:> error unable to connect to database:", err)
	}
	err = db.AutoMigrate(&domain.Boilerplate{})
	if err != nil {
		log.Fatal("['.']:> error unable to migrate database:", err)
	}

	repo := database.NewBoilerPostgresRepository(db)
	repository = &repo
}

func startServers() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		if err := runRESTServer(ctx); err != nil && err != http.ErrServerClosed {
			log.Printf("['.']:> error running REST server: %v", err)
		}
	}()

	go func() {
		defer wg.Done()
		if err := runGRPCServer(ctx); err != nil {
			log.Printf("['.']:> error running gRPC server: %v", err)
		}
	}()
	<-stop
	log.Println("['.']:>------- shutting down servers gracefully...")
	cancel()
	wg.Wait()
}

func runRESTServer(ctx context.Context) error {
	apiServer := api.NewServer(config.HTTPServerAddress, repository)
	server := &http.Server{
		Addr:    config.HTTPServerAddress,
		Handler: apiServer.Router(),
	}

	go func() {
		<-ctx.Done()
		log.Println("['.']:> shutting down REST server...")
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := server.Shutdown(shutdownCtx); err != nil {
			log.Printf("['.']:> error shutting down REST server: %v", err)
		}
	}()

	log.Printf("['.']:> REST server running at %s", config.HTTPServerAddress)
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
	server := gapi.NewServer(config.HTTPServerAddress, repository)

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
