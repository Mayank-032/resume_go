package routes

import (
	"log"
	"net/http"
	upsertResumeRequest "resume-go/domain/dto/requests"
	"resume-go/domain/dto/responses"
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

func upsertResume(c *gin.Context) {
	response := responses.BaseResponse{}
	request := upsertResumeRequest.UpsertResume{}
	err := c.ShouldBindJSON(&request);
	if err == nil {
		err = Validate(request)
	}
	if err != nil {
		log.Printf("Error %v\n, invalid request made\n\n", err.Error())
		c.JSON(http.StatusBadRequest, response.Response(false, "invalid request made"))
		return
	}

	resumeData := ToResumeDto(request)
	userId, err := ResumeUCase.UpsertPDF(c, resumeData)
	if err != nil {
		log.Printf("Error %v\n, unable to upsert resume\n\n", err.Error())
		c.JSON(http.StatusBadRequest, response.Response(false, "unable to upsert resume"))
		return
	}
	c.JSON(http.StatusOK, response.Response(true, userId))
} 

func removeResume(c *gin.Context) {} 