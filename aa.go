package main

import (
	"fmt"

	"github.com/konimarti/opc"
)

func main() {
	var directory string = "hello"
	fmt.Printf(directory)

	client, _ := opc.NewConnection(
		"MatriKon.OPC.Simulation.1",     // ProgId
		[]string{"localhost"},           // Nodes
		[]string{"Group0.Random.Real8"}, // Tags
	)
	defer client.Close()
	client.ReadItem("Group0.Random.Real8")
	fmt.Printf("ssaddsd")
}
