package main

func NewGraph() *Graph {
	return &Graph{Nodes: make(map[string]*Node)}
}

func (g *Graph) addNode(k string) {
	g.Nodes[k] = &Node{Name: k}
}

func (g *Graph) findNode(Name string) *Node {
	for _, Node := range g.Nodes {
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
