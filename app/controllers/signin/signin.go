package signin

import (
	"github.com/gin-gonic/gin"

	"log"
	"net/http"

	"gorm.io/gorm"
	"errors"

	"controllers/dbpkg"
	"controllers/crypto"
)

type SigninData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
  gorm.Model
  Username string `json:"username"`
  Password string `json:"password"`
}

func GetByUsername(username string) (*User, error) {
	db := dbpkg.GormConnect()
	var dbUser User
	if err := db.
						Select("username", "password").
						First(&dbUser, "username = ?", username).
						Error; err != nil {
							if errors.Is(err, gorm.ErrRecordNotFound) {
							  return nil, nil
							}
							  return nil, err
						}
	return &dbUser, nil
}

func Signin(c *gin.Context) {
	var input SigninData
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dbUser, err := GetByUsername(input.Username)

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
