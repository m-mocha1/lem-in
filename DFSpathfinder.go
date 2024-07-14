package main

import (
	"fmt"
)

func DFS(graph *Graph, startName, endName string) (pathes [][]string) {
	start := graph.findNode(startName)
	if start == nil {
		fmt.Println("start node not found")
		return nil
	}

	end := graph.findNode(endName)
	if end == nil {
		fmt.Println("end node not found")
		return nil
	}

	visited := make(map[string]bool)
	path := []string{}
	dfsVisit(start, endName, visited, path, &pathes)
	return pathes
}

func dfsVisit(node *Node, endName string, visited map[string]bool, path []string, pathes *[][]string) {
	if visited[node.Name] {
		return
	}

	visited[node.Name] = true
	path = append(path, node.Name)

	if node.Name == endName {
		*pathes = append(*pathes, path)
	} else {
		for _, jary := range node.Close {
			dfsVisit(jary, endName, visited, path, pathes)
		}
	}
	// if we reached here, then we need to backtrack

	visited[node.Name] = false

	path = path[:len(path)-1]
}
