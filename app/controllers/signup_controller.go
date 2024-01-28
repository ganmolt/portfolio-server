package controllers

import (
	"github.com/gin-gonic/gin"

	"models/user"
)

func Signup(c *gin.Context) {
	user := usermodel.Create(c)
	c.JSON(200, user)
}
