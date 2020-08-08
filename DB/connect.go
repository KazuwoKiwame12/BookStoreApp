package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

//Connect ...DB型を返す
func Connect() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}

	connStr := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@/" + os.Getenv("DB_NAME")
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal(err.Error())
	}
	return db
}
