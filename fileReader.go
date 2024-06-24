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
	numAnts := 0
	for reader.Scan() {
		line := reader.Text()

		Fields := strings.Fields(line)
		if numAnts == 0 {
			numAnts, _ = strconv.Atoi(Fields[0])
			if numAnts > 0 {
				farmInfo.Ants = numAnts
			}
		}

		if line == "##start" {
			start = true
			continue
		}

		if line == "##end" {
			end = true
			continue
		}

		if strings.Contains(line, " ") {
			x, _ := strconv.Atoi(Fields[1])
			y, _ := strconv.Atoi(Fields[2])

			room := Rooms{
				Name:     Fields[0],
				Location: Locations{X: x, Y: y},
			}
			farmInfo.Rooms = append(farmInfo.Rooms, room)
			roomNum := Fields[0]
			g.addNode(roomNum)

			if start {
				farmInfo.Start = room.Name
				start = false
			} else if end {
				farmInfo.End = room.Name
				end = false
			}

		} else if strings.Contains(line, "-") {
			parts2 := strings.Split(line, "-")
			tunnel := tunnels{
				From: parts2[0],
				To:   parts2[1],
			}
			farmInfo.tunnels = append(farmInfo.tunnels, tunnel)
			from := tunnel.From
			to := tunnel.To
			g.AddEdge(from, to)
		}
	}
}
