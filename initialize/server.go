package initialize

import (
	"github.com/gin-gonic/gin"
	"golang-demo/controller"
)


func Routers() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Recovery())
	authorized := r.Group("/")
	{
		authorized.POST("/login",controller.LoginHandler)
	}
	return r
}

