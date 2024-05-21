package main

import "fmt"

func NewGraph() *Graph {
	return &Graph{Heads: []*Head{}}
}

func (g *Graph) addHead(k int) {
	g.Heads = append(g.Heads, &Head{Key: k})
}

func (g *Graph) findHead(key int) *Head {
	for _, head := range g.Heads {
		if head.Key == key {
			return head
		}
	}
	return nil
}
func (g *Graph) AddEdge(from, to int) {
	fromHead := g.findHead(from)
	ToHead := g.findHead(to)
	fromHead.Close = append(fromHead.Close, ToHead)
	//if u want 2 dirictons uncomment this
	// ToHead.Close = append(ToHead.Close, fromHead)
}

func (g *Graph) Print() {
	for _, v := range g.Heads {
		for _, t := range v.Close {
			fmt.Println("head --->", v.Key, "Tunnel to ", t.Key)
		}
	}
	fmt.Println()
}
