package response

import "github.com/gin-gonic/gin"

type JwtToken struct {
	Token    string
	ExpireIn int
}

func buildResponse(code int, body interface{}, c *gin.Context) {
	c.JSON(code, gin.H{
		"code": code,
		"body": body,
	})
}

func ResponseWithToken(code int, token string, c *gin.Context) {
	buildResponse(code, JwtToken{
		Token:    token,
		ExpireIn: 3600,
	}, c)
}

func ResponseWithError(code int, err error, c *gin.Context) {
	buildResponse(code, err.Error(), c)
}

func SuccessResponse(code int, resp string, c *gin.Context) {
	buildResponse(code, resp, c)
}
