package main

import (
	"context"
	"log"
	"os"
	"postgres-go-echo-htmx-bulma/internal/handler"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
)

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

	// Define routes
  e.GET("/heroes", func(c echo.Context) error {
      return handler.ListHeroHandler(dbpool)(c)
  })
  e.POST("/heroes", func(c echo.Context) error {
      return handler.CreateHeroHandler(dbpool)(c)
  })
  e.GET("/heroes/:id", func(c echo.Context) error {
      return handler.GetHeroHandler(dbpool)(c)
  })
  e.PUT("/heroes/:id", func(c echo.Context) error {
      return handler.UpdateHeroHandler(dbpool)(c)
  })
  e.DELETE("/heroes/:id", func(c echo.Context) error {
      return handler.DeleteHeroHandler(dbpool)(c)
  })

	// Start the server
	e.Logger.Fatal(e.Start(":8080"))
}
