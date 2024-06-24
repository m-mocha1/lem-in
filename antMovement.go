package main

// move the ants
func Move(g *Graph, ants []Ant, endNode string) {
	for i := range ants {
		if ants[i].Room == endNode {
			continue
			// ant reached the end check next ant
		}
		for i, jary := range g.Nodes[ants[i].Room].Close {
			if fady(jary.Name, ants) || jary.Name == endNode {
				ants[i].Room = jary.Name
				break // for moving the ant once each time
			}
		}
	}
}

//check if next room is emapty
func fady(room string, ants []Ant) bool {
	for _, ant := range ants {
		if ant.Room == room {
			return false
		}
	}
	return true
}

// if an ant reached the end return true
func End(ants []Ant, endNode string) bool {
	for _, ant := range ants {
		if ant.Room == endNode {
			return true
		}
	}
	return false
}
