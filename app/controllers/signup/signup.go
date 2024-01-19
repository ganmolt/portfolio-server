package signup

import (
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"controllers/dbpkg"
)

type User struct {
  gorm.Model
  Id  int `gorm:"primaryKey" json:"id"`
  Username string `json:"username"`
  Password string `json:"-"`
}

func Signup(c *gin.Context) {
	db := dbpkg.GormConnect()

	var newUser User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := PasswordEncrypt(newUser.Password)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to hash password"})
		return
	}
	newUser.Password = string(hashedPassword)

	db.Create(&newUser)
	c.JSON(200, newUser)
}


func PasswordEncrypt(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

