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

func handleRegistrationResponse() string {
	// Decide later if it fails
	//
	return "nil"
}

func handleDeregistrationResponse() string {
	// Drop it?
	return "nil"
}

func handleNodeRegistry() string {
	// Check if you can connect to your neighbours
	// Fault tolerant? Get a new table if connections in your table don't work?
	return "nil"
}

func handleInitiateTask() string {
	// Start sending out packets to sink
	return "nil"
}

func handleNodeData() string {
	// Check if we are the sink, otherwise send to highest but not over id number
	// Add ourselfs to trace, increment hop
	return "nil"
}

func handleRequestTrafficSummary() string {
	// Respond with report traffic summary containing all our info about our sent and received packets
	return "nil"
}

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
	//addressObject := minichord.Registration{}
	/*
		testAddress := minichord.Registration{Address: "tetet"}
		address := minichord.MiniChord{
			Message: &minichord.MiniChord_Registration{Registration: &testAddress},
		}
		conn, _ := net.Dial("tcp", "localhost:45000")
		data, _ := proto.Marshal(&address)
		conn.Write(data)
	*/
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
			message := &minichord.MiniChord{}
			proto.Unmarshal(data[:length], message)

			switch x := message.Message.(type) {
			case *minichord.MiniChord_RegistrationResponse:

				fmt.Println(x.RegistrationResponse.GetInfo())
				handleRegistrationResponse()

			case *minichord.MiniChord_DeregistrationResponse:

				fmt.Println("We have been deregistered!")
				handleDeregistrationResponse()

			case *minichord.MiniChord_NodeRegistry:

				fmt.Println("Node registry response comes here!")
				handleNodeRegistry()

			case *minichord.MiniChord_InitiateTask:

				fmt.Println("Response to initiate task")
				handleInitiateTask()

			case *minichord.MiniChord_NodeData:

				fmt.Println("Response to a message node")
				handleNodeData()

			case *minichord.MiniChord_RequestTrafficSummary:

				fmt.Println("Send TrafficSummary")
				handleRequestTrafficSummary()

			default:
				fmt.Println("default option triggered!")

			}
		}
	}
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
