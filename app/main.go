package main 

import (
  "os"
  "fmt"
  "github.com/gin-gonic/gin"
  "golang.org/x/crypto/bcrypt"
  "gorm.io/driver/mysql"
  "gorm.io/gorm"
)

type User struct {
  gorm.Model
  Id  int `gorm:"primaryKey" json:"id"`
  Username string `json:"username"`
  Password string `json:"-"`
}

func main() {
  dsn := fmt.Sprintf(
    "%s:%s@tcp(%s)/%s?tls=true",
    os.Getenv("DB_USERNAME"),
    os.Getenv("DB_PASSWORD"),
    os.Getenv("DB_HOST"),
    os.Getenv("DB_DATABASE"),
  )
  
  db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

  router := gin.Default()

  router.LoadHTMLGlob("templates/*.html")

  data := "Hello Go/Gin!!"
  router.GET("/", func(c *gin.Context) {
      c.HTML(200, "index.html", gin.H{"data": data})
  })

  router.GET("/users", func(c *gin.Context) {
    var users []User
    db.Unscoped().Find(&users)
    c.JSON(200, users)
  })

  router.POST("/signup", func(c *gin.Context) {
		var newUser User
		if err := c.ShouldBindJSON(&newUser); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(500, gin.H{"error": "failed to hash password"})
			return
		}
		newUser.Password = string(hashedPassword)

		db.Create(&newUser)
		c.JSON(200, newUser)
	})

  router.Run(":3001")
}
