package main

import (
	
	"fmt"

	"overlay_network/minichord"
)

func main() {
	
	a := minichord.TrafficSummary {Id: 0, Sent: 0, Relayed: 0, Received: 0, TotalSent: 0, TotalReceived: 0}
	fmt.Println(a)

}



