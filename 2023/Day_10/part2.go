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

	// We change the S from the original slice to it's corresponding pipe
	removeS(field, s)

	// We duplicate/clone the field
	// Note: Cloning a 1D slice is trivial, a 2D slice not so much.
	modField := [][]byte{}
	for i := 0; i < len(field); i++ {
		modField = append(modField, []byte{})
		modField[i] = append([]byte{}, field[i]...)
	}

	// We then fill the straight path with '+' and the corners with 'c'
	var pathTiles int
	s.Next(modField)
	markPath(modField, s)
	pathTiles++
	for modField[s.y][s.x] != 'S' && modField[s.y][s.x] != '+' && modField[s.y][s.x] != 'c' {
		s.Next(modField)
		markPath(modField, s)
		pathTiles++
	}

	// We then check for every cell if we are inside or outside the loop
	var insideLoop, outsideLoop int
	for j := 0; j < len(field); j++ {
		for i := 0; i < len(field[j]); i++ {
			if isInsideLoop(modField, field, i, j) {
				insideLoop++
				field[j][i] = 'I'
				modField[j][i] = 'I'
			} else {
				outsideLoop++
			}
		}
	}

	// Result
	fmt.Printf(" Path tiles: %d\n Outside the loop: %d\n Inside the loop: \033[1m%d\033[0m\n", pathTiles, outsideLoop-pathTiles, insideLoop)
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
	case 'S', '+':
		// Doesn't matter from where, this is the end and we stay here
		s.oldX, s.oldY = s.x, s.y
	}
}

// markPath sets the previous tile from the path as '+' or 'c' if it is
// a corner in order to count it later
func markPath(f [][]byte, s position) {
	if f[s.oldY][s.oldX] == '-' || f[s.oldY][s.oldX] == '|' {
		f[s.oldY][s.oldX] = '+'
	} else { // S seems to be always a corner but it could not be (PENDING, check S)
		f[s.oldY][s.oldX] = 'c'
	}
}

// isInsideLoop checks if a point is inside the loop using the
// Ray Casting algorithm
func isInsideLoop(f, g [][]byte, x, y int) bool {
	if f[y][x] == '+' || f[y][x] == 'c' { // On the path, it counts as outside
		return false
	}

	var count int
	var up, down bool = false, false
	for i := x; i >= 0; i-- {
		if f[y][i] == '+' && !(up || down) {
			count++
		}
		if f[y][i] == 'c' { // if we step on a border we don't count the +
			border(g[y][i], &count, &up, &down)
		}
	}
	return count%2 != 0
}

// used on the function "isInsideLoop" for when we setep on a border
func border(c byte, count *int, up, down *bool) {
	if c == 'J' {
		*up = !*up
	} else if c == '7' {
		*down = !*down
	} else if *up && c == 'L' {
		*up = !*up
	} else if *up && c == 'F' {
		*up = !*up
		*count++
	} else if *down && c == 'F' {
		*down = !*down
	} else if *down && c == 'L' {
		*down = !*down
		*count++
	}
}

// removeS changes the starting point S to it's corresponding pipe.
func removeS(f [][]byte, s position) {
	// Check if it is '|'
	if canGoN(f, &s) && canGoS(f, &s) {
		f[s.y][s.x] = '|'
		return
	}
	// Check if it is '-'
	if canGoE(f, &s) && canGoW(f, &s) {
		f[s.y][s.x] = '-'
		return
	}
	// Check if it is 'L'
	if canGoN(f, &s) && canGoE(f, &s) {
		f[s.y][s.x] = 'L'
		return
	}
	// Check if it is 'J'
	if canGoN(f, &s) && canGoW(f, &s) {
		f[s.y][s.x] = 'J'
		return
	}
	// Check if it is '7'
	if canGoS(f, &s) && canGoW(f, &s) {
		f[s.y][s.x] = '7'
		return
	}
	// Check if it is 'F'
	if canGoS(f, &s) && canGoE(f, &s) {
		f[s.y][s.x] = 'F'
		return
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
