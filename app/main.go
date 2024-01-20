package main 

import (
  "github.com/gin-gonic/gin"
  "gorm.io/gorm"

  "controllers/signin"
  "controllers/signup"
  "controllers/auth"
  "controllers/dbpkg"
)

type User struct {
  gorm.Model
  Id  int `gorm:"primaryKey" json:"id"`
  Username string `json:"Username"`
  Password string `json:"Password"`
}

func main() {
  db := dbpkg.GormConnect()

  router := gin.Default()

  router.LoadHTMLGlob("templates/*.html")

  data := "Hello Go/Gin!!"
  router.GET("/", func(c *gin.Context) {
      c.HTML(200, "index.html", gin.H{"data": data})
  })

  router.POST("/register", auth.Register)

  router.GET("/users", func(c *gin.Context) {
    var users []User
    db.Unscoped().Find(&users)
    c.JSON(200, users)
  })

  router.GET("/signup", func(c *gin.Context) {
    c.HTML(200, "signup.html", gin.H{})
  })

  router.POST("/signup", signup.Signup)

  router.GET("/signin", func(c *gin.Context) {
    c.HTML(200, "signin.html", gin.H{})
  })

  router.POST("/signin", signin.Signin)

  router.Run(":3001")
}


