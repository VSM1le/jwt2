package database

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Database struct {
	*sqlx.DB
}

func DBinstance() (*Database, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("erro loading .env file.")
	}
	portDB := os.Getenv("DB_URL")
	if portDB == "" {
		log.Fatal("error getting portDb.")
	}
	db, err := sqlx.Connect("postgres", portDB)
	if err != nil {
		log.Fatalln("error connecting to database:", err)
	}
	return &Database{db}, nil
}

func (d *Database) Close() error {
	return d.DB.Close()
}
