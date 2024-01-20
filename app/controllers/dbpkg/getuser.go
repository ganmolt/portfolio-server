package dbpkg

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id        int        `gorm:"primaryKey" json:"id"`
	Username  string     `json:"username"`
	Password  string     `json:"password"`
	CreatedAt string     `json:"created_at"`
	UpdatedAt string     `json:"updated_at"`
}

func GetUser(username string) User {
	db := GormConnect()
	var user User
	db.First(&user, "username = ?", username)
	return user
}
