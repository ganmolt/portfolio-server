package basicauth

import (
	"encoding/base64"
)

// Base64エンコード関数
func EncodeBase64(data string) string {
	input := []byte(data)

	encoded := base64.StdEncoding.EncodeToString(input)
	return encoded
}

// Base64デコード関数
func DecodeBase64(encoded string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return "", err
	}

	decodedString := string(decoded)
	return decodedString, nil
}
