package main

import (
	"fmt"
)

func main() {
	graph := NewGraph()
	fileReader("examples/example05.txt", graph)
	// fileReader("examples/badexample00", graph)
	fmt.Println("Start Room:", farmInfo.Start)
	fmt.Println("End Room:", farmInfo.End)
	fmt.Println("ants num:", farmInfo.Ants)

	graph.Print()

	start := farmInfo.Start
	end := farmInfo.End

	path := DFS(graph, start, end)
	if len(path) == 0 {
		fmt.Println("no tunnels found from", start, "to", end)
	} else {
		fmt.Println("fastest path from", start, "to", end, ":", path)
	}
}
