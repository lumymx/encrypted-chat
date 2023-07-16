package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

type Client struct {
	conn   net.Conn
	nick   string
	reader *bufio.Reader
	key    []byte
	iv     []byte
}

func NewClient(conn net.Conn, nick string, key, iv []byte) (*Client, error) {
	return &Client{
		conn:   conn,
		nick:   nick,
		reader: bufio.NewReader(os.Stdin),
		key:    key,
		iv:     iv,
	}, nil
}

func (c *Client) Run() error {
	fmt.Println("Welcome to the chat room!")
	for {
		msg, err := c.reader.ReadString('\n')
		if err != nil {
			return err
		}
		msg = strings.TrimSpace(msg)
		if msg == "/quit" {
			return nil
		}
		encryptedMsg, err := EncryptMessage([]byte(msg), c.key, c.iv)
		if err != nil {
			return err
		}
		if _, err := c.conn.Write(encryptedMsg); err != nil {
			return err
		}
	}
}

func (c *Client) ReadMessage() (string, error) {
	encryptedMsg := make([]byte, 1024)
	if err != nil {
		return "", err
	}
	n, err := c.conn.Read(encryptedMsg)
	if err != nil {
		return "", err
	}
	decryptedMsg, err := DecryptMessage(encryptedMsg[:n], c.key, c.iv)
	if err != nil {
		return "", err
	}
	return string(decryptedMsg), nil
}
