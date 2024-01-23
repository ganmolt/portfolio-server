package dbpkg

import (
	"gorm.io/gorm"

  "errors"
)

type User struct {
  Username string `json:"username"`
  Password string `json:"password"`
}

func GetByUsername(username string) (*User, error) {
	db := GormConnect()
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

