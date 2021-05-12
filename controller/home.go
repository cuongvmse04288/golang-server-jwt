package controller

import (
	"github.com/gin-gonic/gin"
	"golang-demo/model/response"
	"golang-demo/service"
)

func HomeHandler(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")

	userName, err := service.VerifyJWT(token, c)
	if err != nil {
		response.ResponseWithError(500, err, c)
		return
	}
	if userName != "" {
		response.SuccessResponse(200, "Hello "+userName, c)
	}
}
