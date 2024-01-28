package usermodel

import (
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"

	"errors"
	"strings"

	"controllers/dbpkg"
)

type User struct {
  gorm.Model
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
	hashedPassword := PasswordEncrypt(newUser.Password)
	newUser.Password = hashedPassword

	db.Create(&newUser)

	return &newUser
}

func Signin(c *gin.Context) (string, string) {
	var input User
	if err := c.ShouldBindJSON(&input); err != nil {
		return "", string(err.Error())
	}

	dbUser, err := GetByUsername(input.Username)

	if err != nil || dbUser == nil {
		return "", "Invalid username or password"
	}

	if CompareHashAndPassword(dbUser.Password, input.Password) {
		raw_token := input.Username + ":" + input.Password
		access_token := EncodeBase64(raw_token)
		return access_token, ""
	} else {
		return "", "Invalid username or password"
	}
}

func Session(accessToken string) (*User, string) {
  decodedToken, err := DecodeBase64(accessToken)

	ok, username, password := splitToken(decodedToken)
	if !ok {
		return nil, "アクセストークンが誤っています"
  }

	dbUser, err := GetByUsername(username)
	if err != nil || dbUser == nil {
		return nil, "ユーザー名またはパスワードが違います"
	}

	if CompareHashAndPassword(dbUser.Password, password) {
		return dbUser, ""
	} else {
		return nil, "ユーザー名またはパスワードが違います"
	}
}

func splitToken(input string) (bool, string, string) {
	index := strings.Index(input, ":")
  if index == -1 {
		return false, "", ""
	}
  user, password := input[:index], input[index+1:]
  return true, user, password
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
