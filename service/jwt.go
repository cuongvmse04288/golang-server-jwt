package service

import (
	"crypto/rsa"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang-demo/model"
	"golang-demo/model/request"
	"golang-demo/model/response"
	"io/ioutil"
	"strings"
)

//CustomClaim for JWT
type CustomClaims struct {
	*jwt.StandardClaims
	User string
	Roll string
}

//JWT use RSA256
func GenerateJWT(user request.User, c *gin.Context) (string, error) {
	config, err := GetConfig()
	if err != nil {
		return "", err
	}
	//Check if username/password exist in DB
	isExist, err := VerifyLogin(config, user)

	if isExist == true {

		_, signKey := getRSAKey(config, c)
		t := jwt.New(jwt.GetSigningMethod("RS256"))

		// set our claims
		t.Claims = &CustomClaims{
			&jwt.StandardClaims{
				Issuer: "root",
			},
			user.Username,
			"demo",
		}

		tokenString, err := t.SignedString(signKey)
		if err != nil {
			return "", err
		}
		return tokenString, nil
	}

	return "", err
}

func getRSAKey(config *model.Config, c *gin.Context) (*rsa.PublicKey, *rsa.PrivateKey) {

	signBytes, err := ioutil.ReadFile(config.JWT.PrivateKeyPath)
	if err != nil {
		response.ResponseWithError(500, err, c)
		return nil, nil
	}

	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		response.ResponseWithError(500, err, c)
		return nil, nil
	}

	verifyBytes, err := ioutil.ReadFile(config.JWT.PublicKeyPath)
	if err != nil {
		response.ResponseWithError(500, err, c)
		return nil, nil
	}

	verifyKey, err := jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		response.ResponseWithError(500, err, c)
		return nil, nil
	}

	return verifyKey, signKey
}

func VerifyJWT(tokenString string, c *gin.Context) (string, error) {
	config, err := GetConfig()
	if err != nil {
		return "", err
	}
	verifyKey, _ := getRSAKey(config, c)
	getToken := strings.Split(tokenString, " ")
	token, err := jwt.ParseWithClaims(getToken[1], &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// since we only use the one private key to sign the tokens,
		// we also only use its public counter part to verify
		return verifyKey, nil
	})
	if err != nil {
		return "", err
	}
	claim := token.Claims.(*CustomClaims)
	return claim.User, err
}
