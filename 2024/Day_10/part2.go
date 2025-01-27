package main

import (
	"bufio"
	"fmt"
	"os"
)

type tTopographicMap [][]byte
type tPos struct {
	x, y int
}
type tTrailhead struct {
	init       tPos
	uniqueEnds map[tPos]int
	score      int
}

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

	// Load map (board)
	var tMap tTopographicMap
	fmt.Println("Loading map ...")
	for scn.Scan() {
		l := scn.Text()
		tMap = append(tMap, []byte(l))
	}

	// We find all the trailheads
	fmt.Println("Finding trailheads ...")
	var ts []tTrailhead
	ts = findTraiheads(tMap)
	fmt.Println("There are", len(ts), "trailheads.")

	// Calculate scores
	fmt.Println("Calculating the score of each trailhead ...")
	sum := 0
	for i, t := range ts {
		t.updateScore(tMap)
		fmt.Printf(" Trailhead %d/%d at position (%d, %d) has a score of %d.\n", i+1, len(ts), t.init.x, t.init.y, t.score)
		sum += t.score
	}

	// Result
	fmt.Printf("The sum of the scores of all trailheads is \033[1m%d\033[0m.\n", sum)
}

// findTrailheads returns a slice of all trailheads found on the topographic map m
func findTraiheads(m tTopographicMap) []tTrailhead {
	var ts []tTrailhead
	for y := 0; y < len(m); y++ {
		for x := 0; x < len(m[y]); x++ {
			if m[y][x] == '0' {
				p := tPos{
					x: x,
					y: y,
				}
				u := map[tPos]int{}
				t := tTrailhead{
					init:       p,
					uniqueEnds: u,
					score:      0,
				}
				ts = append(ts, t)
			}
		}
	}
	return ts
}

// updateScore sets the proper score to the trailhead t according to the topographic map m
func (t *tTrailhead) updateScore(m tTopographicMap) {
	_ = recCalculation(m, *t, t.init)
	for _, n := range t.uniqueEnds {
		t.score += n
	}
}

// recCalculation or recursive calculation explores the map m starting from the point p
// and returns the number of different ends that can be reached by the trailhead t.
func recCalculation(m tTopographicMap, t tTrailhead, p tPos) int {

	// End condition
	if m[p.y][p.x] == '9' {
		if t.uniqueEnds[p] > 0 {
			// We already have this end
			t.uniqueEnds[p]++
			return 0
		}
		// New unique end
		t.uniqueEnds[p] = 1
		return 1
	}

	nextHeight := m[p.y][p.x] + 1
	north, east, south, west := 0, 0, 0, 0
	var nextP tPos

	// North
	nextP = tPos{
		x: p.x,
		y: p.y - 1,
	}
	if nextP.isInside(m) && m[nextP.y][nextP.x] == nextHeight {
		north = recCalculation(m, t, nextP)
	}

	// East
	nextP = tPos{
		x: p.x + 1,
		y: p.y,
	}
	if nextP.isInside(m) && m[nextP.y][nextP.x] == nextHeight {
		east = recCalculation(m, t, nextP)
	}

	// South
	nextP = tPos{
		x: p.x,
		y: p.y + 1,
	}
	if nextP.isInside(m) && m[nextP.y][nextP.x] == nextHeight {
		south = recCalculation(m, t, nextP)
	}

	// West
	nextP = tPos{
		x: p.x - 1,
		y: p.y,
	}
	if nextP.isInside(m) && m[nextP.y][nextP.x] == nextHeight {
		west = recCalculation(m, t, nextP)
	}

	return north + east + south + west
}

// Returns true if the position p is in the map m
func (p tPos) isInside(m tTopographicMap) bool {
	return p.x >= 0 && p.y >= 0 && p.x < len(m[0]) && p.y < len(m)
}
