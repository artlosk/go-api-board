package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func Hash(value string) string {
	return hex.EncodeToString(sha256.New().Sum([]byte(value)))
}
