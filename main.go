package main

import (
	"fmt"
	"time"
)

func main() {
	timestart := time.Now()
	graph := NewGraph()
	fileReader("examples/example05.txt", graph)
	// fileReader("examples/badexample00.txt", graph)

	if farmInfo.numAnts == 0 {
		fmt.Println("ERROR: invalid data format at least 1 ant")
		return
	}
	fmt.Println("Start Room:", farmInfo.Start)
	fmt.Println("End Room:", farmInfo.End)
	fmt.Println("Ants num:", farmInfo.numAnts)
	fmt.Println("there is ", len(graph.Nodes), "nodes")

	start := farmInfo.Start
	end := farmInfo.End

	paths := DFS(graph, start, end)
	if len(paths) <= 0 {
		fmt.Println("ERROR: invalid data format No pathes to end room")
		return
	}

	fmt.Println(paths)

	fmt.Println()

	rooms := makeRooms(paths)

	newPathes := RemoveBadPaths(paths)
	fmt.Println(newPathes)
	ants := creatingAnts(newPathes, farmInfo.numAnts)
	move(ants, rooms, farmInfo.End, 0)

	pased := time.Since(timestart)
	fmt.Println("it took ", pased, " to excute the program")
}
