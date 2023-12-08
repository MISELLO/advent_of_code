package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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

	fmt.Println("Looking for digits")

	var total int

	for scn.Scan() {
		l := scn.Text()

		// Part 2: I had to rething it, we will work with this slice and search for the
		// first and last indexes that appear on the line.
		var digits = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
			"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

		var lowV, hightV int
		var lowI, hightI int = 9999, -1

		for i, v := range digits {
			x := strings.Index(l, v)
			y := strings.LastIndex(l, v)
			if x != -1 && x < lowI {
				lowI = x
				lowV = i % 10
			}
			if y > hightI {
				hightI = y
				hightV = i % 10
			}
		}

		num := lowV*10 + hightV
		fmt.Println(l, "=>", num)
		total += num
	}
	fmt.Println("Total is", total)
}
