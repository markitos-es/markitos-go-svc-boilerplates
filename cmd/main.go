package main

import (
	"log"

	"github.com/markitos-es/markitos-svc-boilerplates/infrastructure/api"
	"github.com/markitos-es/markitos-svc-boilerplates/infrastructure/configuration"
	"github.com/markitos-es/markitos-svc-boilerplates/infrastructure/database"
	"github.com/markitos-es/markitos-svc-boilerplates/internal/domain"
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

	repository = loadDatabase()
	log.Println("['.']:>------- database ok")

	server := createServer()
	log.Println("['.']:>------- server created and running at: ", config.ServerAddress)
	log.Println("['.']:>--------------------------------------------")
	if err := server.Run(); err != nil {
		log.Fatal("['.']:>------- unable to start server: ", err)
	}
}

func loadConfiguration() configuration.BoilerplateConfiguration {
	config, err := configuration.LoadConfiguration(".")
	if err != nil {
		log.Fatal("['.']:>------- unable to load configuration: ", err)
	}

	return config
}

func createServer() *api.Server {
	return api.NewServer(config.ServerAddress, repository)
}

func loadDatabase() domain.BoilerplateRepository {
	db, err := gorm.Open(postgres.Open(config.DatabaseDsn), &gorm.Config{})
	if err != nil {
		log.Fatal("['.']:> error unable to connect to database:", err)
	}
	err = db.AutoMigrate(&domain.Boilerplate{})
	if err != nil {
		log.Fatal("['.']:> error unable to migrate database:", err)
	}

	repo := database.NewBoilerPostgresRepository(db)

	return &repo
}
