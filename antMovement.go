package main

import (
	"fmt"
)

func calculatePathCost(path []string, antsInPath int) int {
	return len(path) + antsInPath
}

func assignPathToAnt(ant *Ant, paths [][]string, antsInPath []int) {
	bestPathIndex := 0
	bestPathCost := calculatePathCost(paths[0], antsInPath[0])

	for i := 1; i < len(paths); i++ {
		currentPathCost := calculatePathCost(paths[i], antsInPath[i])
		if currentPathCost < bestPathCost {
			bestPathIndex = i
			bestPathCost = currentPathCost
		}
	}
	ant.path = paths[bestPathIndex]
	antsInPath[bestPathIndex]++
}

func chosingPathes(pathes [][]string, n int) []Ant {
	ants := make([]Ant, n)
	antsInPath := make([]int, len(pathes))

	for i := 0; i < n; i++ {
		assignPathToAnt(&ants[i], pathes, antsInPath)
		ants[i].id = i + 1
		ants[i].pos = 0
	}
	return ants
}

func move(ants []Ant, rooms map[string]*Rooms, endRoom string, steps int) {
	anyMoved := false

	fmt.Println(steps)
	for i := range ants {
		ant := &ants[i]
		if ant.End {
			continue
		}

		currentRoom := rooms[ant.path[ant.pos]]
		nextRoom := rooms[ant.path[ant.pos+1]]

		if nextRoom.Name != endRoom && nextRoom.occupiedBy == nil {
			currentRoom.occupiedBy = nil
			nextRoom.occupiedBy = ant
			ant.currentRoom = nextRoom.Name
			ant.pos++
			fmt.Print("L", ant.id, "-", nextRoom.Name, " ")
			anyMoved = true
		} else if nextRoom.Name == endRoom {
			ant.End = true
			currentRoom.occupiedBy = nil
			ant.currentRoom = endRoom
			ant.pos++
			fmt.Print("L", ant.id, "-", nextRoom.Name, " ")
			anyMoved = true
		}
	}

	if anyMoved {
		move(ants, rooms, endRoom, steps+1)
	}
}
