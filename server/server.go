package server

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	return r
}

func Start(router *gin.Engine, port string) error {
	return router.Run(port)
}