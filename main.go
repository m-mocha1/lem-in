package main

import (
	"fmt"
)

func main() {
	graph := NewGraph()
	fileReader("examples/example01.txt", graph)
	// fileReader("examples/badexample00", graph)
	fmt.Println("Start Room:", farmInfo.Start)
	fmt.Println("End Room:", farmInfo.End)
	fmt.Println("ants num:", farmInfo.Ants)

	start := farmInfo.Start
	end := farmInfo.End

	path := DFS(graph, start, end)
	if len(path) == 0 {
		fmt.Println("no tunnels found from", start, "to", end)
	} else {
		fmt.Println("fastest path from", start, "to", end, ":", path)
	}
	ants := make([]Ant, farmInfo.Ants)
	for i := range ants {
		ants[i] = Ant{Room: start}
	}
	var step int

	for !End(ants, end) {
		step++
		if step > 20 {
			break
		}
		Move(graph, ants, end)
		fmt.Println("step", step)
		for i, ant := range ants {
			fmt.Print("L", i+1, "-", ant.Room, "\n")
		}
		fmt.Println()
	}
}
