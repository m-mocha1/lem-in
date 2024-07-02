package main

func RemoveBadPaths(paths [][]string) [][]string {

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
