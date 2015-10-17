// Implements a recursive backtracking maze generation algorithm.
// Based upon http://weblog.jamisbuck.org/2010/12/27/maze-generation-recursive-backtracking
package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const usage = "Usage: ./recursive-backtracking <width> <height>"

type dimensions struct {
	width  int
	height int
}

func main() {
	args := os.Args[1:]
	dim, err := parseArgs(args)
	if err != nil {
		fmt.Println(err)
		fmt.Println(usage)
		os.Exit(1)
	}
	maze := NewMaze(dim.width, dim.height)
	maze.generate()
	maze.print()
}

func parseArgs(args []string) (dimensions, error) {
	var (
		dim dimensions
		err error
	)

	if len(args) < 2 {
		return dim, errors.New("2 arguments required")
	}
	dim.width, err = strconv.Atoi(args[0])
	if err != nil {
		return dim, errors.New("Width must be an integer")
	}
	dim.height, err = strconv.Atoi(args[1])
	if err != nil {
		return dim, errors.New("Height must be an integer")
	}
	return dim, nil
}
