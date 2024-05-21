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
	Heads []*Head
}
type Head struct {
	Key   int
	Close []*Head
}
