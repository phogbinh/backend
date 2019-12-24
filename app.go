package main

import (
	"database/sql"
	"log"
	"os"

	db "backend/database"
	DUTU "backend/database_users_table_util"
	"backend/router"

	"github.com/gin-gonic/gin"
)

func getDatabaseHandler() *sql.DB {
	databasePtr, openDatabaseError := sql.Open("mysql", os.Getenv("DATABASE_URL"))
	if openDatabaseError != nil {
		log.Fatalf("Error opening database: %q.", openDatabaseError)
	}
	return databasePtr
}

func run() {
	// databasePtr := getDatabaseHandler()
	// defer databasePtr.Close()
	db.Init()
	databasePtr := db.SqlDb
	createDatabaseUsersTableError := DUTU.CreateTableIfNotExists(databasePtr)
	if createDatabaseUsersTableError != nil {
		log.Fatalf("Error creating database table: %q.", createDatabaseUsersTableError)
	}
	db.CreateDatabases(db.SqlDb)

	var httpServer *gin.Engine

	httpServer = gin.Default()

	router.Register(httpServer, databasePtr)

	serverAddr := "127.0.0.1:8080"

	// listen and serve on 0.0.0.0:8080
	httpServer.Run(serverAddr)
}

func main() {
	run()
}
