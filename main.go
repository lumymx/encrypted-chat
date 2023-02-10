package main

import (
	"bufio"
	"crypto/aes"
	"crypto/rand"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

func main() {
	var ip, port string
	key := make([]byte, 32)
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		fmt.Println(err)
		return
	}
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print("Create a new chat room (y/n)?")
	createRoom, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	createRoom = strings.TrimSpace(createRoom)
	if createRoom == "y" {
		ip = "localhost"
	} else {
		fmt.Print("Chat IP: ")
		ip, err = bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		ip = strings.TrimSpace(ip)
	}
	fmt.Print("Port: ")
	port, err = bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	port = strings.TrimSpace(port)
	fmt.Println(ip)
	conn, err := net.Dial("tcp", ip+":"+port)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	client := newClient(conn, "User")

	go func() {
		for {
			msg, err := client.ReadMessage(key, iv)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(msg)
		}
	}()

	if err := client.Run(key, iv); err != nil {
		fmt.Println(err)
	}
}
