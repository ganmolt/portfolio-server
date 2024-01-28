package usermodel

import (
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"

	"errors"

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

func Session(username string, password string) (*User, string) {
	dbUser, err := GetByUsername(username)
	if err != nil || dbUser == nil {
		return nil, "Invalid username or password"
	}

	if crypto.CompareHashAndPassword(dbUser.Password, password) {
		return dbUser, ""
	} else {
		return nil, "Invalid username or password"
	}
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
