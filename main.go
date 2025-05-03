package main

import (
	"fmt"
	"github.com/rgarofano/MyCloud/common"
	"log"
	"net/http"
	"os"
)

type contactsServiceConfig struct {
	Port             int    `json:"port"`
	ContactsFileName string `json:"contactsFileName"`
}

type contact struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	PhoneNumber string `json:"phoneNumber"`
}

var cfg contactsServiceConfig
var contacts []contact

func init() {
	err := utils.LoadJson("config.json", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := os.Stat(cfg.ContactsFileName); err == nil {
		err = utils.LoadJson(cfg.ContactsFileName, &contacts)
		if err != nil {
			log.Fatal(err)
		}
	} else if os.IsNotExist(err) {
		contacts = make([]contact, 0)
	} else {
		log.Fatal(err)
	}
}

func main() {
	log.Printf("Starting contacts service on port %d", cfg.Port)
	http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), nil)
}
