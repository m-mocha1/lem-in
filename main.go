package main

import (
	"fmt"
)

func main() {
	graph := NewGraph()
	fileReader("examples/example04.txt", graph)
	// fileReader("examples/badexample00", graph)
	fmt.Println("Start Room:", farmInfo.Start)
	fmt.Println("End Room:", farmInfo.End)
	fmt.Println("ants num:", farmInfo.Ants)

	start := farmInfo.Start
	end := farmInfo.End
	// numAnts := farmInfo.Ants

	paths := DFS(graph, start, end)
	if len(paths) == 0 {
		fmt.Println("no tunnels found from", start, "to", end)
	} else {
		fmt.Println("all pathes from", start, "to", end, ":", paths)
	}

	fmt.Println(paths)
}
