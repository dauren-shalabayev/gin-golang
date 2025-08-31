package main

import (
	"database/sql"
	"log"
	"rest-api-in-gin/cmd/internal/database"
	"rest-api-in-gin/cmd/internal/env"

	"rest-api-in-gin/docs"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"
)

// @title Rest API in GIN
// @version 1.0
// @description Rest API in GIN
// @host localhost:8080
// @BasePath /api/v1
// @schemes http
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @description Enter the token with the `Bearer ` prefix, e.g. `Bearer abcde12345`.

type application struct {
	port      int
	jwtSecret string
	models    database.Models
}

func main() {
	// Initialize Swagger docs
	docs.SwaggerInfo.Title = "Rest API in GIN"
	docs.SwaggerInfo.Description = "Rest API in GIN"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http"}

	db, err := sql.Open("sqlite3", "./cmd/migrate/data.db")
	if err != nil {
		log.Fatal(err)
	}

	// Test the database connection
	if err = db.Ping(); err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	defer db.Close()
	models := database.NewModels(db)
	app := &application{
		port:      env.GetEnvInt("PORT", 8080),
		jwtSecret: env.GetEnvString("JWT_SECRET", "secret"),
		models:    models,
	}
	if err := app.serve(); err != nil {
		log.Fatal(err)
	}
}
