package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type tPosition struct {
	x, y int
}

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
	fmt.Println("Calculating ...")

	// We will use the Gauss's area formula (AKA Shoelace formula)
	// A = 1/2 | Sum from i=1 to N of (Xi * Yi+1 - Xi+1 * Yi)|
	// But this formula does not count half of the area of the
	// trench and the initial position. So, we have to add:
	// 1/2 * Perimeter + 1

	var perimeter, gauss int
	var p, oldP tPosition
	p = tPosition{0, 0}

	for scn.Scan() {
		l := scn.Text()
		m := strings.Split(l, " ")

		// Convert the meters in Hex to int64
		meters, _ := strconv.ParseInt(m[2][2:len(m[2])-2], 16, 64)
		perimeter += int(meters)

		// Direction
		dir := m[2][len(m[2])-2 : len(m[2])-1] // Last digit

		oldP = p
		switch dir {
		case "0": // Right
			p.x = p.x + int(meters)
		case "1": // Down
			p.y = p.y + int(meters)
		case "2": // Left
			p.x = p.x - int(meters)
		case "3": // Up
			p.y = p.y - int(meters)
		}
		gauss += (oldP.x * p.y) - (p.x * oldP.y)
	}

	gauss = gauss / 2
	fmt.Println("The result of applying the Gauss area formula (AKA Shoelace algorithm) is", gauss)
	perimeter = perimeter/2 + 1
	fmt.Println("The missing area of the perimeter is", perimeter)

	// Result
	res := gauss + perimeter
	fmt.Printf("The lagoon will hold up to \033[1m%d\033[0m cubic meters of lava.\n", res)
}
