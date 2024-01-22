package session

import (
	"github.com/gin-gonic/gin"

	"log"

	"controllers/dbpkg"

  "controllers/crypto"
  "controllers/basicauth"
  "strings"
)

func Session(c *gin.Context) {
  isExist, user := IsLoginUserExist(c)
  if isExist {
    c.JSON(200, user)
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
