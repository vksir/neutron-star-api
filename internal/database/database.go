package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

var DB *sql.DB

func InitDataBase() {

}

func Connect() {
	var err error
	DB, err = sql.Open("postgres", "postgresql://ns:vksir97634@localhost:5432/ns_api")
	if err != nil {
		log.Fatal("Unable to use data source name", err)
	}
	DB.SetConnMaxLifetime(0)
	DB.SetMaxIdleConns(3)
	DB.SetMaxOpenConns(3)
	Ping()
}

func Close() error {
	return DB.Close()
}

func Ping() {
	if err := DB.Ping(); err != nil {
		log.Fatalf("unable to connect to database: %v", err)
	}
}
