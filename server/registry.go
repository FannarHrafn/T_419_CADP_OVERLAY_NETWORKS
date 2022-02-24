package main

import (
	//"bufio"
	//"os"
	"sort"
	"fmt"
	"net"
	"overlay_network/minichord"
	"strings"
	"math/rand"
	"time"
	"strconv"
	//"reflect"
	//"strings"

	"google.golang.org/protobuf/proto"
)

func orderedRemove(slice [][]string, s int) [][]string {
    return append(slice[:s], slice[s+1:]...)
}

func unOrderedRemove(s []int, i int) []int {
    s[i] = s[len(s)-1]
    return s[:len(s)-1]
}

func deregister(deregisterRequest *minichord.Deregistration, registeredNodes [][]string, conn net.Conn, nodeArray []int) ([][]string, []int){
	socketString := conn.RemoteAddr().String()
	socketIpPort := strings.Split(socketString, ":")
	requestIpPort := strings.Split(deregisterRequest.GetAddress(), ":")
	if socketIpPort[0] != requestIpPort[0]{
		fmt.Println("Mismatched IP ERROR")
	}
	deregisterIndex := -1
	for i := 0; i < len(registeredNodes); i++{
		if registeredNodes[i][1] == requestIpPort[0]{
			if registeredNodes[i][2] == requestIpPort[2]{
				deregisterIndex = i
				break
			}
		}
	}
	if deregisterIndex != -1{
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
	registeredNodes := make([][]string , 0)
	//var registeredNodes[128][3]string
	for i := 0; i < 128; i++{
		nodeArray[i] = i;
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
			message := &minichord.Registration{}
			proto.Unmarshal(data[:length], message)
			socketString := conn.RemoteAddr().String()
			socketIpPort := strings.Split(socketString, ":")
			requestIpPort := strings.Split(message.GetAddress(), ":")
			if socketIpPort[0] != requestIpPort[0]{
				fmt.Println("Mismatched IP ERROR")
			} 

			for i := 0; i < len(registeredNodes); i++ {
				if registeredNodes[i][1] == requestIpPort[0]{
					if registeredNodes[i][2] == requestIpPort[2]{
						fmt.Println("Already registered error")
						break
					}
				}
			}
			
			rand.Seed(time.Now().UnixNano())
			nodeArrayLength := len(nodeArray)
			randomIdentifier := rand.Intn(nodeArrayLength)
			identifier := nodeArray[randomIdentifier]
			newNode := make([]string , 3)
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
