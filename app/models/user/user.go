package usermodel

import (
	"gorm.io/gorm"

	"errors"
	"strings"

	"dbpkg"
)

type User struct {
  gorm.Model
  Username string `json:"username"`
  Password string `json:"password"`
}

func Create(newUser User) (*User, error) {
	hashedPassword := PasswordEncrypt(newUser.Password)
	newUser.Password = hashedPassword

	db := dbpkg.GormConnect()
	result := db.Create(&newUser)
	if result.Error != nil {
		return nil, result.Error
	}
	return &newUser, nil
}

func Signin(username string, password string) (string, string) {
	dbUser, err := getByUsername(username)

	if err != nil || dbUser == nil {
		return "", "ユーザー名またはパスワードが違います"
	}

	if CompareHashAndPassword(dbUser.Password, password) {
		raw_token := username + ":" + password
		access_token := EncodeBase64(raw_token)
		return access_token, ""
	} else {
		return "", "ユーザー名またはパスワードが違います"
	}
}

func Session(accessToken string) (*User, string) {
  decodedToken, err := DecodeBase64(accessToken)

	ok, username, password := splitToken(decodedToken)
	if !ok {
		return nil, "アクセストークンが誤っています"
  }

	dbUser, err := getByUsername(username)
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

func getByUsername(username string) (*User, error) {
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
