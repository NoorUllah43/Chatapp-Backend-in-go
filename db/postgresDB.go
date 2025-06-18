package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2/log"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectPostgresql() {

	connectionString := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	fmt.Print(connectionString)

	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		log.Errorf("database not connecting error: ", err)
	}

	DB = db
	log.Info("database connected to postgresql")

}
