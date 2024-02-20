package main

import (
	"bufio"
	"fmt"
	"os"
)

type position struct {
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
	var contractedUni [][]byte
	fmt.Println("Loading map...")
	for scn.Scan() {
		l := scn.Text()
		contractedUni = append(contractedUni, []byte(l))
	}
	fmt.Println("Map loaded.")
	printUni(contractedUni)

	// We expand the universe
	expandedUni := expandUni(contractedUni)
	fmt.Println("Expanding the universe")
	printUni(expandedUni)

	// We search for galaxies
	fmt.Println("Searching for galaxies")
	var galaxies []position
	for y := 0; y < len(expandedUni); y++ {
		for x := 0; x < len(expandedUni[0]); x++ {
			if expandedUni[y][x] == '#' {
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
			d := dist(expandedUni, galaxies[i], galaxies[j])
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

// expandUni takes a given universe u and duplicates the columns and rows that are empty.
func expandUni(u [][]byte) [][]byte {
	var v [][]byte

	// Check rows
	for i := 0; i < len(u); i++ {
		tmp := make([]byte, len(u))
		copy(tmp, u[i])
		v = append(v, tmp)
		if !rowHasG(u[i]) {
			v = append(v, tmp) // We add it a second time
		}
	}

	// Check columns
	var cols []int
	for i := 0; i < len(u[0]); i++ {
		if !colHasG(v, i) {
			cols = append(cols, i)
		}
	}
	addDuplicateCols(v, cols)

	return v
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

// addDuplicatecols duplicates the columns of (u) that match with the indexes included in (c)
func addDuplicateCols(u [][]byte, c []int) {
	for j := 0; j < len(u); j++ { // Each row
		var k int
		var tmp []byte
		for i := 0; i < len(u[j]); i++ { // Each position
			tmp = append(tmp, u[j][i])
			if k < len(c) && i == c[k] { // We add it twice if the index matches the column we have to duplicate
				tmp = append(tmp, u[j][i])
				k++
			}
		}
		u[j] = tmp
	}
}

// dist returns the distance between the galaxies (g1) and (g2) in the universe (u)
func dist(u [][]byte, g1, g2 position) int {
	var difX, difY int
	if g1.x > g2.x {
		difX = g1.x - g2.x
	} else {
		difX = g2.x - g1.x
	}
	if g1.y > g2.y {
		difY = g1.y - g2.y
	} else {
		difY = g2.y - g1.y
	}
	return difX + difY
}
