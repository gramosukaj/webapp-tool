package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gramosukaj/webtool-app/pkg/config"
)

func main() {
	configPath := os.Args[1]
	configPath, err := filepath.Abs(configPath)

	if err != nil {
		panic(err)
	}

	config := config.NewConfig(configPath)

	fmt.Println(config)
}