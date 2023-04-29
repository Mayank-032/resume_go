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
		log.Printf("Error %v, unable to load .env file", err)
		return
	}

	connectionURI := os.Getenv("MONGODB_URI")
	db.ConnectDB(connectionURI)

	credsFilePath := os.Getenv("CloudCredentialFilePath")
	db.ConnectCloud(credsFilePath)

	r := gin.Default()

	apiGroup := r.Group("/api")
	bootstrap.Init(apiGroup)

	r.Run(":8001")
}
