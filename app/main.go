package main 

import (
  "github.com/gin-gonic/gin"
  "golang.org/x/crypto/bcrypt"
  "gorm.io/gorm"

  "log"
  "net/http"

  "controllers/signup"
  "controllers/auth"
  "controllers/dbpkg"
)

type User struct {
  gorm.Model
  Id  int `gorm:"primaryKey" json:"id"`
  Username string `json:"username"`
  Password string `json:"-"`
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

  router.POST("/signin", func(c *gin.Context) {
    log.Println(c)
    if (1 >= 0) {
      // c.JSON(200, "asdf")
      c.JSON(200, c.PostForm("username"))
      // c.PostForm("Username")
      return
    }
    // DBから取得したユーザーパスワード(Hash)
    dbPassword := dbpkg.GetUser(c.PostForm("username")).Password
    log.Println(dbPassword)
    // フォームから取得したユーザーパスワード
    formPassword := c.PostForm("password")

    // ユーザーパスワードの比較
    if err := CompareHashAndPassword(dbPassword, formPassword); err != nil {
      log.Println(dbPassword, formPassword)
      log.Println("ログインできませんでした")
      c.HTML(http.StatusBadRequest, "signin.html", gin.H{"err": err})
      c.Abort()
    } else {
      log.Println("ログインできました")
      c.Redirect(302, "/")
    }
  })
  router.Run(":3001")
}

func PasswordEncrypt(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}
func CompareHashAndPassword(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

