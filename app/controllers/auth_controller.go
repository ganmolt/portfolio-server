package controllers

import (
	"github.com/gin-gonic/gin"

	"log"
	"net/http"

	"models/user"
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
	access_token := c.Request.Header.Get("access-token")
	user, errMessage := usermodel.Session(access_token)

  if errMessage == "" {
    c.JSON(200, user)
  } else {
    c.JSON(401, gin.H{"err": errMessage})
  }
}
