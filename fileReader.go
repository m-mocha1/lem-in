package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var farmInfo antFarm

func fileReader(name string, g *Graph) {
	file, err := os.Open(name)
	if err != nil {
		fmt.Println("err", err)
	}
	reader := bufio.NewScanner(file)
	var start, end bool
	for reader.Scan() {
		line := reader.Text()

		if line == "##start" {
			start = true
			continue
		}

		if line == "##end" {
			end = true
			continue
		}

		if len(line) == 1 {
			numAnts, _ := strconv.Atoi(line)
			farmInfo.Ants = numAnts

		} else if strings.Contains(line, " ") {
			parts := strings.Fields(line)
			x, _ := strconv.Atoi(parts[1])
			y, _ := strconv.Atoi(parts[2])

			room := Rooms{
				Name:     parts[0],
				Location: Locations{X: x, Y: y},
			}
			farmInfo.Rooms = append(farmInfo.Rooms, room)
			roomNum, _ := strconv.Atoi(parts[0])
			g.addHead(roomNum)

			if start {
				farmInfo.Start = room.Name
				start = false
			} else if end {
				farmInfo.End = room.Name
				end = false
			}

		} else if strings.Contains(line, "-") {
			parts2 := strings.Split(line, "-")
			tunnels := tunnels{
				From: parts2[0],
				To:   parts2[1],
			}
			farmInfo.tunnels = append(farmInfo.tunnels, tunnels)
			from, _ := strconv.Atoi(tunnels.From)
			to, _ := strconv.Atoi(tunnels.To)
			g.AddEdge(from, to)
		}
	}
}
