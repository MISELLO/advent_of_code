package main

import (
	"bufio"
	"fmt"
	"os"
)

type tGardenPlotsMap [][]byte

type tPos struct {
	x, y int
}

var directions = []tPos{
	{0, -1}, {0, 1}, {1, 0}, {-1, 0},
}

type tVisited map[tPos]bool

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please, provide just one file to analize.")
		os.Exit(0)
	}
	fmt.Println("Opening file", os.Args[1], "...")

	f, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fmt.Println("File", os.Args[1], "opened")

	scn := bufio.NewScanner(f)

	// Load map (gardens)
	fmt.Println("Loading gardens ...")
	var gPlots tGardenPlotsMap
	for scn.Scan() {
		l := scn.Text()
		gPlots = append(gPlots, []byte(l))
	}

	// Processing ...
	visited := make(tVisited)
	sum := 0
	for y := 0; y < len(gPlots); y++ {
		for x := 0; x < len(gPlots[y]); x++ {
			p := tPos{x, y}

			if visited[p] {
				continue
			}

			plantType := gPlots[p.y][p.x]
			area, perimeter := floodFill(gPlots, visited, p, plantType)
			fmt.Printf(" - A region of %c plants with price %d x %d = %d\n", plantType, area, perimeter, area*perimeter)
			sum += area * perimeter
		}
	}

	// Result
	fmt.Printf("\033[1mAnswer\033[0m: The total price of fencing all regions is \033[1m%d\033[0m.\n", sum)
}

// floodFill checks a plot of plants and return its area and perimeter using the floot-fill algorithm
func floodFill(gPlots tGardenPlotsMap, visited tVisited, p tPos, plantType byte) (int, int) {

	// We abandon if go out the grid,
	// if we already visited that cell
	// or the plant is not the correct one
	if !inArea(gPlots, p) || visited[p] || gPlots[p.y][p.x] != plantType {
		return 0, 0
	}

	visited[p] = true

	area, prmt := 1, countFences(gPlots, p)

	for _, d := range directions {
		a, p := floodFill(gPlots, visited, tPos{x: p.x + d.x, y: p.y + d.y}, plantType)
		area += a
		prmt += p
	}

	return area, prmt
}

// countFences returns the anount of fences needed to complete the perimeter on this position
func countFences(gPlots tGardenPlotsMap, p tPos) int {
	plantType := gPlots[p.y][p.x]

	numFences := 0

	for _, d := range directions {
		newP := tPos{x: p.x + d.x, y: p.y + d.y}
		if !inArea(gPlots, newP) || gPlots[newP.y][newP.x] != plantType {
			numFences++
		}
	}

	return numFences
}

// inArea returns true if the point p is in the area a
func inArea(a tGardenPlotsMap, p tPos) bool {
	return p.x >= 0 && p.y >= 0 && p.y < len(a) && p.x < len(a[p.y])
}
