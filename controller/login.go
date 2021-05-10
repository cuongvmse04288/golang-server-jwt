package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang-demo/model/request"
	"golang-demo/service"
	"log"
)

func LoginHandler(c *gin.Context){
	var l request.User
	err := c.ShouldBindJSON(&l)
	if err != nil {
		log.Print(err)
	}
	fmt.Println()
	token,err := service.GenerateJWT(l)
	if err != nil {
		c.JSON(500,gin.H{
			"error":err.Error(),
		})
		return
	}
	c.JSON(200,gin.H{
		"jwt":token,
		"expiredIn":3600,
	})
}
