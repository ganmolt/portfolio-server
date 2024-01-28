package usermodel

import (
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
)

func PasswordEncrypt(input string) string {
	// SHA-256ハッシュを計算
	hasher := sha256.New()
	hasher.Write([]byte(input))
	hashedBytes := hasher.Sum(nil)

	return hex.EncodeToString(hashedBytes)
}

func CompareHashAndPassword(hashedPassword string, inputPassword string) bool {
	// 保存されたハッシュをデコード
	savedHash, err := hex.DecodeString(hashedPassword)
	if err != nil {
		return false
	}

	// 入力パスワードをハッシュ化
	inputHash := sha256.Sum256([]byte(inputPassword))

	// ハッシュの比較
	return subtle.ConstantTimeCompare(savedHash, inputHash[:]) == 1
}
