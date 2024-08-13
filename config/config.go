package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/nanda03dev/gnosql_client"
	"github.com/nanda03dev/gque/models"
)

var GnoSQLDB *gnosql_client.Database

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}
	if os.Getenv("PORT") != "" {
		os.Setenv("PORT", "8080")
	}

}

func SetupDatabase() {

	host := os.Getenv("GNOSQL_SERVER")

	if host == "" {
		host = "localhost:5455" // default value
	}

	DatabaseName := "gque"

	collections := models.GetAllGnosqlCollections()

	GnoSQLDB = gnosql_client.Connect(host, DatabaseName, true)
	GnoSQLDB.CreateCollections(collections)

	log.Printf("Successfully connected to GNOSQL Database : %v \n", GnoSQLDB.DBName)

}
