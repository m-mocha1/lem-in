package main

import "fmt"

func NewGraph() *Graph {
	return &Graph{Node: []*Nodes{}}
}

func (g *Graph) addNode(k string) {
	g.Node = append(g.Node, &Nodes{Name: k})
}

func (g *Graph) findNode(Name string) *Nodes {
	for _, Node := range g.Node {
		if Node.Name == Name {
			return Node
		}
	}
	return nil
}
func (g *Graph) AddEdge(from, to string) {
	fromNode := g.findNode(from)
	ToNode := g.findNode(to)
	fromNode.Close = append(fromNode.Close, ToNode)
	//if u want 2 dirictons uncomment this
	// ToNode.Close = append(ToNode.Close, fromNode)
}

func (g *Graph) Print() {
	for _, v := range g.Node {
		for _, t := range v.Close {
			fmt.Println("Node-Name", v.Name, "Tunnel to ", t.Name, "Node", "\n")
		}
	}
	fmt.Println()
}
