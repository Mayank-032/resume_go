package main

import (
	"resume-go/bootstrap"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	apiGroup := r.Group("/api")
	bootstrap.Init(apiGroup)
}