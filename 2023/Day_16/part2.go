package main

import (
	"bufio"
	"fmt"
	"os"
)

type tPosition struct {
	x, y, oldX, oldY int
}

var laserGrid [][]byte
var energizedGrid [][]bool
var dejaVu map[tPosition]bool

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
	fmt.Println("Loading ...")
	for scn.Scan() {
		l := scn.Text()
		laserGrid = append(laserGrid, []byte(l))
		energizedGrid = append(energizedGrid, make([]bool, len(l)))
	}
	fmt.Println("Grid loaded.")
	fmt.Println("Simulating the laser...")
	var max int
	dejaVu = make(map[tPosition]bool)
	l := AllPossibleStartingPositions()
	for i := 0; i < len(l); i++ {
		p := l[i]
		p.Run()
		num := EnergizedTiles()
		if num > max {
			max = num
		}
		clear(dejaVu)
	}

	// Result
	fmt.Printf("The maximum energized tiles we can find is \033[1m%d\033[0m\n", max)
}

// New creates a new position for the laser starting at coordinates (0, 0)
// and coming from the left (-1, 0)
func (p *tPosition) New() {
	p.oldX, p.oldY, p.x, p.y = -1, 0, 0, 0
}

// Energize makes sure the tile at position (p) is energized
func (p *tPosition) Energize() {
	energizedGrid[p.y][p.x] = true
}

// EnergizedTiles returns the number of tiles that are enrgized.
// And restores the energized grid back to it's default state.
func EnergizedTiles() int {
	n := 0
	for y := 0; y < len(energizedGrid); y++ {
		for x := 0; x < len(energizedGrid[y]); x++ {
			if energizedGrid[y][x] {
				n++
				energizedGrid[y][x] = false
			}
		}
	}
	return n
}

// Run makes the laser keeps going from position (p)
// If the position is not in the greed it just ends.
// If current tile is empty it keeps going.
// If there is a mirror it bounces.
// If there is a splitter it forks.
func (p *tPosition) Run() {

	// Check if out of bounds
	if !isInBounds(*p) {
		return
	}

	// Check if we have already been in this position with the same direction
	if dejaVu[*p] {
		return
	} else {
		dejaVu[*p] = true
	}

	// Energize current position
	p.Energize()

	// Performing next move
	if goingN(*p) {
		performN(p)
	} else if goingE(*p) {
		performE(p)
	} else if goingS(*p) {
		performS(p)
	} else if goingW(*p) {
		performW(p)
	}
}

// AllPossibleStartingPositions returns a slice of all possible ways to start the grid
func AllPossibleStartingPositions() []tPosition {
	var pos []tPosition

	// North and South
	b := len(laserGrid)
	for i := 0; i < len(laserGrid[0]); i++ {
		pos = append(pos, tPosition{i, 0, i, -1}, tPosition{i, b - 1, i, b})
	}

	// West and East
	r := len(laserGrid[0])
	for i := 0; i < len(laserGrid); i++ {
		pos = append(pos, tPosition{0, i, -1, i}, tPosition{r - 1, i, r, i})
	}

	return pos
}

// isInBounds returns true if the point (p) is
// inside the laserGrid.
func isInBounds(p tPosition) bool {
	return p.x >= 0 && p.x < len(laserGrid[0]) && p.y >= 0 && p.y < len(laserGrid)
}

// Going North
func goingN(p tPosition) bool {
	return p.x == p.oldX && p.y < p.oldY
}

// Going East
func goingE(p tPosition) bool {
	return p.x > p.oldX && p.y == p.oldY
}

// Going South
func goingS(p tPosition) bool {
	return p.x == p.oldX && p.y > p.oldY
}

// Going West
func goingW(p tPosition) bool {
	return p.x < p.oldX && p.y == p.oldY
}

// Perform North
func performN(p *tPosition) {
	c := laserGrid[p.y][p.x]
	p.oldX, p.oldY = p.x, p.y
	switch c {
	case '.':
		p.y = p.y - 1
		p.Run()
	case '/':
		p.x = p.x + 1
		p.Run()
	case '\\':
		p.x = p.x - 1
		p.Run()
	case '|':
		p.y = p.y - 1
		p.Run()
	case '-':
		q := *p
		p.x = p.x + 1
		q.x = q.x - 1
		p.Run()
		q.Run()
	}
}

// Perform East
func performE(p *tPosition) {
	c := laserGrid[p.y][p.x]
	p.oldX, p.oldY = p.x, p.y
	switch c {
	case '.':
		p.x = p.x + 1
		p.Run()
	case '/':
		p.y = p.y - 1
		p.Run()
	case '\\':
		p.y = p.y + 1
		p.Run()
	case '|':
		q := *p
		p.y = p.y - 1
		q.y = q.y + 1
		p.Run()
		q.Run()
	case '-':
		p.x = p.x + 1
		p.Run()
	}
}

// Perform South
func performS(p *tPosition) {
	c := laserGrid[p.y][p.x]
	p.oldX, p.oldY = p.x, p.y
	switch c {
	case '.':
		p.y = p.y + 1
		p.Run()
	case '/':
		p.x = p.x - 1
		p.Run()
	case '\\':
		p.x = p.x + 1
		p.Run()
	case '|':
		p.y = p.y + 1
		p.Run()
	case '-':
		q := *p
		p.x = p.x - 1
		q.x = q.x + 1
		p.Run()
		q.Run()
	}
}

// Perform West
func performW(p *tPosition) {
	c := laserGrid[p.y][p.x]
	p.oldX, p.oldY = p.x, p.y
	switch c {
	case '.':
		p.x = p.x - 1
		p.Run()
	case '/':
		p.y = p.y + 1
		p.Run()
	case '\\':
		p.y = p.y - 1
		p.Run()
	case '|':
		q := *p
		p.y = p.y - 1
		q.y = q.y + 1
		p.Run()
		q.Run()
	case '-':
		p.x = p.x - 1
		p.Run()
	}
}
