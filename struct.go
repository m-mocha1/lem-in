package main

type antFarm struct {
	Ants     int
	Start    string
	End      string
	Paths    []string
	Rooms    []string
	RoomName Room
	Location Location
}

type Room struct {
	Name string
}

type Path struct {
	PathNum   int
	NumOfAnts int
}
type Location struct {
	X string
	Y string
}

type sarsor struct {
	Name string
}
type Move struct {
	X int
	Y int
}
