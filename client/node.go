package main

import (
	"bufio"
	//"fmt"
	"fmt"
	"net"
	"os"
	"overlay_network/minichord"
	"strings"

	"google.golang.org/protobuf/proto"
)

func main() {
	/*
		//a := minichord.TrafficSummary {Id: 0, Sent: 0, Relayed: 0, Received: 0, TotalSent: 0, TotalReceived: 0}
		conn, err := net.Dial("tcp", "localhost:45000")
		if err != nil {
			// handle error
		}
		data, _ := proto.Marshal(message.ProtoReflect().Interface())
		end, _ := conn.Write(data)
		fmt.Println(end)
	*/
	// ...
	//message := minichord.TrafficSummary{Id: 5, Sent: 2, Relayed: 1, Received: 2, TotalSent: 2, TotalReceived: 2}
	address := minichord.Registration{Address: "127.0.0.1:45000"}
	conn, _ := net.Dial("tcp", "localhost:45000")
	data, _ := proto.Marshal(&address)
	conn.Write(data)

	reader := bufio.NewReader(os.Stdin)
repl:
	for {
		cmd, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			break
		}
		cmd = strings.Trim(cmd, "\n")
		switch cmd {
		case "print":
			fmt.Println("print")
		case "exit":
			fmt.Println("Bye.")
			break repl
		default:
			fmt.Printf("command not understood: %s\n", cmd)
		}
	}
}
