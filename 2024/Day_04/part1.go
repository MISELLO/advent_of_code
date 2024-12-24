package main

import (
	"bufio"
	"fmt"
	"os"
)

type tPos struct {
	x, y int
}

// All possible directions
var North = tPos{x: +0, y: -1}
var NortEast = tPos{x: +1, y: -1}
var East = tPos{x: +1, y: +0}
var SouthEast = tPos{x: +1, y: +1}
var South = tPos{x: +0, y: +1}
var SouthWest = tPos{x: -1, y: +1}
var West = tPos{x: -1, y: +0}
var NorthWest = tPos{x: -1, y: -1}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please, provide just one file to analize.")
		os.Exit(0)
	}
	fmt.Println("Opening file", os.Args[1])

	f, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fmt.Println("File", os.Args[1], "opened")

	scn := bufio.NewScanner(f)

	var wordSearchBoard [][]byte

	// Load input
	fmt.Println("Loading word search board ...")
	for scn.Scan() {
		l := scn.Text()
		wordSearchBoard = append(wordSearchBoard, []byte(l))
	}

	// print(wordSearchBoard)

	// Calculate
	fmt.Println(" Counting the word \"XMAS\"...")
	var appearances int
	for y := 0; y < len(wordSearchBoard); y++ {
		for x := 0; x < len(wordSearchBoard[y]); x++ {
			appearances += countWordPos(wordSearchBoard, tPos{x: x, y: y}, North, "XMAS")
			appearances += countWordPos(wordSearchBoard, tPos{x: x, y: y}, NortEast, "XMAS")
			appearances += countWordPos(wordSearchBoard, tPos{x: x, y: y}, East, "XMAS")
			appearances += countWordPos(wordSearchBoard, tPos{x: x, y: y}, SouthEast, "XMAS")
			appearances += countWordPos(wordSearchBoard, tPos{x: x, y: y}, South, "XMAS")
			appearances += countWordPos(wordSearchBoard, tPos{x: x, y: y}, SouthWest, "XMAS")
			appearances += countWordPos(wordSearchBoard, tPos{x: x, y: y}, West, "XMAS")
			appearances += countWordPos(wordSearchBoard, tPos{x: x, y: y}, NorthWest, "XMAS")
		}
	}

	// Result
	fmt.Printf("The word \"XMAS\" appears \033[1m%d\033[0m times.\n", appearances)
}

// print just prints the board b just like the input file
func print(b [][]byte) {
	for i := 0; i < len(b); i++ {
		fmt.Println(string(b[i]))
	}
}

// isInBoard returns true if the position p is inside the board b
func isInBoard(p tPos, b [][]byte) bool {
	return p.x >= 0 && p.y >= 0 && p.x < len(b[0]) && p.y < len(b)
}

// countWordPos searches in the board b at position p and direction d the word w.
// It returns 1 if it finds the word w (it counts it)
// And returns 0 if it doesn't find it (does not count it)
func countWordPos(b [][]byte, p tPos, d tPos, w string) int {
	index := 0
	if b[p.y][p.x] != w[index] {
		return 0
	}
	for index = 1; index < len(w); index++ {
		p.x += d.x
		p.y += d.y
		if !isInBoard(p, b) || b[p.y][p.x] != w[index] {
			return 0
		}
	}
	return 1
}
