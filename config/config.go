package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

type Config struct {
	InitialBalanceAmount int `json:"initialBalanceAmount"`
	MinumumBalanceAmount int `json:"minumumBalanceAmount"`
}

var C = &Config{}

func init() {
	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	mainPath := strings.TrimSuffix(path, "controllers")
	mainPath = strings.TrimSuffix(mainPath, "repository")
	mainPath = strings.TrimSuffix(mainPath, "service")
	mainPath = strings.TrimSuffix(mainPath, "handler")

	file, err := os.Open(mainPath + ".config/" + env + ".json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	read, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(read, C)
	if err != nil {
		panic(err)
	}
}
