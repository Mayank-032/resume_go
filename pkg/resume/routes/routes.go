package routes

import (
	"resume-go/domain/interfaces/usecase"

	"github.com/gin-gonic/gin"
)

var ResumeUCase usecase.IResumeUCase
func Init(apiGroup *gin.RouterGroup, resumeUCase usecase.IResumeUCase) {
	ResumeUCase = resumeUCase

	apiGroup.GET("/resume/{user_id}", fetchResume)
	apiGroup.PUT("/resume", upsertResume)
	apiGroup.PATCH("/resume/{user_id}", removeResume)
}

func fetchResume(c *gin.Context) {} 

func upsertResume(c *gin.Context) {} 

func removeResume(c *gin.Context) {} 