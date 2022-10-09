package utils

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateAppId(appName string) string {
	bytes := make([]byte, 8)
	if _, err := rand.Read(bytes); err != nil {
		return ""
	}
	hex := hex.EncodeToString(bytes)

	return (appName + "-webapp-" + hex)
}