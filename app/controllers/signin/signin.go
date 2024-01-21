package signin

import (
	"github.com/gin-gonic/gin"

	"log"
	"net/http"

	"controllers/dbpkg"
	"controllers/crypto"

	"os"
)

type SigninData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Signin(c *gin.Context) {
	var input SigninData
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

func Login(c *gin.Context, user SigninData) {
  c.SetSameSite(http.SameSiteNoneMode) // samesiteをnonemodeにする
  if os.Getenv("ENV") == "local" {
    log.Println("cookieをセットする")
    c.SetCookie("username", user.Username, 3600, "/", "localhost:3001", true, true)
  }
}
