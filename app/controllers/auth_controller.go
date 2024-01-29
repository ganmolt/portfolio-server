package controllers

import (
	"github.com/gin-gonic/gin"

	"log"
	"net/http"

	"models/user"
)

type AuthController struct{}

func (ac AuthController) Signup(c *gin.Context) {
	var newUser usermodel.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	user, err := usermodel.Create(newUser)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, user)
}

func (ac AuthController) Signin(c *gin.Context) {
	var input usermodel.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}

	token, err_message := usermodel.Signin(input.Username, input.Password)
	if err_message != "" {
		log.Println("ログインできませんでした")
		c.JSON(400, gin.H{"error": err_message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"access-token": token})
	log.Println("ログインできました")
	c.Redirect(302, "/")
}

func (ac AuthController) Session(c *gin.Context) {
	access_token := c.Request.Header.Get("access-token")
	user, errMessage := usermodel.Session(access_token)

  if errMessage == "" {
    c.JSON(200, user)
  } else {
    c.JSON(401, gin.H{"err": errMessage})
  }
}
