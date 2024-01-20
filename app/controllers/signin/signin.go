package signin

import (
	"github.com/gin-gonic/gin"

	"log"
	"net/http"

	"controllers/dbpkg"
	"controllers/crypto"
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
		log.Println("ログインできました")
		c.Redirect(302, "/")
	} else {
		log.Println("ログインできませんでした")
    c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid username or password"})
    c.Abort()
	}
}
