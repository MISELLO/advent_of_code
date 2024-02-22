package main

import (
	"bufio"
	"fmt"
	"os"
)

type position struct {
	x, y, oldX, oldY int
}

var field [][]byte

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please, provide just one file to analize.")
		os.Exit(0)
	}
	fmt.Println("Opening file", os.Args[1])

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Could not open", os.Args[1])
		os.Exit(1)
	}
	defer f.Close()

	fmt.Println("File", os.Args[1], "opened")

	scn := bufio.NewScanner(f)

	// Load input
	fmt.Println("Loading map...")
	for scn.Scan() {
		l := scn.Text()
		field = append(field, []byte(l))
	}
	fmt.Println("Map loaded.")

	// We find the starting point
	s := findStartingPoint(field)
	fmt.Printf("Starting point is (%d, %d)\n", s.x, s.y)
	//fmt.Printf("Next should be (%d, %d) --> %c\n", s.x+1, s.y, field[s.y][s.x+1])

	// We count the steps
	s.Next(field)
	var steps int = 1
	for field[s.y][s.x] != 'S' {
		s.Next(field)
		steps++
	}

	// Result
	fmt.Printf("Steps to perform a loop: %d, \033[1m%d\033[0m to reach the furthest point.\n", steps, steps/2)
}

// findStartingPoint returns the position where to start
func findStartingPoint(f [][]byte) position {
	for j := 0; j < len(f); j++ {
		for i := 0; i < len(f[j]); i++ {
			if f[j][i] == 'S' {
				return position{i, j, i, j}
			}
		}
	}
	return position{-1, -1, -1, -1}
}

// Next is a method that makes the step further the previous position
// inside a given field (2D byte slice)
// In any direction if it is a starting point
func (s *position) Next(f [][]byte) {
	// Starting position
	if s.x == s.oldX && s.y == s.oldY {
		// North
		if canGoN(f, s) {
			s.y = s.y - 1
			return
		}
		// East
		if canGoE(f, s) {
			s.x = s.x + 1
			return
		}
		// South
		if canGoS(f, s) {
			s.y = s.y + 1
			return
		}
		// West
		if canGoW(f, s) {
			s.x = s.x - 1
			return
		}
	}

	// Middle and ending positions
	switch f[s.y][s.x] {
	case '|':
		dirNorthSouth(s)
	case '-':
		dirWestEast(s)
	case 'L':
		dirNorthEast(s)
	case 'J':
		dirNorthWest(s)
	case '7':
		dirSouthWest(s)
	case 'F':
		dirSouthEast(s)
	case 'S':
		// Doesn't matter from where, this is the end and we stay here
		s.oldX, s.oldY = s.x, s.y
	}
}

// isInBounds returns true if the point represented by x and y is
// inside the 2D slice f.
func isInBounds(f [][]byte, x, y int) bool {
	return x >= 0 && x < len(f[0]) && y >= 0 && y < len(f)
}

// canGoN returns true if we can go to the North
func canGoN(f [][]byte, p *position) bool {
	return isInBounds(f, p.x, p.y-1) && (f[p.y-1][p.x] == '|' || f[p.y-1][p.x] == '7' || f[p.y-1][p.x] == 'F')
}

// canGoE returns true if we can go to the East
func canGoE(f [][]byte, p *position) bool {
	return isInBounds(f, p.x+1, p.y) && (f[p.y+1][p.x] == '-' || f[p.y+1][p.x] == '7' || f[p.y+1][p.x] == 'J')
}

// canGoS returns true if we can go to the South
func canGoS(f [][]byte, p *position) bool {
	return isInBounds(f, p.x, p.y+1) && (f[p.y+1][p.x] == '|' || f[p.y+1][p.x] == 'L' || f[p.y+1][p.x] == 'J')
}

// cangoW returns true if we can go to the West
func canGoW(f [][]byte, p *position) bool {
	return isInBounds(f, p.x-1, p.y) && (f[p.y][p.x-1] == '-' || f[p.y][p.x-1] == 'L' || f[p.y][p.x-1] == 'F')
}

// Direction North South
func dirNorthSouth(s *position) {
	if s.y > s.oldY { // From North to South
		s.oldX, s.oldY = s.x, s.y
		s.y = s.y + 1
	} else { //          From South to North
		s.oldX, s.oldY = s.x, s.y
		s.y = s.y - 1
	}
}

// Direction West East
func dirWestEast(s *position) {
	if s.x > s.oldX { // From West to East
		s.oldX, s.oldY = s.x, s.y
		s.x = s.x + 1
	} else { //          From East to West
		s.oldX, s.oldY = s.x, s.y
		s.x = s.x - 1
	}
}

// Direction North East
func dirNorthEast(s *position) {
	if s.y > s.oldY { // From North to East
		s.oldX, s.oldY = s.x, s.y
		s.x = s.x + 1
	} else { //          From East to North
		s.oldX, s.oldY = s.x, s.y
		s.y = s.y - 1
	}
}

// Direction North West
func dirNorthWest(s *position) {
	if s.y > s.oldY { // From North to West
		s.oldX, s.oldY = s.x, s.y
		s.x = s.x - 1
	} else { //          From West to North
		s.oldX, s.oldY = s.x, s.y
		s.y = s.y - 1
	}
}

// Direction South West
func dirSouthWest(s *position) {
	if s.y < s.oldY { // From South to West
		s.oldX, s.oldY = s.x, s.y
		s.x = s.x - 1
	} else { //          From West to South
		s.oldX, s.oldY = s.x, s.y
		s.y = s.y + 1
	}
}

// Direction South East
func dirSouthEast(s *position) {
	if s.y < s.oldY { // From South to East
		s.oldX, s.oldY = s.x, s.y
		s.x = s.x + 1
	} else { //          From East to South
		s.oldX, s.oldY = s.x, s.y
		s.y = s.y + 1
	}
}
