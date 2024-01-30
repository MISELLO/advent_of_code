package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type tInstruction struct {
	dir    string
	meters int
	color  string
}

var instrList []tInstruction

type tPosition struct {
	x, y int
}

// digMap is a representation of what surface is being excavated
// false means not excavated, true means excavated
var digMap [][]bool

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
	fmt.Println("Loading instructions ...")
	for scn.Scan() {
		l := scn.Text()
		m := strings.Split(l, " ")
		var instr tInstruction
		instr.dir = m[0]
		instr.meters, _ = strconv.Atoi(m[1])
		instr.color = m[2]
		instrList = append(instrList, instr)
	}
	fmt.Println("Instructions loaded.")
	fmt.Println("Generating map ...")
	start := initMap()
	fmt.Println("Starting position:", start)
	fmt.Println("Digging the edge of the lagoon...")
	digPerimeter(start)
	fmt.Println("Digging the interior of the lagoon...")

	// Note: In order to fill the resulting figure we will use the Flood Fill algorithm.
	// In order to work we are doing 3 asumptions:
	// 1.- The siluete has no narrow corridors like this:
	// ####....####
	// #..######..#
	// #..######..#
	// ####....####
	// 2.- The right-down side of the starting position is the inside of the perimeter
	// 3.- The algorithm will only fill inside the perimeter and, as a consequence of
	// that, it will never reach the border of the map.
	insidePos := tPosition{start.x + 1, start.y + 1}
	floodFill(insidePos)
	//printMap()

	// Result
	res := count()
	fmt.Printf("The lagoon will hold up to \033[1m%d\033[0m cubic meters of lava.\n", res)
}

// printMap is a temporary function that represents the digMap on the screen
func printMap() {
	for y := 0; y < len(digMap); y++ {
		for x := 0; x < len(digMap[0]); x++ {
			if digMap[y][x] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println("")
	}
}

// initMap initializes the digMap and returns the starting position
func initMap() tPosition {
	var p = tPosition{0, 0}
	var xMin, xMax, yMin, yMax int

	for i := 0; i < len(instrList); i++ {
		switch instrList[i].dir {
		case "U":
			p.y = p.y - instrList[i].meters
		case "R":
			p.x = p.x + instrList[i].meters
		case "D":
			p.y = p.y + instrList[i].meters
		case "L":
			p.x = p.x - instrList[i].meters
		}

		if p.x > xMax {
			xMax = p.x
		}
		if p.x < xMin {
			xMin = p.x
		}
		if p.y > yMax {
			yMax = p.y
		}
		if p.y < yMin {
			yMin = p.y
		}
	}

	width := xMax - xMin + 1
	height := yMax - yMin + 1

	fmt.Println("X goes from", xMin, "to", xMax, "Width:", width)
	fmt.Println("Y goes from", yMin, "to", yMax, "Height:", height)

	for i := 0; i < height; i++ {
		line := make([]bool, width)
		digMap = append(digMap, line)
	}

	return tPosition{p.x - xMin, p.y - yMin}
}

// digPerimeter follows the instructions and digs a trench that will be the perimeter of the lagoon
func digPerimeter(p tPosition) {
	for i := 0; i < len(instrList); i++ {
		switch instrList[i].dir {
		case "U":
			for j := p.y - instrList[i].meters; j < p.y; j++ {
				digMap[j][p.x] = true
			}
			p.y = p.y - instrList[i].meters
		case "R":
			for j := p.x + instrList[i].meters; j > p.x; j-- {
				digMap[p.y][j] = true
			}
			p.x = p.x + instrList[i].meters
		case "D":
			for j := p.y + instrList[i].meters; j > p.y; j-- {
				digMap[j][p.x] = true
			}
			p.y = p.y + instrList[i].meters
		case "L":
			for j := p.x - instrList[i].meters; j < p.x; j++ {
				digMap[p.y][j] = true
			}
			p.x = p.x - instrList[i].meters
		}
	}
}

// floodFill fills the content of a perimeter recursively
func floodFill(p tPosition) {
	if !digMap[p.y][p.x] {
		digMap[p.y][p.x] = true
		floodFill(tPosition{p.x, p.y - 1})
		floodFill(tPosition{p.x + 1, p.y})
		floodFill(tPosition{p.x, p.y + 1})
		floodFill(tPosition{p.x - 1, p.y})
	}
}

// count returns the number of m³ that have ben dig
func count() int {
	n := 0
	for y := 0; y < len(digMap); y++ {
		for x := 0; x < len(digMap[0]); x++ {
			if digMap[y][x] {
				n++
			}
		}
	}
	return n
}
