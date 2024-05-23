package main

import "fmt"

func DFS(graph *Graph, startName, endName string) []string {
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
	if dfsVisit(start, endName, visited, &path) {
		return path
	}
	return []string{} // Return empty if no path found
}

func dfsVisit(node *Nodes, endName string, visited map[string]bool, path *[]string) bool {
	if visited[node.Name] {
		return false
	}

	visited[node.Name] = true
	*path = append(*path, node.Name)

	if node.Name == endName {
		return true
	}

	for _, neighbor := range node.Close {
		if dfsVisit(neighbor, endName, visited, path) {
			return true
		}
	}

	// If we reached here, then we need to backtrack
	*path = (*path)[:len(*path)-1]
	return false
}
