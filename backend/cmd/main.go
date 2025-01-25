package main

import (
	"log"
	"github.com/Cleaach/threads/backend/cmd/api"
	"github.com/Cleaach/threads/backend/db"
	"github.com/Cleaach/threads/backend/config"
	"github.com/go-sql-driver/mysql"
	"database/sql"
)

func main() {
	
	// Connect to database
	db, err := db.NewMySQLStorage(mysql.Config{
		User: config.Envs.DBUser,
		Passwd: config.Envs.DBPassword,
		Addr: config.Envs.DBAddress,
		DBName: config.Envs.DBName,
		Net: "tcp",
		AllowNativePasswords: true,
		ParseTime: true,
	})

	// Error check
	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)

	// Initialize API Server
	server := api.NewAPIServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

// Check if database is online
func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB: Successfully connected!")
}