package main 

import (
  "os"
  "fmt"
  "github.com/gin-gonic/gin"
  "gorm.io/driver/mysql"
  "gorm.io/gorm"
)

type User struct {
  gorm.Model
  Id  int `gorm:"primaryKey" json:"id"`
  Name string `json:"name"`
  Age  int `json:"age"`
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

  router.Run(":3001")
}
