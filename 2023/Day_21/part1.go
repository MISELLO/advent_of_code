package main

import (
	"bufio"
	"fmt"
	"os"
)

// steps is the number of steps the elf has to perform.
const steps int = 64

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
	moveElf(steps)

	// Result
	fmt.Printf("With %d steps, the elf can reach \033[1m%d\033[0m garden plots.\n", steps, len(oldPos))
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
func moveElf(s int) {

	if s <= 0 {
		return
	}

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

	clear(oldPos)
	copyMap(oldPos, newPos)
	clear(newPos)
	moveElf(s - 1)
}

// valid returns true if the position (p) is inside bounds and not a rock (#)
func valid(p tPosition) bool {
	return p.x >= 0 && p.x < len(gardenMap[0]) && p.y >= 0 && p.y < len(gardenMap) && gardenMap[p.y][p.x] != '#'
}

// copyMap copies the content of "b" into "a"
// "a" must be empty
func copyMap(a, b map[tPosition]bool) {
	for p := range b {
		a[p] = true
	}
}
