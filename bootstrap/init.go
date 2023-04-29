package bootstrap

import "github.com/gin-gonic/gin"

func initRepo() {}

func initUseCase() {}

func initRoutes(apiGroup *gin.RouterGroup) {

} 

func Init(apiGroup *gin.RouterGroup) {
	initRepo()
	initUseCase()
	initRoutes(apiGroup)
}