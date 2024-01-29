package controllers

import (
	"github.com/gin-gonic/gin"

	"dbpkg"

  "models/user"
)

type UsersController struct{}

func (uc UsersController) Users(c *gin.Context) {
  access_token := c.Request.Header.Get("access-token")
	_, errMessage := usermodel.Session(access_token)
  if errMessage != "" {
    c.JSON(401, gin.H{"err": errMessage})
  }

  db := dbpkg.GormConnect()

  type User struct {
    Username string `json:"username"`
  }

  var usernames []User
  db.Unscoped().Select("Username").Find(&usernames)
  c.JSON(200, usernames)
}
