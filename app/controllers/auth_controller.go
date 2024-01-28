package controllers

import (
	"github.com/gin-gonic/gin"

	"log"
	"net/http"
	"strings"

	"models/user"
	"controllers/basicauth"
)

func Signup(c *gin.Context) {
	user := usermodel.Create(c)
	c.JSON(200, user)
}

func Signin(c *gin.Context) {
	token, err_message := usermodel.Signin(c)
	if err_message != "" {
		log.Println("ログインできませんでした")
		c.JSON(http.StatusBadRequest, gin.H{"error": err_message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"access-token": token})
	log.Println("ログインできました")
	c.Redirect(302, "/")
}

func Session(c *gin.Context) {
	encodedToken := c.Request.Header.Get("access-token")
  // encoded-tokenをdecode
  decodedToken, err := basicauth.DecodeBase64(encodedToken)

	// decoded-tokenの分割
	ok, tokenUsername, tokenPassword := splitToken(decodedToken)

	if !ok {
		c.JSON(401, gin.H{"err": "分割失敗！"})
    return
  }

	user, err_message := usermodel.Session(tokenUsername, tokenPassword)

  if err == nil {
    c.JSON(200, user)
  } else {
    c.JSON(401, gin.H{"err": err_message})
  }
}

func splitToken(input string) (bool, string, string) {
	index := strings.Index(input, ":")
  if index == -1 {
		return false, "", ""
	}
  user, password := input[:index], input[index+1:]
  return true, user, password
}
