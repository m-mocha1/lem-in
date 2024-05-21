package main

import (
	"fmt"
	"strconv"
)

func main() {
	graph := NewGraph()
	fileReader("examples/example03", graph)
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

	start, _ := strconv.Atoi(farmInfo.Start)
	end, _ := strconv.Atoi(farmInfo.End)

	path := BFS(graph, start, end)
	fmt.Println("Shortest path", path)
}
