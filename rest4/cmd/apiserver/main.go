package main

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"

	"go-learn/base"
	"go-learn/rest4/internal/app/apiserver"

	"github.com/BurntSushi/toml"
)

// Путь
var path = base.Dir("rest4")

// Путь конфигурации
var configPath string

func init() {
	flag.StringVar(&configPath, "config-path",
		"configs/apiserver.toml", "Путь к файлу конфигурации")
}

func main() {
	fmt.Println(" \n[ REST (GOPHER SCHOOL) ]\n ")

	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(filepath.Join(path, configPath), config)
	if err != nil {
		log.Fatal(err)
	}

	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}
}
