package main

import (
	"bufio"
	"fmt"
	"os"
)

type tPos struct {
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

	// Calculate
	fmt.Println(" Counting the \"X-MAS\"...")
	var appearances int
	for y := 1; y < len(wordSearchBoard)-1; y++ {
		for x := 1; x < len(wordSearchBoard[y])-1; x++ {
			if wordSearchBoard[y][x] != 'A' {
				continue
			}
			if !(wordSearchBoard[y-1][x-1] == 'M' && wordSearchBoard[y+1][x+1] == 'S' ||
				wordSearchBoard[y-1][x-1] == 'S' && wordSearchBoard[y+1][x+1] == 'M') {
				continue
			}
			if !(wordSearchBoard[y-1][x+1] == 'M' && wordSearchBoard[y+1][x-1] == 'S' ||
				wordSearchBoard[y-1][x+1] == 'S' && wordSearchBoard[y+1][x-1] == 'M') {
				continue
			}
			appearances++
		}
	}

	// Result
	fmt.Printf("The \"X\" shaped \"MAS\" appears \033[1m%d\033[0m times.\n", appearances)
}
