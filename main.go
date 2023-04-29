package main

import (
	"log"
	"os"
	"resume-go/bootstrap"
	"resume-go/db"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	connectionURI := os.Getenv("MONGODB_URI")
	db.ConnectDB(connectionURI)

	r := gin.Default()

	apiGroup := r.Group("/api")
	bootstrap.Init(apiGroup)
}