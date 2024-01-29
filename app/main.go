package main 

import (
  "github.com/gin-gonic/gin"
  "github.com/gin-contrib/cors"
  "time"

  "controllers"
)

type User struct {
  ID  int `gorm:"primaryKey" json:"id"`
  Username string `json:"Username"`
  Password string `json:"Password"`
}

func main() {
  router := gin.Default()

  router.Use(cors.New(cors.Config{
    // アクセスを許可したいアクセス元
    AllowOrigins: []string{
        "http://localhost:3000",
        "https://ganmolt.github.io",
    },
    // アクセスを許可したいHTTPメソッド(以下の例だとPUTやDELETEはアクセスできません)
    AllowMethods: []string{
        "POST",
        "GET",
        "OPTIONS",
    },
    // 許可したいHTTPリクエストヘッダ
    AllowHeaders: []string{
        "Access-Control-Allow-Credentials",
        "Access-Control-Allow-Headers",
        "Content-Type",
        "Content-Length",
        "Accept-Encoding",
        "Authorization",
        "access-token",
        "Permissions-Policy",
    },
    // cookieなどの情報を必要とするかどうか
    AllowCredentials: true,
    // preflightリクエストの結果をキャッシュする時間
    MaxAge: 24 * time.Hour,
  }))

  data := "Hello Go/Gin!!"
  router.GET("/", func(c *gin.Context) {
    c.JSON(200, gin.H{"data": data})
  })
  // router.POST("/auth/signup", controllers.AuthController{}.Signup)
  router.POST("/auth/signin", controllers.AuthController{}.Signin)
  router.GET("/auth/session", controllers.AuthController{}.Session)

  router.GET("/works", controllers.WorksController{}.Show)

  authorized := router.Group("/admin")
  {
    authorized.GET("/users", controllers.UsersController{}.Users)
    authorized.POST("/works/create", controllers.WorksController{}.Create)
  }

  router.Run(":3001")
}
