package main

import (
	"encoding/json"
	"log"
	"os"
)

type AppConfiguration struct {
	LootDir string
	LogDir  string
}

func main() {
	var configFile = "./config.json"
	c := getConfiguration(configFile)
}

func getConfiguration(configurationFile string) *AppConfiguration {
	file, err := os.Open(configurationFile)
	if err != nil {
		log.Fatalf("Can't open a configuration file")
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	config := new(AppConfiguration)
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatalf("Can't decode a configuration file")
	}
	return config
}
