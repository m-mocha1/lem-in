package main

import "fmt"

func DFS(graph *Graph, startKey, endKey string) []string {
	start := graph.findHead(startKey)
	if start == nil {
		fmt.Println("Start node not found")
		return nil
	}

	end := graph.findHead(endKey)
	if end == nil {
		fmt.Println("End node not found")
		return nil
	}

	visited := make(map[string]bool)
	path := []string{}
	if dfsVisit(start, endKey, visited, &path) {
		return path
	}
	return []string{} // Return empty if no path found
}

func dfsVisit(node *Head, endKey string, visited map[string]bool, path *[]string) bool {
	if visited[node.Key] {
		return false
	}

	visited[node.Key] = true
	*path = append(*path, node.Key)

	if node.Key == endKey {
		return true
	}

	for _, neighbor := range node.Close {
		if dfsVisit(neighbor, endKey, visited, path) {
			return true
		}
	}

	// If we reached here, then we need to backtrack
	*path = (*path)[:len(*path)-1]
	return false
}
