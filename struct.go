package main

type antFarm struct {
	Ants    int
	Start   string
	End     string
	Rooms   []Rooms
	tunnels []tunnels
}

type Rooms struct {
	Name     string
	Location Locations
}

type tunnels struct {
	From string
	To   string
}

type Locations struct {
	X int
	Y int
}

type Graph struct {
	Nodes map[string]*Nodes
}
type Nodes struct {
	Name  string
	Close []*Nodes
}
type Ant struct {
	Room string
	Id   int
}
