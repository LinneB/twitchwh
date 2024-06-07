package twitchwh

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func generateHmac(secret, message string) string {
	hash := hmac.New(sha256.New, []byte(secret))
	hash.Write([]byte(message))
	signature := hash.Sum(nil)
	return hex.EncodeToString(signature)
}

func verifyHmac(hmac1, hmac2 string) bool {
	return hmac.Equal([]byte(hmac1), []byte(hmac2))
}
