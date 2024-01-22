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

	// Load a single input line
	scn.Scan()
	l := scn.Text()

	// Separate the commands
	steps := strings.Split(l, ",")

	// For each command get the hash
	var sum int
	for i := 0; i < len(steps); i++ {
		h := HolidayAsciiStringHelperAlgorithm(steps[i])
		fmt.Printf("- \"%s\" becomes %d.\n", steps[i], h)
		sum += h
	}

	// Result
	fmt.Printf("The sum of each hash is \033[1m%d\033[0m\n", sum)
}

// HolidayAsciiStringHelperAlgorithm (HASH for short) is a hash function that
// given a string (s) it returns an integer between 0-255
func HolidayAsciiStringHelperAlgorithm(s string) int {
	var cv int
	chars := []byte(s)
	for i := 0; i < len(chars); i++ {
		cv += int(chars[i])
		cv *= 17
		cv %= 256
	}
	return cv
}
