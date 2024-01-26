package works

import (
  "gorm.io/gorm"
	"github.com/gin-gonic/gin"

  "log"

	"controllers/dbpkg"

  "controllers/crypto"
  "controllers/basicauth"
  "strings"
)

func Create(c *gin.Context) {
  isExist, _ := IsLoginUserExist(c)
  if !isExist {
    c.JSON(401, gin.H{"msg": "Unauthorized"})
    return
  }
  type Work struct {
    gorm.Model
    Name string `json:"name"`
    Url string `json:"url"`
    Description string `json:"description"`
    EncodedImg string `json:"encodedImg"`
    Tech string `json:"tech"`
  }
	db := dbpkg.GormConnect()

	var newWork Work
	if err := c.ShouldBindJSON(&newWork); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	db.Create(&newWork)
	c.JSON(200, newWork)
}

func Works(c *gin.Context) {  
  type Work struct {
    Name string `json:"name"`
    Url string `json:"url"`
    Description string `json:"description"`
    EncodedImg string `json:"encodedImg"`
    Tech string `json:"tech"`
  }

  db := dbpkg.GormConnect()

  var works []Work
  db.Unscoped().Find(&works)
  c.JSON(200, works)
}

// ログイン確認
func IsLoginUserExist(c *gin.Context) (bool, *dbpkg.User) {
  encodedToken := c.Request.Header.Get("access-token")
  // encoded-tokenをdecode
  decodedToken, err := basicauth.DecodeBase64(encodedToken)

  // decode失敗
  if err != nil {
    log.Println("認証失敗！")
    return false, nil
  }

  // decoded-tokenを分割して、ユーザとパスワードにする
  ok, tokenUsername, tokenPassword := splitToken(decodedToken)
  if !ok {
    log.Println("分割失敗！")
    return false, nil
  }

  // ユーザとパスワードがデータベースと一致するかどうか確認
  user, err := dbpkg.GetByUsername(tokenUsername)
  if err != nil || user == nil {
    return false, nil
  }
  if !crypto.CompareHashAndPassword(user.Password, tokenPassword) {
    log.Println("認証できませんでした")
    return false, nil
  }
  log.Println("認証できました")
  return true, user
}

func splitToken(input string) (bool, string, string) {
	index := strings.Index(input, ":")
  if index == -1 {
		return false, "", ""
	}
  user, password := input[:index], input[index+1:]
  return true, user, password
}
