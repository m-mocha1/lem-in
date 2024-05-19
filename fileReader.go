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
			farmInfo.Ants, _ = strconv.Atoi(line)
			fmt.Println(farmInfo.Ants)
		}
		
		if strings.Contains(line, " ") {
			parts := strings.Fields(line)
			farmInfo.RoomName = Room{parts[0]}
			farmInfo.Location.X = parts[1]
			farmInfo.Location.Y = parts[2]
		}
	}
}
