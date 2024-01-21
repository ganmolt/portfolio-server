package users

import (
	"github.com/gin-gonic/gin"

	"log"

	"controllers/dbpkg"
)

type User struct {
  ID  int `gorm:"primaryKey" json:"id"`
  Username string `json:"Username"`
  Password string `json:"Password"`
}

func Users(c *gin.Context) {
  db := dbpkg.GormConnect()

  isExist, _ := IsLoginUserExist(c)

  if isExist {
    log.Println("isExist!")
    var users []User
    db.Unscoped().Find(&users)
    c.JSON(200, users)
  } else {
    c.JSON(401, gin.H{"msg": "Unauthorized"})
  }
}

func IsLoginUserExist(c *gin.Context) (bool, *dbpkg.User) {
  username, err := c.Cookie("username")
  if err != nil {
    log.Println("Guest")
    return false, nil
  }
  log.Println(username)

  user, err := dbpkg.GetByUsername(username)
  if err != nil || user == nil {
    return false, nil
  }
  return true, user
}

