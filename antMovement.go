package main

import "fmt"

func simulateAntMovement(path []string, numAnts int) {
    positions := make([]int, numAnts)
    occupied := make(map[string]bool)

    for step := 0; step < len(path); step++ {
        for i := 0; i < numAnts; i++ {
            if positions[i] >= len(path) - 1 {
                continue
            }

            currentRoom := path[positions[i]]
            nextRoom := path[positions[i]+1]

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