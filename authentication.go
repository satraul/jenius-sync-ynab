package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/emersion/go-imap/client"
	"github.com/emersion/go-sasl"
)

func input() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return strings.TrimSuffix(text, "\n")
}

func askConfig() (i string, u string, p string) {
	fmt.Print("Full Name or Display Name: ")
	i = input()
	fmt.Println()

	fmt.Print("Account Name, User name, or Email address: ")
	u = input()
	fmt.Println()

	fmt.Print("Password: ")
	p = input()
	fmt.Println()

	return
}

func authenticate(c *client.Client) error {
	ok, err := c.SupportAuth(sasl.Plain)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("PLANI auth not supported by the server")
	}

	// Login to the IMAP server
	saslClient := sasl.NewPlainClient(askConfig())
	return c.Authenticate(saslClient)
}
