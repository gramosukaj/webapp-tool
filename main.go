package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gramosukaj/webapp-tool/pkg/config"
	"github.com/gramosukaj/webapp-tool/pkg/utils"
)

func main() {
	configPath, err := filepath.Abs(os.Args[1])

	if err != nil {
		panic(err)
	}

	config := config.NewConfig(configPath)
	config.AppId = utils.GenerateAppId(config.DirName)

	fmt.Println(config.AppId)
}