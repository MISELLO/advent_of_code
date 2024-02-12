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
}

var startPos, endPos tPosition

var forestMap [][]byte

var pathLength []int

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

	cursor := tCursor{0, startPos, startPos}
	walk(cursor)

	// Result
	var res int
	for _, n := range pathLength {
		if n > res {
			res = n
		}
	}
	fmt.Printf("The longest hike is \033[1m%d\033[0m steps.\n", res)
}

// walk performs a walk into the forest labyrinth and stores into pathLength
// the number of steps it has taken to reach the end.
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
		pathLength = append(pathLength, c.steps)
		return
	}

	var newPos tPosition

	// If we are on a slope, we have to go to that direction
	if forestMap[c.current.y][c.current.x] == '^' { // slope North
		newPos = tPosition{c.current.x, c.current.y - 1}
		c.previous = c.current
		c.current = newPos
		c.steps++
		walk(c)
		return
	} else if forestMap[c.current.y][c.current.x] == '>' { // slope East
		newPos = tPosition{c.current.x + 1, c.current.y}
		c.previous = c.current
		c.current = newPos
		c.steps++
		walk(c)
		return
	} else if forestMap[c.current.y][c.current.x] == 'v' { // slope South
		newPos = tPosition{c.current.x, c.current.y + 1}
		c.previous = c.current
		c.current = newPos
		c.steps++
		walk(c)
		return
	} else if forestMap[c.current.y][c.current.x] == '<' { // slope West
		newPos = tPosition{c.current.x - 1, c.current.y}
		c.previous = c.current
		c.current = newPos
		c.steps++
		walk(c)
		return
	}

	// We step on any possible direction
	pDir := possibleDirections(c)

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
	if newPos != c.previous && forestMap[newPos.y][newPos.x] != '#' && forestMap[newPos.y][newPos.x] != 'v' {
		p = append(p, newPos)
	}
	// East
	newPos = tPosition{c.current.x + 1, c.current.y}
	if newPos != c.previous && forestMap[newPos.y][newPos.x] != '#' && forestMap[newPos.y][newPos.x] != '<' {
		p = append(p, newPos)
	}
	// South
	newPos = tPosition{c.current.x, c.current.y + 1}
	if newPos != c.previous && forestMap[newPos.y][newPos.x] != '#' && forestMap[newPos.y][newPos.x] != '^' {
		p = append(p, newPos)
	}
	// West
	newPos = tPosition{c.current.x - 1, c.current.y}
	if newPos != c.previous && forestMap[newPos.y][newPos.x] != '#' && forestMap[newPos.y][newPos.x] != '>' {
		p = append(p, newPos)
	}

	return p
}

