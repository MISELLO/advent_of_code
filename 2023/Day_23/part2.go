package main

import (
	"bufio"
	"fmt"
	"os"
)

type tPosition struct {
	x, y int
}

type tCursor struct {
	steps    int
	previous tPosition
	current  tPosition
	ckPoint  []tPosition
}

var startPos, endPos tPosition

var forestMap [][]byte

var maxLength int

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
	fmt.Println("Loading map ...")
	for scn.Scan() {
		l := scn.Text()
		forestMap = append(forestMap, []byte(l))
	}

	fmt.Println("Exploring the map ...")
	startPos = tPosition{1, 0}
	endPos = tPosition{len(forestMap[0]) - 2, len(forestMap) - 1}
	cursor := tCursor{0, startPos, startPos, []tPosition{}}
	walk(cursor)

	// Result
	fmt.Printf("The longest hike is \033[1m%d\033[0m steps.\n", maxLength)
}

// walk performs a walk into the forest labyrinth and prints if it founds a longer path
func walk(c tCursor) {

	// Initial step
	if c.current == startPos {
		c.current.y++
		c.steps++
		walk(c)
		return
	}

	// We have arrived
	if c.current == endPos {
		if c.steps > maxLength {
			maxLength = c.steps
			fmt.Println("Found a longer path:", maxLength)
		}
		return
	}

	// We check all possible directions
	pDir := possibleDirections(c)

	// In order to avoid infinite loops (or visit a place where we have been before)
	// we will store a check point if more than one direction is possible.
	// And check if we have been there before
	if len(pDir) > 1 {
		for _, p := range c.ckPoint {
			if c.current.x == p.x && c.current.y == p.y {
				// We have already been here, abort
				return
			}
		}
		// We haven't been here before, we store this check point
		c.ckPoint = append(c.ckPoint, c.current)
	}

	c.previous = c.current
	c.steps ++
	for _, p := range pDir {
		c.current = p
		walk(c)
	}
}

// possibleDirections returns a slice with all possible directions the cursor can go
func possibleDirections(c tCursor) []tPosition {
	var p []tPosition
	var newPos tPosition
	// North
	newPos = tPosition{c.current.x, c.current.y - 1}
	if newPos != c.previous && forestMap[newPos.y][newPos.x] != '#' {
		p = append(p, newPos)
	}
	// East
	newPos = tPosition{c.current.x + 1, c.current.y}
	if newPos != c.previous && forestMap[newPos.y][newPos.x] != '#' {
		p = append(p, newPos)
	}
	// South
	newPos = tPosition{c.current.x, c.current.y + 1}
	if newPos != c.previous && forestMap[newPos.y][newPos.x] != '#' {
		p = append(p, newPos)
	}
	// West
	newPos = tPosition{c.current.x - 1, c.current.y}
	if newPos != c.previous && forestMap[newPos.y][newPos.x] != '#' {
		p = append(p, newPos)
	}

	return p
}

