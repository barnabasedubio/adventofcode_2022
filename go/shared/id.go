package shared

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateUID(length int) string {
	if length <= 0 {
		panic("Can not create ID from length <= 0")
	}
	randomBytes := make([]byte, length)
	if _, err := rand.Read(randomBytes); err != nil {
		panic("Failed to generate random bytes")
	}
	randomID := base64.URLEncoding.EncodeToString(randomBytes)
	return randomID[:length]
}
