package main

import "fmt"

func main() {
	graph := NewGraph()
	fileReader("examples/example00", graph)
	fmt.Println("Start Room:", farmInfo.Start)
	fmt.Println("End Room:", farmInfo.End)
	fmt.Println("ants num:", farmInfo.Ants)

	for _, room := range farmInfo.Rooms {
		fmt.Println("room:", room.Name, "location:", room.Location)
	}
	for _, tunnel := range farmInfo.tunnels {
		fmt.Println("tunnel from", tunnel.From, "to", tunnel.To)
	}

	graph.Print()
}
