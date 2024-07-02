package main

import (
	"fmt"
	"time"
)

/* working
ex0 ex3 ex5 ex6-6.1264ms- ex7-24.5509ms-  be00 be01 be02
not working
ex01 ex04
*/

func main() {
	timestart := time.Now()
	graph := NewGraph()
	fileReader("examples/example07.txt", graph)
	// fileReader("examples/badexample02.txt", graph)

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

	sorting(paths)
	fmt.Println(paths)

	
	fmt.Println()

	// fmt.Print(chosingPathes(paths, farmInfo.numAnts))

	rooms := make(map[string]*Rooms)
	for _, path := range paths {
		for _, room := range path {
			if _, exists := rooms[room]; !exists {
				rooms[room] = &Rooms{Name: room, occupiedBy: nil}
			}
		}
	}

	newPathes := RemoveTooLongPaths(paths)
	fmt.Println(newPathes)
	ants := chosingPathes(paths, farmInfo.numAnts)
	// moveAnts(ants, rooms, farmInfo.Start, farmInfo.End)
	move(ants, rooms, farmInfo.End, 0)

	pased := time.Since(timestart)
	fmt.Println(pased)
}

