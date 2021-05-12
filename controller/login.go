package controller

import (
	"github.com/gin-gonic/gin"
	"golang-demo/model/request"
	"golang-demo/model/response"
	"golang-demo/service"
)

func LoginHandler(c *gin.Context) {
	var user request.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		response.ResponseWithError(500, err, c)
		return
	}

	token, err := service.GenerateJWT(user, c)
	if err != nil {
		response.ResponseWithError(500, err, c)
		return
	}
	response.ResponseWithToken(200, token, c)
}
