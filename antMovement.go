package main

import "fmt"

func simulateAntMovement(path []string, numAnts int) {
    positions := make([]int, numAnts)
    occupied := make(map[string]bool)

    // Initially, mark room 0 as occupied by all ants
    for i := 0; i < numAnts; i++ {
        occupied[path[0]] = true
    }

    for step := 0; step < len(path); step++ {
        for i := 0; i < numAnts; i++ {
            // Skip ants that have reached the end of the path
            if positions[i] >= len(path) - 1 {
                continue
            }

            currentRoom := path[positions[i]]
            nextRoom := path[positions[i]+1]

            // If next room is not occupied, move the ant to the next room
            if !occupied[nextRoom] {
                positions[i]++
                occupied[currentRoom] = false
                occupied[nextRoom] = true
                fmt.Printf("Ant %d moved from %s to %s\n", i+1, currentRoom, nextRoom)
            }
        }
        fmt.Println("---")
    }
}
