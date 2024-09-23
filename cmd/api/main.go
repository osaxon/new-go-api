package main

import (
	"context"
	_ "embed"
	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
	"log"
	database "new-go-api/internal/database"
	"new-go-api/internal/server"
)

func main() {
	ctx := context.Background()

	// Initialize the database
	db, err := database.NewDB(ctx)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	queries := database.New(db)

	// Start the app
	app := server.NewApp(queries)
	app.Start(":1323")
}
