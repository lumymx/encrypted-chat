package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	var ip, port string
	key, iv, err := GenerateKeyAndIV()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print("Create a new chat room (y/n)? ")
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

	var conn net.Conn
	if ip == "localhost" {
		ln, err := net.Listen("tcp", ":"+port)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Listening on port " + port + "...")
		conn, err = ln.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		defer ln.Close()
	} else {
		conn, err = net.Dial("tcp", ip+":"+port)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer conn.Close()
	}

	client, err := NewClient(conn, "User", key, iv)
	if err != nil {
		fmt.Println(err)
		return
	}

	go func() {
		for {
			msg, err := client.ReadMessage()
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(msg)
		}
	}()

	if err := client.Run(); err != nil {
		fmt.Println(err)
	}
}
