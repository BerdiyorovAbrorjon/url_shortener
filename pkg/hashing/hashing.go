package hashing

import (
	"crypto/sha256"
	"encoding/base64"
)

func GenerateShortURL(orgURL string, len int) string {
	hash := sha256.New()
	hash.Write([]byte(orgURL))
	hashSum := hash.Sum(nil)
	base64URL := base64.RawURLEncoding.EncodeToString(hashSum)
	shortURL := base64URL[:len]
	return shortURL
}
