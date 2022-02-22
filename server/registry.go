package main

import (
	//"bufio"
	//"os"
	"fmt"
	"net"
	"overlay_network/minichord"

	//"strings"

	"google.golang.org/protobuf/proto"
)

func main() {

	ln, err := net.Listen("tcp", ":45000")
	if err != nil {
		// handle error
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
		}
		for {
			data := make([]byte, 65535)
			length, _ := conn.Read(data)
			if err != nil {
				fmt.Println("err", err)
			}
			message := &minichord.Registration{}
			proto.Unmarshal(data[:length], message)
			fmt.Println(conn.RemoteAddr())
			//if message.GetAddress() ==
			break
		}
	}

	/*
		reader := bufio.NewReader(os.Stdin)
		repl:
		for {
			cmd, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println(err)
				break
			}
			cmd = strings.Trim(cmd, "\n")
			switch (cmd) {
				case "list":
					fmt.Println("print")
				case "setup":
					fmt.Println("Bye.")
					break repl
				case "route":
					fmt.Println("print")
				case "start":
					fmt.Println("Bye.")
					break repl
				default:
					fmt.Printf("command not understood: %s\n", cmd)
			}
		}
	*/

}
