package main 

import (
  "github.com/gin-gonic/gin"

  "controllers/signin"
  "controllers/signup"
  "controllers/auth"

  "controllers/users"
)

type User struct {
  ID  int `gorm:"primaryKey" json:"id"`
  Username string `json:"Username"`
  Password string `json:"Password"`
}

func main() {
  router := gin.Default()

  router.LoadHTMLGlob("templates/*.html")

  data := "Hello Go/Gin!!"
  router.GET("/", func(c *gin.Context) {
      c.HTML(200, "index.html", gin.H{"data": data})
  })

  router.POST("/register", auth.Register)

  router.GET("/signup", func(c *gin.Context) {
    c.HTML(200, "signup.html", gin.H{})
  })

  router.POST("/signup", signup.Signup)

  router.GET("/signin", func(c *gin.Context) {
    c.HTML(200, "signin.html", gin.H{})
  })

  router.POST("/signin", signin.Signin)

  authorized := router.Group("/admin")
  {
    authorized.GET("/users", users.Users)
  }

  router.Run(":3001")
}
