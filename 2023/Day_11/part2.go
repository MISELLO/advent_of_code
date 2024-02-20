package main

import (
	"bufio"
	"fmt"
	"os"
)

type position struct {
	x, y int
}

const expander int = 1000000

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
	var Uni [][]byte
	fmt.Println("Loading map...")
	for scn.Scan() {
		l := scn.Text()
		Uni = append(Uni, []byte(l))
	}
	fmt.Println("Map loaded.")
	printUni(Uni)

	// We expand the universe
	expandUni(Uni)
	fmt.Println("Expanding the universe")
	printUni(Uni)

	// We search for galaxies
	fmt.Println("Searching for galaxies")
	var galaxies []position
	for y := 0; y < len(Uni); y++ {
		for x := 0; x < len(Uni[0]); x++ {
			if Uni[y][x] == '#' {
				galaxies = append(galaxies, position{x, y})
			}
		}
	}
	fmt.Println("I have found", len(galaxies), "galaxies:")
	fmt.Println(galaxies)

	// We get the length
	var sum int
	for i := 0; i < len(galaxies)-1; i++ {
		for j := i + 1; j < len(galaxies); j++ {
			d := dist(Uni, galaxies[i], galaxies[j])
			fmt.Printf("The distance between the galaxy %d and the galaxy %d is %d. (%d + %d = %d)\n", i, j, d, sum, d, sum+d)
			sum += d
		}
	}

	// Result
	fmt.Printf("The sum of the length of the shortest paths between every pair of galaxies is \033[1m%d\033[0m.\n", sum)
}

// printUni prints a representation of the universe
func printUni(u [][]byte) {
	for j := 0; j < len(u); j++ {
		for i := 0; i < len(u[j]); i++ {
			fmt.Printf("%c", u[j][i])
		}
		fmt.Println("")
	}
}

// expandUni takes a given universe u and marks the columns and rows that are empty.
// Then, each marked column or row will be counted as many times as necessary.
func expandUni(u [][]byte) {

	// Check rows
	for i := 0; i < len(u); i++ {
		if !rowHasG(u[i]) {
			markRow(u, i)
		}
	}

	// Check columns
	for i := 0; i < len(u[0]); i++ {
		if !colHasG(u, i) {
			markCol(u, i)
		}
	}
}

// rowHasG returns true if the row (r) contains a galaxy (#)
func rowHasG(r []byte) bool {
	for i := 0; i < len(r); i++ {
		if r[i] == '#' {
			return true
		}
	}
	return false
}

// colHasG returns true if the column (c) from the universe (u)
// contains a galaxy (#)
func colHasG(u [][]byte, c int) bool {
	for i := 0; i < len(u); i++ {
		if u[i][c] == '#' {
			return true
		}
	}
	return false
}

// markRow sets the row (i) from the universe (u) as '+'.
// This means it represents "expander" distance units instead of one.
func markRow(u [][]byte, i int) {
	for j := 0; j < len(u[0]); j++ {
		u[i][j] = '+'
	}
}

// markCol sets the column (i) from the universe (u) as '+'.
// This means it represents "expander" distance units instead of one.
func markCol(u [][]byte, i int) {
	for j := 0; j < len(u); j++ {
		u[j][i] = '+'
	}
}

// dist returns the distance between the galaxies (g1) and (g2) in the universe (u)
func dist(u [][]byte, g1, g2 position) int {
	var difX, difY int
	if g1.x > g2.x {
		for i := g1.x; i > g2.x; i-- {
			if u[g1.y][i] == '+' {
				difX = difX + expander
			} else {
				difX++
			}
		}
	}
	if g1.x < g2.x {
		for i := g1.x; i < g2.x; i++ {
			if u[g1.y][i] == '+' {
				difX = difX + expander
			} else {
				difX++
			}
		}
	}
	if g1.y > g2.y {
		for i := g1.y; i > g2.y; i-- {
			if u[i][g1.x] == '+' {
				difY = difY + expander
			} else {
				difY++
			}
		}
	}
	if g1.y < g2.y {
		for i := g1.y; i < g2.y; i++ {
			if u[i][g1.x] == '+' {
				difY = difY + expander
			} else {
				difY++
			}
		}
	}
	return difX + difY
}
