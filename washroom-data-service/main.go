package main

import (
	"database/sql"
	"io/ioutil"
	"log"
	"os"

	"washroom-data-service/handler"
	"washroom-data-service/middleware"
	"washroom-data-service/repository/sqlite"
	"washroom-data-service/service"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	if _, err := os.Stat("./washrooms.db"); os.IsNotExist(err) {
		// Create and initialize database
		file, err := os.Create("./washrooms.db")
		if err != nil {
			log.Fatal(err)
		}
		file.Close()

		// Open DB and run schema
		db, err := sql.Open("sqlite3", "./washrooms.db")
		if err != nil {
			log.Fatal(err)
		}
		schemaBytes, err := ioutil.ReadFile("washroom-data-service/repository/sqlite/schema.sql")
		if err != nil {
			log.Fatal(err)
		}
		if _, err := db.Exec(string(schemaBytes)); err != nil {
			log.Fatal(err)
		}
		db.Close()
	}

	db, err := sql.Open("sqlite3", "./washrooms.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Initialize repositories
	washroomRepo := sqlite.NewSQLiteRepository(db)
	locationQuery := sqlite.NewSQLiteLocationQuery(db)
	eventStore := sqlite.NewSQLiteEventStore(db)

	// Initialize services
	washroomService := service.NewWashroomService(washroomRepo, locationQuery, eventStore, nil)

	// Initialize handlers
	washroomHandler := handler.NewWashroomHandler(washroomService)

	// Set up Gin router
	r := gin.Default()
	r.Use(middleware.AuthMiddleware())

	// API routes
	api := r.Group("/api/v1")
	{
		washrooms := api.Group("/washrooms")
		{
			washrooms.POST("/", washroomHandler.Create)
			washrooms.GET("/:id", washroomHandler.GetByID)
			washrooms.GET("/nearby", washroomHandler.FindNearby)
			// Add more routes as needed
		}
	}

	// Start server
	log.Println("API server listening")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
