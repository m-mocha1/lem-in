package main

import "fmt"

func roomsPlusAnts(path []string, antsInPath int) int {
	return len(path) + antsInPath
}

func chosingPathe(ant *Ant, paths [][]string, antsInPath []int) {
	bestPathIndex := 0
	bestPathCost := roomsPlusAnts(paths[0], antsInPath[0])

	for i := 1; i < len(paths); i++ {
		currentPathCost := roomsPlusAnts(paths[i], antsInPath[i])
		if currentPathCost < bestPathCost {
			bestPathIndex = i
			bestPathCost = currentPathCost
		}
	}
	ant.path = paths[bestPathIndex]
	antsInPath[bestPathIndex]++
}

func creatingAnts(pathes [][]string, n int) []Ant {
	ants := make([]Ant, n)
	antsInPath := make([]int, len(pathes))

	for i := 0; i < n; i++ {
		chosingPathe(&ants[i], pathes, antsInPath)
		ants[i].id = i + 1
		ants[i].pos = 0
	}
	return ants
}

func makeRooms(pathes [][]string) map[string]*Rooms {
	rooms := make(map[string]*Rooms)
	for _, path := range pathes {
		for _, room := range path {
			rooms[room] = &Rooms{Name: room, occupiedBy: nil}
		}
	}
	return rooms
}

func move(ants []Ant, rooms map[string]*Rooms, endRoom string, steps int) {
	fmt.Println(steps)
	tunnel := make(map[string]bool) 

	for i := range ants {
		ant := &ants[i]
		if ant.End {
			continue
		}

		currentRoom := rooms[ant.path[ant.pos]]
		nextRoom := rooms[ant.path[ant.pos+1]]

		tunnelName := currentRoom.Name + " to " + nextRoom.Name 
		if tunnel[tunnelName] {
			continue 
		}

		if nextRoom.Name != endRoom && nextRoom.occupiedBy == nil {
			currentRoom.occupiedBy = nil
			nextRoom.occupiedBy = ant
			ant.currentRoom = nextRoom.Name
			ant.pos++
			fmt.Print("L", ant.id, "-", nextRoom.Name, " ")
			tunnel[tunnelName] = true
		} else if nextRoom.Name == endRoom {
			ant.End = true
			currentRoom.occupiedBy = nil
			ant.currentRoom = endRoom
			ant.pos++
			fmt.Print("L", ant.id, "-", nextRoom.Name, " ")
			tunnel[tunnelName] = true
		}
	}
	if len(tunnel) > 0 {
		move(ants, rooms, endRoom, steps+1)
	}
}

// func move(ants []Ant, rooms map[string]*Rooms, endRoom string, steps int) {
// 	Moving := false

// 	fmt.Println(steps)

// 	for i := range ants {
// 		ant := &ants[i]
// 		if ant.End {
// 			continue
// 		}

// 		currentRoom := rooms[ant.path[ant.pos]]
// 		nextRoom := rooms[ant.path[ant.pos+1]]

// 		if nextRoom.Name != endRoom && nextRoom.occupiedBy == nil {
// 			currentRoom.occupiedBy = nil
// 			nextRoom.occupiedBy = ant
// 			ant.currentRoom = nextRoom.Name
// 			ant.pos++
// 			fmt.Print("L", ant.id, "-", nextRoom.Name, " ")
// 			Moving = true
// 		} else if nextRoom.Name == endRoom {
// 			ant.End = true
// 			currentRoom.occupiedBy = nil
// 			ant.currentRoom = endRoom
// 			ant.pos++
// 			fmt.Print("L", ant.id, "-", nextRoom.Name, " ")
// 			Moving = true
// 		}
// 	}

// 	if Moving {
// 		move(ants, rooms, endRoom, steps+1)
// 		}
// 		}
