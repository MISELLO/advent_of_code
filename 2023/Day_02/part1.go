package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const maxRed int = 12
const maxGreen int = 13
const maxBlue int = 14

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

	fmt.Println("Checking games")

	var sum int

	for scn.Scan() {
		l := scn.Text()

		// Game ID
		parts := strings.Split(l, ":")
		fields := strings.Fields(parts[0])
		gameID, _ := strconv.Atoi(fields[1])
		fmt.Println("Game ID:", gameID)
		var gameCorrect bool = true

		// Sets
		sets := strings.Split(parts[1], ";")
		for _, s := range sets {
			var red, green, blue int
			fmt.Println(" set -->", s)
			cubes := strings.Split(s, ",")
			for _, c := range cubes {
				f := strings.Fields(c)
				fmt.Println("  ", f[1], "=", f[0])
				if f[1] == "red" {
					red, _ = strconv.Atoi(f[0])
				} else if f[1] == "green" {
					green, _ = strconv.Atoi(f[0])
				} else { // blue
					blue, _ = strconv.Atoi(f[0])
				}
			}
			fmt.Println("-> red:", red, "green:", green, "blue:", blue, "<-")
			if red > maxRed || green > maxGreen || blue > maxBlue {
				gameCorrect = false
			}
		}
		if gameCorrect {
			sum += gameID
		}
	}

	fmt.Println("Games checked")
	fmt.Println("The sum of the IDs is", sum)
}
