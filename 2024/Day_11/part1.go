package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const BLINKS = 25

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

	// Load stones
	fmt.Println("Loading stones ...")
	scn.Scan()
	l := scn.Text()
	var stones []string
	stones = strings.Split(l, " ")
	fmt.Printf(" Originally we have %d stones: %v\n", len(stones), stones)

	// Blinking
	for b := range BLINKS {
		stones = process(stones)
		bl := "blink"
		if b > 0 {
			bl += "s"
		}
		// fmt.Printf(" After %d %s we have %d stones: %v\n", b+1, bl, len(stones), stones)
		fmt.Printf(" After %d %s we have %d stones.\n", b+1, bl, len(stones))
	}

	// Result
	fmt.Printf("\033[1mAnswer\033[0m: After %d blinks, there will be \033[1m%d\033[0m stones.\n", BLINKS, len(stones))
}

func process(s []string) []string {
	var result []string
	for _, r := range s {
		if r == "0" {
			result = append(result, "1")
		} else if len(r)%2 == 0 {
			left := r[:len(r)/2]
			right := r[len(r)/2:]
			right = strings.TrimLeft(right, "0")
			if right == "" {
				right = "0"
			}
			result = append(result, left, right)
		} else {
			n, err := strconv.Atoi(r)
			if err != nil {
				panic(err)
			}
			n *= 2024
			result = append(result, strconv.Itoa(n))
		}
	}
	return result
}
