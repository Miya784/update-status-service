package initials

import (
	"database/sql"
	"log"
	"os"
)

var DB *sql.DB

func InitDB() {
	var err error
	//connStr := utils.Strip(os.Getenv("DB_URL"))
	connStr := os.Getenv("DB_URL")
	// Connect to database
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		panic("Failed to connect to  db")
	}
	log.Println("database connect")
}
