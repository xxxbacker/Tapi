package internal

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func GenPort() string {
	er := godotenv.Load(".env")
	if er != nil {
		log.Fatal(".env file not found")
	}
	port, err := os.LookupEnv("PORT")
	if err != true {
		log.Fatal(" var env not found")
	}
	return port
}

func GenConStr() string {
	er := godotenv.Load(".env")
	if er != nil {
		log.Fatal(".env file not found")
	}
	// Подключиться к базе данных
	connStr, err := os.LookupEnv("Con_Str")
	if err != true {
		log.Fatal(" var env not found")
	}
	return connStr
}
