package main

import (
	_ "HL_project_management/docs"
	"HL_project_management/internal/repository"
	"HL_project_management/internal/router"
	"flag"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)

func main() {
	var cfg repository.Config
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file	")
	}
	url := os.Getenv("url")
	flag.IntVar(&cfg.Port, "port", 8080, "API server port")
	flag.StringVar(&cfg.Env, "env", "development", "Environment (development|staging|production)")
	flag.StringVar(&cfg.Db.Dsn, "db-dsn", url, "PostgreSQL DSN")
	flag.Parse()
	_, err = repository.OpenDB(cfg)
	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	} else {
		log.Println("Connected to the database")
	}
	defer repository.CloseDB()
	r := router.SetupRouter()

	// Add CORS support
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	handler := c.Handler(r)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}
