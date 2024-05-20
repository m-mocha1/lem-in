package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var farmInfo antFarm

func fileReader(name string) {
	file, err := os.Open(name)
	if err != nil {
		fmt.Println("err", err)
	}
	reader := bufio.NewScanner(file)

	for reader.Scan() {
		line := reader.Text()

		if len(line) == 1 {
			numAnts, _ := strconv.Atoi(line)
			farmInfo.Ants = numAnts
		} else if strings.Contains(line, " ") {
			parts := strings.Fields(line)
			farmInfo.RoomName = parts[0]
			farmInfo.X = parts[1]
			farmInfo.Y = parts[2]
		}
	}
	fmt.Println(farmInfo)
}
