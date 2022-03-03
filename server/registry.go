package main

import (
	//"bufio"
	//"os"
	"fmt"
	//"math/rand"
	"net"
	"overlay_network/minichord"

	//"sort"
	"strconv"
	"strings"

	//"time"

	//"strings"

	"google.golang.org/protobuf/proto"
)

func handleRegistration() string {
	// Add registered node to registry, respond with successful or unsuccesful
	// Check to see if node is already registered and make sure ip address in request matches socket
	return "nil"
}

func handleDeregistration() string {
	// Remove from registry, check if node was registered previously, respond with successful or unsuccessful
	// Check if ip address and request matches socket
	return "nil"
}

func handleNodeRegistryResponse() string {
	// Not sure, drop it?
	return "nil"
}

func handleTaskFinished() string {
	// Once every registered node has sent a task finished message then request traffic summary from all nodes
	return "nil"
}

func handleReportTrafficSummary() string {
	// Collect and process info, make sure numbers add up
	return "nil"
}

func orderedRemove(slice [][]string, s int) [][]string {
	return append(slice[:s], slice[s+1:]...)
}

func unOrderedRemove(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func deregister(deregisterRequest *minichord.Deregistration, registeredNodes [][]string, conn net.Conn, nodeArray []int) ([][]string, []int) {
	socketString := conn.RemoteAddr().String()
	socketIpPort := strings.Split(socketString, ":")
	requestIpPort := strings.Split(deregisterRequest.GetAddress(), ":")
	if socketIpPort[0] != requestIpPort[0] {
		fmt.Println("Mismatched IP ERROR")
	}
	deregisterIndex := -1
	for i := 0; i < len(registeredNodes); i++ {
		if registeredNodes[i][1] == requestIpPort[0] {
			if registeredNodes[i][2] == requestIpPort[2] {
				deregisterIndex = i
				break
			}
		}
	}
	if deregisterIndex != -1 {
		// TODO
		fmt.Println("REMOVE FROM SLICE/ARRAY")
		identifier := registeredNodes[deregisterIndex][0]
		val, _ := strconv.Atoi(identifier)
		registeredNodes = orderedRemove(registeredNodes, deregisterIndex)
		nodeArray = append(nodeArray, val)

	}
	return registeredNodes, nodeArray

}

func main() {
	nodeArray := make([]int, 128)
	//registeredNodes := make([][]string, 0)
	//var registeredNodes[128][3]string
	for i := 0; i < 128; i++ {
		nodeArray[i] = i
	}
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
			case *minichord.MiniChord_Registration:
				fmt.Println(x.Registration.GetAddress(), "Registration response comes here!")
				handleRegistration()
			case *minichord.MiniChord_Deregistration:
				handleDeregistration()
			case *minichord.MiniChord_NodeRegistryResponse:
				handleNodeRegistryResponse()
			case *minichord.MiniChord_TaskFinished:
				handleTaskFinished()
			case *minichord.MiniChord_ReportTrafficSummary:
				handleReportTrafficSummary()
			case nil:
				fmt.Println("Error - switch case nil")
			default:
				fmt.Println("default option triggered!")

			}

			/*
				socketString := conn.RemoteAddr().String()
				socketIpPort := strings.Split(socketString, ":")
				//requestIpPort := strings.Split(message.GetAddress(), ":")
				if socketIpPort[0] != requestIpPort[0] {
					fmt.Println("Mismatched IP ERROR")
				}

				for i := 0; i < len(registeredNodes); i++ {
					if registeredNodes[i][1] == requestIpPort[0] {
						if registeredNodes[i][2] == requestIpPort[2] {
							fmt.Println("Already registered error")
							break
						}
					}
				}

				rand.Seed(time.Now().UnixNano())
				nodeArrayLength := len(nodeArray)
				randomIdentifier := rand.Intn(nodeArrayLength)
				identifier := nodeArray[randomIdentifier]
				newNode := make([]string, 3)
				newNode[0] = strconv.Itoa(identifier)
				newNode[1] = requestIpPort[0]
				newNode[2] = requestIpPort[1]
				//newNode := [...]string {strconv.Itoa(identifier), requestIpPort[0], requestIpPort[1]}
				registeredNodes = append(registeredNodes, newNode)
				sort.Slice(registeredNodes, func(i, j int) bool {
					return registeredNodes[i][0] < registeredNodes[j][0]
				})
				fmt.Println(registeredNodes)

				break
			*/
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
