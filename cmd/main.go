package main

import (
	"log"
	"net"

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

	config = loadConfiguration()
	log.Println("['.']:>------- configuration loaded")

	loadDatabase()
	log.Println("['.']:>------- database ok")

	// log.Println("['.']:>------- server REST created and running at: ", config.ServerAddress)
	// runRESTServer()

	log.Println("['.']:>------- server GRPC created and running at: ", config.GRPCServerAddress)
	runGRPCServer()
	log.Println("['.']:>--------------------------------------------")
	log.Println("['.']:>--- <markitos-svc-boilerplates started>  ---")
	log.Println("['.']:>")
}

func loadConfiguration() configuration.BoilerplateConfiguration {
	config, err := configuration.LoadConfiguration(".")
	if err != nil {
		log.Fatal("['.']:>------- unable to load configuration: ", err)
	}

	return config
}

func runGRPCServer() {
	grpcServer := grpc.NewServer()
	server := gapi.NewServer(config.ServerAddress, repository)

	gapi.RegisterBoilerplateServiceServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal("['.']:> error unable to listen to GRPC server:", err)
	}

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("['.']:> error unable to serve GRPC server:", err)
	}
}

func runRESTServer() {
	server := api.NewServer(config.ServerAddress, repository)
	if err := server.Run(); err != nil {
		log.Fatal("['.']:> error unable to start server:", err)
	}
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
