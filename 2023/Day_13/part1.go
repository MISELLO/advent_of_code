package main

import (
	"bufio"
	"fmt"
	"os"
)

type pattern [][]byte

var notes []pattern

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
	var singleNote pattern
	for scn.Scan() {
		l := scn.Text()
		if l != "" {
			singleNote = append(singleNote, []byte(l))
		} else {
			notes = append(notes, singleNote)
			singleNote = [][]byte{}
		}
	}
	notes = append(notes, singleNote)
	fmt.Println(len(notes), "notes loaded.")

	// Checking for simetry
	fmt.Println("Checking for simetry...")
	var sum int
	for i := 0; i < len(notes); i++ {
		n := getVerticalSimetry(notes[i])
		if n == -1 {
			n = getHorizontalSimetry(notes[i])
			n = n * 100
		}
		sum += n
	}

	fmt.Printf("The sum of my notes is \033[1m%d\033[0m\n", sum)
}

// getVerticalSimetry returns the number of columns on the left of the
// vertical simetry. It returns -1 if there is no vertical simetry
func getVerticalSimetry(p pattern) int {
	for i := 1; i < len(p[0]); i++ {
		if trueVerticalSimetry(p, i-1, i) {
			return i
		}
	}
	return -1
}

// getHorizontalSimetry returns the number of rows on the top of the
// horizontal simetry. It returns -1 if there is no horizontal simetry
func getHorizontalSimetry(p pattern) int {
	for i := 1; i < len(p); i++ {
		if trueHorizontalSimetry(p, i-1, i) {
			return i
		}
	}
	return -1
}

// trueVerticalSimetry returns true if exists a vertical simetry
// between columns (a) and (b). Returns false otherwise
func trueVerticalSimetry(p pattern, a, b int) bool {
	if a < 0 || b >= len(p[0]) {
		return true
	}

	for i := 0; i < len(p); i++ {
		if p[i][a] != p[i][b] {
			return false
		}
	}

	return trueVerticalSimetry(p, a-1, b+1)
}

// trueHorizontalSimetry returns true if exists an horizontal simetry
// between rows (a) and (b). Returns false otherwise
func trueHorizontalSimetry(p pattern, a, b int) bool {
	if a < 0 || b >= len(p) {
		return true
	}

	for i := 0; i < len(p[0]); i++ {
		if p[a][i] != p[b][i] {
			return false
		}
	}

	return trueHorizontalSimetry(p, a-1, b+1)
}
