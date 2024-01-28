package usermodel

import (
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"

	"controllers/dbpkg"
	"controllers/crypto"
	"controllers/basicauth"
)

type User struct {
  gorm.Model
  Id  int `gorm:"primaryKey" json:"id"`
  Username string `json:"username"`
  Password string `json:"password"`
}

func Create(c *gin.Context) (*User) {
	db := dbpkg.GormConnect()

	var newUser User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return nil
	}
	hashedPassword := crypto.PasswordEncrypt(newUser.Password)
	newUser.Password = hashedPassword

	db.Create(&newUser)

	return &newUser
}

func Signin(c *gin.Context) (string, string) {
	var input User
	if err := c.ShouldBindJSON(&input); err != nil {
		return "", string(err.Error())
	}

	dbUser, err := dbpkg.GetByUsername(input.Username)

	if err != nil || dbUser == nil {
		return "", "Invalid username or password"
	}

	if crypto.CompareHashAndPassword(dbUser.Password, input.Password) {
		raw_token := input.Username + ":" + input.Password
		access_token := basicauth.EncodeBase64(raw_token)
		return access_token, ""
	} else {
		return "", "Invalid username or password"
	}
}
