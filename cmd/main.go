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

// #[.'.]:> Entry point of the application
func main() {
	log.Println("['.']:>")
	log.Println("['.']:>--------------------------------------------")
	log.Println("['.']:>--- <starting markitos-svc-boilerplates>  ---")

	//#[.'.]:> Load configuration
	config = loadConfiguration()
	log.Println("['.']:>------- configuration loaded")

	//#[.'.]:> Initialize database
	loadDatabase()
	log.Println("['.']:>------- database initialized")

	//#[.'.]:> Start servers (REST and gRPC)
	startServers()
	log.Println("['.']:>--------------------------------------------")
	log.Println("['.']:>--- <markitos-svc-boilerplates stopped>  ---")
	log.Println("['.']:>")
}

// #[.'.]:> Load application configuration
func loadConfiguration() configuration.BoilerplateConfiguration {
	config, err := configuration.LoadConfiguration(".")
	if err != nil {
		log.Fatal("['.']:>------- unable to load configuration: ", err)
	}
	return config
}

// #[.'.]:> Initialize and migrate the database
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

// #[.'.]:> Start both REST and gRPC servers concurrently
func startServers() {
	// Create a context for graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Channel to capture OS signals
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// WaitGroup to wait for both servers to finish
	var wg sync.WaitGroup
	wg.Add(2)

	// Start REST server
	go func() {
		defer wg.Done()
		if err := runRESTServer(ctx); err != nil && err != http.ErrServerClosed {
			log.Printf("['.']:> error running REST server: %v", err)
		}
	}()

	// Start gRPC server
	go func() {
		defer wg.Done()
		if err := runGRPCServer(ctx); err != nil {
			log.Printf("['.']:> error running gRPC server: %v", err)
		}
	}()

	// Wait for termination signal
	<-stop
	log.Println("['.']:>------- shutting down servers gracefully...")

	// Cancel the context to signal shutdown
	cancel()

	// Wait for both servers to finish
	wg.Wait()
}

// #[.'.]:> Run the REST server
func runRESTServer(ctx context.Context) error {
	apiServer := api.NewServer(config.HTTPServerAddress, repository)
	server := &http.Server{
		Addr:    config.HTTPServerAddress,
		Handler: apiServer.Router(),
	}

	// Listen for context cancellation to shut down the server
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

// #[.'.]:> Run the gRPC server
func runGRPCServer(ctx context.Context) error {
	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	server := gapi.NewServer(config.HTTPServerAddress, repository)
	gapi.RegisterBoilerplateServiceServer(grpcServer, server)
	reflection.Register(grpcServer)

	// Listen for context cancellation to gracefully stop the server
	go func() {
		<-ctx.Done()
		log.Println("['.']:> shutting down gRPC server...")
		grpcServer.GracefulStop()
	}()

	log.Printf("['.']:> gRPC server running at %s", config.GRPCServerAddress)
	return grpcServer.Serve(listener)
}
