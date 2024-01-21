package users

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"log"
	"os"

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

    var users []User
    db.Unscoped().Find(&users)
    // log.Println(users[0])
    // Login(c, users[0])
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

func Login(c *gin.Context, user User) {
  c.SetSameSite(http.SameSiteNoneMode) // samesiteをnonemodeにする
  if os.Getenv("ENV") == "local" {
    log.Println("cookieをセットする")
    c.SetCookie("username", user.Username, 3600, "/", "localhost:3001", true, true)
  }
}

