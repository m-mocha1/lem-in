package main

import (
	"fmt"
	"time"
)

func RemoveBadPaths(paths [][]string) [][]string {
	start := time.Now()
	var newPaths [][]string

	for _, path := range paths {
		duplicate := false
		for _, existingPath := range newPaths {
			if checkDuplicates(path, existingPath) {
				duplicate = true
				break
			}
		}
		if !duplicate {
			newPaths = append(newPaths, path)
		}
	}
	end := time.Since(start)
	fmt.Println("the time it took to filter the pathes ", end)
	return newPaths
}
func checkDuplicates(path1, path2 []string) bool {
	for _, room1 := range path1 {
		if room1 == farmInfo.Start || room1 == farmInfo.End {
			continue
		}
		for _, room2 := range path2 {
			if room2 == farmInfo.Start || room2 == farmInfo.End {
				continue
			}
			if room1 == room2 {
				return true
			}
		}
	}
	return false
}
