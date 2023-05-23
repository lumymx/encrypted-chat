package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

type Client struct {
	conn   net.Conn
	nick   string
	reader *bufio.Reader
}

func newClient(conn net.Conn, nick string) *Client {
	return &Client{
		conn:   conn,
		nick:   nick,
		reader: bufio.NewReader(conn),
	}
}

func (c *Client) ReadMessage(key, iv []byte) (string, error) {
	encryptedMessage, err := c.reader.ReadBytes('\n')
	if err != nil {
		return "", err
	}
	message, err := DecryptMessage(encryptedMessage, key, iv)
	if err != nil {
		return "", err
	}
	return string(message), nil
}

func (c *Client) Run(key, iv []byte) error {
	for {
		fmt.Print("Enter your message or a command: ")
		input, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return err
		}
		if input[0] == '/' {
			if input == "/help\n" {
				fmt.Print("help -- display this message\n",
					"quit -- close the app\n",
					"nick [your nickname] -- choose a nickname\n")
				continue
			}
			if input == "/quit\n" {
				return nil
			}
			parts := strings.Split(input, " ")
			if parts[0] == "/nick" {
				if len(parts) > 1 {
					c.nick = strings.Join(parts[1:], " ")
					fmt.Printf("Your nickname is %s now\n", c.nick)
				} else {
					fmt.Println("Please enter your nickname")
				}
				continue
			} else {
				fmt.Println("Undefined command.")
			}
		}
		encryptedMessage, err := EncryptMessage([]byte(input), key, iv)
		if err != nil {
			fmt.Println(err)
			return err
		}
		_, err = c.conn.Write(encryptedMessage)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
}
