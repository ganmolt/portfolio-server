package users

import (
	"github.com/gin-gonic/gin"

	"log"

	"controllers/dbpkg"

  "controllers/basicauth"
  "strings"

  "models/user"
)

func Users(c *gin.Context) {
  db := dbpkg.GormConnect()

  isExist, _ := IsLoginUserExist(c)

  if isExist {
    log.Println("isExist!")
    var users []dbpkg.User
    db.Unscoped().Find(&users)
    c.JSON(200, users)
  } else {
    c.JSON(401, gin.H{"msg": "Unauthorized"})
  }
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
  if !usermodel.CompareHashAndPassword(user.Password, tokenPassword) {
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
