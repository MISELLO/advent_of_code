package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
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

		var first, last rune
		var firstFound bool = false

		for i := 0; i < len(l); i++ {
			if unicode.IsDigit(rune(l[i])) {
				if !firstFound {
					first = rune(l[i])
					last = first
					firstFound = true
				} else {
					last = rune(l[i])
				}
			}
		}
		var s string = string(first) + string(last)
		num, _ := strconv.Atoi(s)
		fmt.Println(l, "=>", num)
		total += num
	}
	fmt.Println("Total is", total)
}
