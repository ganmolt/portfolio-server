package users

import (
	"github.com/gin-gonic/gin"

	"controllers/dbpkg"

  "models/user"
)

func Users(c *gin.Context) {
  access_token := c.Request.Header.Get("access-token")
	_, errMessage := usermodel.Session(access_token)
  if errMessage != "" {
    c.JSON(401, gin.H{"err": errMessage})
  }

  db := dbpkg.GormConnect()

  var users []dbpkg.User
  db.Unscoped().Find(&users)
  c.JSON(200, users)
}
