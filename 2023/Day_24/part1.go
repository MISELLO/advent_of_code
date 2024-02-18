package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const testAreaXMin int = 200000000000000
const testAreaXMax int = 400000000000000
const testAreaYMin int = 200000000000000
const testAreaYMax int = 400000000000000

type tHailstone struct {
	px, py, pz, vx, vy, vz int
}

var hailList []tHailstone

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
	fmt.Println("Loading hailstones ...")
	for scn.Scan() {
		var h tHailstone
		l := scn.Text()
		l = strings.ReplaceAll(l, "  ", " ") // Example had a double space.
		// This turns every double space into a single space.
		part := strings.Split(l, " @ ")
		position := strings.Split(part[0], ", ")
		speed := strings.Split(part[1], ", ")
		h.px, _ = strconv.Atoi(position[0])
		h.py, _ = strconv.Atoi(position[1])
		h.pz, _ = strconv.Atoi(position[2])
		h.vx, _ = strconv.Atoi(speed[0])
		h.vy, _ = strconv.Atoi(speed[1])
		h.vz, _ = strconv.Atoi(speed[2])
		hailList = append(hailList, h)
	}

	// Calculate
	fmt.Println("Calculating collisions in work area ...")
	var res int
	for i := 0; i < len(hailList)-1; i++ {
		for j := i + 1; j < len(hailList); j++ {
			if CollisionInsideTestArea(hailList[i], hailList[j]) {
				res++
			}
		}
	}

	// Result
	fmt.Printf("In the test area there will be \033[1m%d\033[0m collisions.\n", res)
}

// CollisionInsideTestArea returns true if the hailstones (a) and (b) will collide
// inside the test area (ignoring the Z coordinate).
func CollisionInsideTestArea(a, b tHailstone) bool {
	//fmt.Println("Hailstone A:", a)
	//fmt.Println("Hailstone B:", b)

	// Line formula: Y = m X + b
	// Where m = vy / vx and b = Y - (m X)

	ma := float64(a.vy) / float64(a.vx)
	mb := float64(b.vy) / float64(b.vx)

	ba := float64(a.py) - (ma * float64(a.px))
	bb := float64(b.py) - (mb * float64(b.px))

	// Lines are parallel
	if ma == mb {
		//fmt.Println("Hailstones' paths are parallel; they never intersect.")
		return false
	}

	// Intersections
	intX := (bb - ba) / (ma - mb)
	intY := (intX * mb) + bb
	if intX < float64(testAreaXMin) || intX > float64(testAreaXMax) || intY < float64(testAreaYMin) || intY > float64(testAreaYMax) {
		// Out of Test Area
		//fmt.Printf("Hailstones' paths will cross outside the test area (at x=%0.3f, y=%0.3f).\n", intX, intY)
		return false
	}

	// Time
	Ta := (intX - float64(a.px)) / float64(a.vx)
	Tb := (intX - float64(b.px)) / float64(b.vx)
	if Ta < 0.0 || Tb < 0.0 {
		//fmt.Println("Hailstones' paths crossed in the past.")
		return false
	}

	//fmt.Printf("Formula for hailstone A is Y=%0.3fX + %0.3f\n", ma, ba)
	//fmt.Printf("Formula for hailstone B is Y=%0.3fX + %0.3f\n", mb, bb)
	//fmt.Printf("Trajectories intersect in the test area at: x=%0.3f, y=%0.3f\n", intX, intY)

	return true
}
