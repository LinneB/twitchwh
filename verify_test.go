package twitchwh

import (
	"testing"
)

func TestGenerateHmac(t *testing.T) {
	expected := "72f7e4e306649a53f01d7353b36e9b50d49871d2d33d4588bb068356e25c6f5d"
	message := "hello world"
	secret := "supersecretstring"
	hmac := generateHmac(secret, message)
	if !verifyHmac(expected, hmac) {
		t.Fatal("HMAC verification failed")
	}
}
