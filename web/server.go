package main

import (
	"context"
	"github.com/swaggo/echo-swagger"
	"log"
	"os"
	"postgres-go-echo-htmx-bulma/internal/handler"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	_ "postgres-go-echo-htmx-bulma/docs"
)

//	@title			Super API
//	@version		1.0
//	@description	This is a super API. Don't think otherwise

//	@contact.name	Pedro Flores
//	@contact.url	http://codelab.com.py
//	@contact.email	hola@codelab.com.py

// @license.name	Apache 2.0
// @license.url	http://www.apache.org/licenses/LICENSE-2.0.html
func main() {

	// Get the database connection string from environment variable or configuration file
	connectionString := os.Getenv("DATABASE_URI")
	if connectionString == "" {
		log.Fatal("Database URI is not set")
	}

	// Initialize the configuration for the connection pool
	config, err := pgxpool.ParseConfig(connectionString)
	if err != nil {
		log.Fatalf("Error parsing connection string: %v", err)
	}

	// Initialize the connection pool
	dbpool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("Error establishing connection pool: %v", err)
	}
	defer dbpool.Close()

	// Check if the connection is successful
	if err := dbpool.Ping(context.Background()); err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	log.Println("Database connection established")

	// Create a new Echo instance
	e := echo.New()

	// Initialize handlers and pass dbpool to them
	initHandlers(e, dbpool)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Start the server
	e.Logger.Fatal(e.Start(":8080"))
}

// Initialize handlers and pass dbpool to them
func initHandlers(e *echo.Echo, dbpool *pgxpool.Pool) {
	// Create handlers and pass dbpool to them
	e.GET("/heroes", handler.ListHeroHandler(dbpool))
	e.POST("/heroes", handler.CreateHeroHandler(dbpool))
	e.GET("/heroes/:id", handler.GetHeroHandler(dbpool))
	e.PUT("/heroes/:id", handler.UpdateHeroHandler(dbpool))
	e.DELETE("/heroes/:id", handler.DeleteHeroHandler(dbpool))
}
