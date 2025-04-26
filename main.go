package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type config struct {
	Port int `json:"port"`
}

func loadConfig() (*config, error) {
	jsonFile, err := os.Open("config.json")
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	jsonBytes, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var cfg config
	err = json.Unmarshal(jsonBytes, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func main() {
	config, err := loadConfig()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Starting contacts service on port %d", config.Port)
	http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil)
}
