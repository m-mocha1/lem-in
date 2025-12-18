# Overview
lem-in is a Go program that reads an ant-farm description from a file, builds a graph of rooms and tunnels, finds a quickest way to move all ants from ##start to ##end, and prints:


The original input content (ants, rooms, links)


The ants’ moves turn-by-turn in the format Lx-room


The program also validates input and prints:
ERROR: invalid data format
when the file is malformed or impossible to solve.

## Features


Parses rooms, tunnels, ##start, ##end, and ignores unknown commands/comments (# ...)


Builds a graph (adjacency list) from tunnels


Finds path(s) using a recursive depth-first search


Simulates ant movements while respecting rules:


One ant per room (except ##start and ##end)


Each tunnel used at most once per turn


Only moved ants are printed each turn





## Algorithm


Graph representation: rooms are nodes, tunnels are edges (undirected).


Search: recursive DFS is used to explore possible paths from start to end while avoiding cycles.


Path choice: among discovered valid paths, the program selects the quickest solution (fewest total turns) based on path lengths and traffic constraints.


Simulation: ants are dispatched across the chosen path(s) and moved turn-by-turn, outputting only the moves performed.



## Input Format


First non-comment line: number of ants (integer)


Rooms: name x y


Links: name1-name2


Special markers:


### start followed by the start room definition


### end followed by the end room definition




Room names never start with L or # and contain no spaces.

## Output Format
Print the full input content exactly, then a blank line, then moves:
L1-roomA L2-roomB ...
Only ants that moved that turn are shown.

## Error Handling
The program prints ERROR: invalid data format for cases such as (examples):


missing/invalid ant count


missing ##start or ##end


duplicate room names


invalid room coordinates


links to unknown rooms


duplicated tunnels


no path between start and end



## Usage
```bash
go run . <file>
```
### Example:
```bash
go run . test0.txt
```

## Project Structure (example)


main.go — argument handling + orchestration


parser/ — input parsing and validation


graph/ — graph structures and helpers


solver/ — recursive DFS + path selection


sim/ — ant movement simulation and printing



## Notes


Only standard Go packages are used.


Test files can be placed in a tests/ folder and used for manual or unit testing.

