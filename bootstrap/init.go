package bootstrap

import (
	DB "resume-go/db"
	"resume-go/domain/interfaces/repository"
	"resume-go/domain/interfaces/usecase"
	_resumeRepo "resume-go/pkg/resume/repository"
	_resumeUCase "resume-go/pkg/resume/usecase"
	_resumeRoutes "resume-go/pkg/resume/routes"

	"github.com/gin-gonic/gin"
)


var (
	resumeUCase usecase.IResumeUCase
	resumeRepo repository.IResumeRepo
)


func initRepo() {
	resumeRepo = _resumeRepo.NewResumeRepository(DB.ResumeDataCollection)
}

func initUseCase() {
	resumeUCase = _resumeUCase.NewResumeUseCase(resumeRepo)
}

func initRoutes(apiGroup *gin.RouterGroup) {
	_resumeRoutes.Init(apiGroup, resumeUCase)
} 

func Init(apiGroup *gin.RouterGroup) {
	initRepo()
	initUseCase()
	initRoutes(apiGroup)
}