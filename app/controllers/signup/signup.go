package signup

import (
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"

	"controllers/dbpkg"
	"controllers/crypto"
)

type User struct {
  gorm.Model
  Id  int `gorm:"primaryKey" json:"id"`
  Username string `json:"username"`
  Password string `json:"password"`
}

func Signup(c *gin.Context) {
	db := dbpkg.GormConnect()

	var newUser User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	hashedPassword := crypto.PasswordEncrypt(newUser.Password)
	newUser.Password = hashedPassword

	db.Create(&newUser)
	c.JSON(200, newUser)
}
