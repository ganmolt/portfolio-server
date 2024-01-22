package signin

import (
	"github.com/gin-gonic/gin"

	"log"
	"net/http"

	"controllers/dbpkg"
	"controllers/crypto"

	"os"
	"controllers/basicauth"
)

func Signin(c *gin.Context) {
	var input dbpkg.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dbUser, err := dbpkg.GetByUsername(input.Username)

	if err != nil || dbUser == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid username or password"})
		c.Abort()
		return
	}

	if crypto.CompareHashAndPassword(dbUser.Password, input.Password) {
		Login(c, input)
		log.Println("ログインできました")
		c.Redirect(302, "/")
	} else {
		log.Println("ログインできませんでした")
    c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid username or password"})
    c.Abort()
	}
}

func Login(c *gin.Context, user dbpkg.User) {
  c.SetSameSite(http.SameSiteNoneMode) // samesiteをnonemodeにする
  if os.Getenv("ENV") == "local" {
		raw_token := user.Username + ":" + user.Password
		access_token := basicauth.EncodeBase64(raw_token)

		c.SetCookie("access-token", access_token, 3600, "/", "localhost:3001", true, true)
  }
}
