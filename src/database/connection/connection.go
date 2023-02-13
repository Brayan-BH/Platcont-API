package connection

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func Connection() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error configuración de variables de entorno")
	}
	server := os.Getenv("ENV_DDBB_SERVER")
	user := os.Getenv("ENV_DDBB_USER")
	password := os.Getenv("ENV_DDBB_PASSWORD")
	port := os.Getenv("ENV_DDBB_PORT")
	database := os.Getenv("ENV_DDBB_DATABASE")
	// fmt.Println("database: ", database)
	connection_string := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", server, port, user, password, database)
	db, err := sql.Open("postgres", connection_string)
	if err != nil {
		log.Fatal("Error connection: ", err.Error())
	}
	err = db.Ping()
	if err != nil {
		db.Close()
		log.Fatal("Error creating connection: ", err.Error())
	}
	return db
}

func ConnectionCloud() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error configuración de variables de entorno")
	}
	server := os.Getenv("ENV_DDBB_SERVER")
	user := os.Getenv("ENV_DDBB_USER")
	password := os.Getenv("ENV_DDBB_PASSWORD")
	database := os.Getenv("ENV_DDBB_DATABASE")
	port := os.Getenv("ENV_DDBB_PORT")

	connection_string := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", server, port, user, password, database)

	db, err := sql.Open("postgres", connection_string)
	if err != nil {
		log.Fatal("Error connection: ", err.Error())
	}
	err = db.Ping()
	if err != nil {
		db.Close()
		log.Fatal("Error creating connection: ", err.Error())
	}
	return db
}
