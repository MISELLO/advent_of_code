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
		n := getVerticalSmudgeSimetry(notes[i])
		if n == -1 {
			n = getHorizontalSmudgeSimetry(notes[i])
			n = n * 100
		}
		sum += n
	}

	fmt.Printf("The sum of my notes is \033[1m%d\033[0m\n", sum)
}

// getVerticalSmudgeSimetry returns the number of columns on the left of the
// vertical simetry having exactly 1 difference. It returns -1 if there is no
// vertical simetry with exactly 1 difference.
func getVerticalSmudgeSimetry(p pattern) int {
	for i := 1; i < len(p[0]); i++ {
		if smudgeVerticalSimetry(p, i-1, i, 0) {
			return i
		}
	}
	return -1
}

// getHorizontalSmudgeSimetry returns the number of rows on the top of the
// horizontal simetry having exactly 1 difference. It returns -1 if there is
// no horizontal simetry with exactly 1 difference.
func getHorizontalSmudgeSimetry(p pattern) int {
	for i := 1; i < len(p); i++ {
		if smudgeHorizontalSimetry(p, i-1, i, 0) {
			return i
		}
	}
	return -1
}

// smudgeVerticalSimetry returns true if exists a vertical simetry with only
// one difference between columns (a) and (b). Returns false otherwise
func smudgeVerticalSimetry(p pattern, a, b, c int) bool {
	if a < 0 || b >= len(p[0]) {
		return c == 1
	}

	for i := 0; i < len(p); i++ {
		if p[i][a] != p[i][b] {
			c++
			if c > 1 {
				return false
			}
		}
	}

	return smudgeVerticalSimetry(p, a-1, b+1, c)
}

// smudgeHorizontalSimetry returns true if exists an horizontal simetry with only
// one difference between rows (a) and (b). Returns false otherwise
func smudgeHorizontalSimetry(p pattern, a, b, c int) bool {
	if a < 0 || b >= len(p) {
		return c == 1
	}

	for i := 0; i < len(p[0]); i++ {
		if p[a][i] != p[b][i] {
			c++
			if c > 1 {
				return false
			}
		}
	}

	return smudgeHorizontalSimetry(p, a-1, b+1, c)
}
