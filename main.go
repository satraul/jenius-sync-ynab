package main

import (
	"log"

	"github.com/emersion/go-imap/client"
)

// Gmail IMAP server
const Gmail = "imap.gmail.com:993"

func main() {
	// Connect to server
	c, err := client.DialTLS(Gmail, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected")

	defer c.Logout()

	// Login
	if err := authenticate(c); err != nil {
		log.Fatal(err)
	}
	log.Println("Logged in")
}
