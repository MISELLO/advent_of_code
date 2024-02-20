package main

import (
	"bufio"
	"fmt"
	"os"
)

// steps is the number of steps the elf has to perform.
const steps int = 26501365

type tPosition struct {
	x, y int
}

var gardenMap [][]byte

// We will keep track of the positions the elf can go with this two maps.
var oldPos, newPos map[tPosition]bool

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
		gardenMap = append(gardenMap, []byte(l))
	}

	// Initialization
	p := findStartingPosition()
	oldPos = make(map[tPosition]bool)
	newPos = make(map[tPosition]bool)
	oldPos[p] = true

	// Starting
	fmt.Println("Doing the steps ...")
	oldPos = make(map[tPosition]bool)
	newPos = make(map[tPosition]bool)
	oldPos[p] = true

	// There is a pattern that is repeating on the input but not on the example.
	// Once the pattern is found, the number of plots grows as a parabola (quadratic function).
	// plots = ax² + bx + c (x is the number of steps)
	// In order to find the value of a, b and c we need 3 points.
	// This 3 points will be the plots at 65 steps [(len(gardenMap) - 1)/2],
	// At 196 steps (65 + len(gardenMap))
	// And at 327 (196 + len(gardenMap))
	values := make([]int, 3) // This will contain the number of plots at 65, 196 and 327 steps
	moveElf(steps, values)

	fmt.Println("At 65 steps the elf can reach", values[0], "garden plots.")
	fmt.Println("At 196 steps the elf can reach", values[1], "garden plots.")
	fmt.Println("At 327 steps the elf can reach", values[2], "garden plots.")
	fmt.Println("Finding the values of the formula ax² + bx + c")

	a := (values[2] - (2 * values[1]) + values[0]) / 2
	b := values[1] - values[0] - a
	c := values[0]

	// The function starts after reaching the first border (65 steps)
	// and is only valid every time it completes a full map (131 more steps)
	x := (steps - (len(gardenMap)-1)/2) / 131

	res := (a * (x * x)) + (b * x) + c

	// Result
	fmt.Printf("With %d steps, the elf can reach \033[1m%d\033[0m garden plots.\n", steps, res)
}

// findStartingPosition returns the position the elf has to start
func findStartingPosition() tPosition {
	for y := 0; y < len(gardenMap); y++ {
		for x := 0; x < len(gardenMap[y]); x++ {
			if gardenMap[y][x] == 'S' {
				return tPosition{x, y}
			}
		}
	}
	return tPosition{-1, -1}
}

// moveElf sets the porper positions on the variables oldPos and newPos
// until there are no more steps to do.
func moveElf(s int, v []int) {

	// Exit condition
	if s <= 0 {
		return
	}

	// The values we have to store (steps_done = total_steps - remaining_steps)
	if steps-s == (len(gardenMap)-1)/2 {
		v[0] = len(oldPos)
	} else if steps-s == (len(gardenMap)-1)/2+len(gardenMap) {
		v[1] = len(oldPos)
	} else if steps-s == (len(gardenMap)-1)/2+len(gardenMap)+len(gardenMap) {
		v[2] = len(oldPos)
		return
	}

	// Doing the new steps
	for p := range oldPos {
		var n tPosition

		// North
		n = tPosition{p.x, p.y - 1}
		if valid(n) {
			newPos[n] = true
		}

		// East
		n = tPosition{p.x + 1, p.y}
		if valid(n) {
			newPos[n] = true
		}

		// South
		n = tPosition{p.x, p.y + 1}
		if valid(n) {
			newPos[n] = true
		}

		// West
		n = tPosition{p.x - 1, p.y}
		if valid(n) {
			newPos[n] = true
		}
	}

	// Next steps
	clear(oldPos)
	copyMap(oldPos, newPos)
	clear(newPos)
	moveElf(s-1, v)
}

// valid returns true if the position (p) is inside bounds and not a rock (#)
func valid(p tPosition) bool {
	x, y := p.x, p.y
	for x < 0 {
		x += len(gardenMap[0])
	}
	for y < 0 {
		y += len(gardenMap)
	}
	x = x % len(gardenMap[0])
	y = y % len(gardenMap)
	return gardenMap[y][x] != '#'
}

// copyMap copies the content of "b" into "a"
// "a" must be empty
func copyMap(a, b map[tPosition]bool) {
	for p := range b {
		a[p] = true
	}
}
