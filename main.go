package main

import (
	"fmt"
	"time"
)

/* working
ex0 ex01 ex3 ex04 ex5 ex6-6.6724ms- ex7-71.4004ms-  be00 be01 be02
not working
ex2
*/

func main() {
	timestart := time.Now()
	graph := NewGraph()
	fileReader("examples/example04.txt", graph)

	fmt.Println("Start Room:", farmInfo.Start)
	fmt.Println("End Room:", farmInfo.End)
	if farmInfo.numAnts == 0 {
		fmt.Println("ERROR: invalid data format at least 1 ant")
		return
	}
	fmt.Println("Ants num:", farmInfo.numAnts)

	start := farmInfo.Start
	end := farmInfo.End

	paths := DFS(graph, start, end)
	if len(paths) <= 0 {
		fmt.Println("ERROR: invalid data format No pathes to end room")
		return
	}

	fmt.Println(paths)

	fmt.Println()

	rooms := make(map[string]*Rooms)
	for _, path := range paths {
		for _, room := range path {
			rooms[room] = &Rooms{Name: room, occupiedBy: nil}
		}
	}

	newPathes := RemoveBadPaths(paths)
	fmt.Println(newPathes)
	ants := chosingPathes(newPathes, farmInfo.numAnts)
	move(ants, rooms, farmInfo.End, 0)

	pased := time.Since(timestart)
	fmt.Println(pased)
}
