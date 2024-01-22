package main 

import (
  "github.com/gin-gonic/gin"
  "github.com/gin-contrib/cors"
  "time"

  "controllers/signin"
  "controllers/signup"
  "controllers/auth"
  "controllers/session"

  "controllers/users"
  "controllers/works"
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
        // 'https://example2.com',
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
    },
    // cookieなどの情報を必要とするかどうか
    AllowCredentials: true,
    // preflightリクエストの結果をキャッシュする時間
    MaxAge: 24 * time.Hour,
  }))

  router.LoadHTMLGlob("templates/*.html")

  data := "Hello Go/Gin!!"
  router.GET("/", func(c *gin.Context) {
      c.HTML(200, "index.html", gin.H{"data": data})
  })

  router.POST("/register", auth.Register)

  router.GET("/signup", func(c *gin.Context) {
    c.HTML(200, "signup.html", gin.H{})
  })

  router.POST("/auth/signup", signup.Signup)

  // router.GET("/signin", func(c *gin.Context) {
  //   c.HTML(200, "signin.html", gin.H{})
  // })
  router.GET("/auth/session", session.Session)

  router.POST("/auth/signin", signin.Signin)

  authorized := router.Group("/admin")
  {
    authorized.GET("/users", users.Users)
    authorized.GET("/works", works.Works)
  }

  router.Run(":3001")
}
