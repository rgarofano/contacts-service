package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/rgarofano/MyCloud/common"
)

type contactsServiceConfig struct {
	Port int `json:"port"`
}

type contact struct {
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	PhoneNumber string `json:"phoneNumber"`
}

func main() {
	var cfg contactsServiceConfig
	err := config.Load("./config.json", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Starting contacts service on port %d", cfg.Port)
	http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), nil)
}
