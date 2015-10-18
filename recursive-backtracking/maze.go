package main

import (
	"fmt"
	"math/rand"
	"time"
)

// north | south | east | west
type direction int

// Constants to signify which walls of a maze cell have been removed.
const (
	north = 1 << iota
	east
	south
	west
)

// Maps directions to Δx
var dx = map[direction]int{
	north: 0,
	east:  1,
	south: 0,
	west:  -1,
}

// Maps directions to Δy
var dy = map[direction]int{
	north: -1,
	east:  0,
	south: 1,
	west:  0,
}

// Opposite directions
var Opposite = map[direction]direction{
	north: south,
	east:  west,
	south: north,
	west:  east,
}

// Cell is a single position in a Maze.
type Cell int

// Maze of N x M dimensions
type Maze struct {
	width  int
	height int
	cells  [][]Cell
}

// NewMaze creates a new width x height Maze
func NewMaze(width, height int) *Maze {
	m := Maze{
		width,
		height,
		nil,
	}
	m.cells = make([][]Cell, height)
	for i := range m.cells {
		m.cells[i] = make([]Cell, width)
	}
	return &m
}

func (maze *Maze) generate() {
	rand.Seed(time.Now().UnixNano())
	maze.carvePassagesFrom(0, 0)
}

func between(x, min, max int) bool {
	return (x >= min && x <= max)
}

func shuffleDirections(slice []direction) {
	for i := range slice {
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
}

// carvePassagesFrom creates a maze starting from cell cx, cy
func (maze *Maze) carvePassagesFrom(cx, cy int) {
	var (
		d          direction
		directions = []direction{north, east, south, west}
	)
	shuffleDirections(directions)
	for i := range directions {
		d = directions[i]
		nx, ny := cx+dx[d], cy+dy[d]

		if between(nx, 0, maze.width-1) && between(ny, 0, maze.height-1) && maze.cells[ny][nx] == 0 {
			// Hacky cast to Cell so that we can encode the carved walls
			maze.cells[cy][cx] |= Cell(d)
			maze.cells[ny][nx] |= Cell(Opposite[d])
			maze.carvePassagesFrom(nx, ny)
		}
	}
}

func (maze *Maze) isExit(x, y int) bool {
	return x == maze.width-1 && y == maze.height-1
}

// Pretty prints the Maze
func (maze *Maze) print() {
	fmt.Print("  ") // 2 spaces, one for the left wall & one for the entrance
	for i := 0; i < maze.width*2-2; i++ {
		fmt.Print("_")
	}
	fmt.Print("\n")
	for y, row := range maze.cells {
		fmt.Print("|")
		for x, cell := range row {
			if cell&south != 0 || maze.isExit(x, y) {
				fmt.Print(" ")
			} else {
				fmt.Print("_")
			}
			if cell&east != 0 {
				if (cell|row[x+1])&south != 0 {
					fmt.Print(" ")
				} else {
					fmt.Print("_")
				}
			} else {
				fmt.Print("|")
			}
		}
		fmt.Print("\n")
	}
}
