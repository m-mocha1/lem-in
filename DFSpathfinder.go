package main

import (
	"fmt"
	"sort"
)

func DFS(graph *Graph, startName, endName string) (pathes [][]string) {
	start := graph.findNode(startName)
	if start == nil {
		fmt.Println("Start node not found")
		return nil
	}

	end := graph.findNode(endName)
	if end == nil {
		fmt.Println("End node not found")
		return nil
	}

	visited := make(map[string]bool)
	path := []string{}
	dfsVisit(start, endName, visited, path, &pathes)
	return pathes
}

func dfsVisit(node *Nodes, endName string, visited map[string]bool, path []string, pathes *[][]string) {
	if visited[node.Name] {
		return
	}

	visited[node.Name] = true
	path = append(path, node.Name)

	if node.Name == endName {
		onePath := make([]string, len(path))
		copy(onePath, path)
		*pathes = append(*pathes, onePath)
	} else {
		for _, jary := range node.Close {
			dfsVisit(jary, endName, visited, path, pathes)

		}
	}
	// if we reached here, then we need to backtrack
	visited[node.Name] = false
	path = path[:len(path)-1]

}
func sorting(arr [][]string) {
	sort.Slice(arr, func(i, j int) bool {
		return len(arr[i]) < len(arr[j])
	})
}
