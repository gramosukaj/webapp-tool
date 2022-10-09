package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"

	"github.com/gramosukaj/webapp-tool/pkg/config"
)

func main() {
	configPath, err := filepath.Abs(os.Args[1])

	if err != nil {
		panic(err)
	}

	config := config.NewConfig(configPath)
	config.AppId = generateAppId(config.DirName)

	fmt.Println(config.AppId)
}

func generateAppId(appName string) string {
	bytes := make([]byte, 8)
	if _, err := rand.Read(bytes); err != nil {
		return ""
	}
	hex := hex.EncodeToString(bytes)

	return (appName + "-webapp-" + hex)
}