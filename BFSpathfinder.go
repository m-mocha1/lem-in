 package main

// import "fmt"

// func NewQueue() *Queue {
// 	return &Queue{}
// }

// func (q *Queue) Enque(h *Head) {
// 	q.nodes = append(q.nodes, h)
// }

// func (q *Queue) remove() *Head {
// 	if len(q.nodes) == 0 {
// 		return nil
// 	}
// 	head := q.nodes[0]
// 	q.nodes = q.nodes[1:]
// 	return head
// }

// func (q *Queue) empty() bool {
// 	return len(q.nodes) == 0
// }

// func BFS(g *Graph, startKey, endKey int) []int {
// 	start := g.findHead(startKey)
// 	end := g.findHead(endKey)

// 	if start == nil || end == nil {
// 		fmt.Println("start or end not found")
// 		return nil
// 	}

// 	queue := NewQueue()
// 	visited := make([]bool, len(g.Heads))
// 	prev := make([]int, len(g.Heads))
// 	for i := range prev {
// 		prev[i] = -1
// 	}

// 	queue.Enque(start)
// 	visited[start.Key] = true

// 	for !queue.empty() {
// 		node := queue.remove()

// 		if node.Key == end.Key {
// 			break //
// 		}

// 		for _, next := range node.Close {
// 			if !visited[next.Key] {
// 				queue.Enque(next)
// 				visited[next.Key] = true
// 				prev[next.Key] = node.Key
// 			}
// 		}
// 	}

// 	return reconstructPath(startKey, endKey, prev)
// }
// func reconstructPath(startKey, endKey int, prev []int) []int {
// 	path := []int{}
// 	for at := endKey; at != -1; at = prev[at] {
// 		path = append([]int{at}, path...)
// 	}
// 	if len(path) > 0 && path[0] == startKey {
// 		return path
// 	}
// 	return []int{}
// }
