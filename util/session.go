package util

import (
	"crypto/rand"
	"encoding/hex"
)

type Session struct {
	Token string
	Expiry int64
}


var Sessions []Session


// https://www.jetbrains.com/guide/go/tutorials/authentication-for-go-apps/auth/

func RandomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
