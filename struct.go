package main

type antFarm struct {
	numAnts int
	Start   string
	End     string
	Rooms   []*Rooms
	tunnels []Tunnels
	Ants    []Ant
	Paths   [][]string
}

type Rooms struct {
	Name       string
	occupiedBy *Ant
}

type Tunnels struct {
	From string
	To   string
}

type Graph struct {
	Nodes map[string]*Node
}
type Node struct {
	Name  string
	Close []*Node
}
type Ant struct {
	id          int
	path        []string
	pos         int
	currentRoom string
	End bool
}
type Path struct {
	Rooms     []Rooms
	AntsCount int
}
