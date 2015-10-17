package main

import (
	"fmt"
	"math/rand"
	"time"
)

// North | South | East | West
type direction int

// Constants to signify which walls of a maze cell have been removed.
const (
	North = 1 << iota
	East
	South
	West
)

// Maps directions to Δx
var dx = map[direction]int{
	North: 0,
	East:  1,
	South: 0,
	West:  -1,
}

// Maps directions to Δy
var dy = map[direction]int{
	North: -1,
	East:  0,
	South: 1,
	West:  0,
}

// Opposite directions
var Opposite = map[direction]direction{
	North: South,
	East:  West,
	South: North,
	West:  East,
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
func NewMaze(width int, height int) Maze {
	m := Maze{
		width,
		height,
		nil,
	}
	m.cells = make([][]Cell, height)
	for i := range m.cells {
		m.cells[i] = make([]Cell, width)
	}
	return m
}

func (maze *Maze) generate() {
	rand.Seed(time.Now().UnixNano())
	maze.carvePassagesFrom(0, 0)
}

func between(x int, min int, max int) bool {
	return (x >= min && x <= max)
}

func shuffleDirections(slice []direction) {
	for i := range slice {
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
}

// carvePassagesFrom creates a maze starting from cell cx, cy
func (maze *Maze) carvePassagesFrom(cx int, cy int) {
	var (
		d          direction
		directions = []direction{North, East, South, West}
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

func (maze *Maze) isExit(x int, y int) bool {
	return x == maze.width-1 && y == maze.height-1
}

// Pretty prints the Maze
func (maze *Maze) print() {
	fmt.Print("  ") // 2 spaces, one for the left wall & one for the entrance
	for i := 0; i < maze.width*2-2; i++ {
		fmt.Print("_")
	}
	fmt.Print("\n")
	for y := 0; y < maze.height; y++ {
		fmt.Print("|")
		for x := 0; x < maze.width; x++ {
			if maze.cells[y][x]&South != 0 || maze.isExit(x, y) {
				fmt.Print(" ")
			} else {
				fmt.Print("_")
			}
			if maze.cells[y][x]&East != 0 {
				if (maze.cells[y][x]|maze.cells[y][x+1])&South != 0 {
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
