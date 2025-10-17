package collector

import (
	"github.com/gin-gonic/gin"
)

func SetupAPI() *gin.Engine {
	router := gin.Default()
	router.Use(ErrorHandler())
	router.POST("/getlinks", Posthandler)
	return router
}
