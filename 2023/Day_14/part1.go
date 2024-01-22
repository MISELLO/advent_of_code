package main

import (
	"bufio"
	"fmt"
	"os"
)

var platform [][]byte

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
	for scn.Scan() {
		l := scn.Text()
		platform = append(platform, []byte(l))
	}
	fmt.Println("Platform loaded.")
	//printPlatform()

	// Tild
	fmt.Println("Tilding the platform...")
	moveRocksNorth()
	//printPlatform()

	// Calculating load...
	fmt.Println("Calculating load...")
	var sum int
	multiplier := len(platform)
	for y := 0; y < len(platform); y++ {
		for x := 0; x < len(platform[y]); x++ {
			if platform[y][x] == 'O' {
				sum += multiplier
			}
		}
		multiplier--
	}

	fmt.Printf("The total load on the north support beams is \033[1m%d\033[0m\n", sum)
}

// printPlatform is a debug function that prints on the screen the current state of the platform
func printPlatform() {
	for i := 0; i < len(platform); i++ {
		fmt.Println(string(platform[i]))
	}
}

// moveRocksNorth moves all the rocks to north
func moveRocksNorth() {
	for y := 1; y < len(platform); y++ {
		for x := 0; x < len(platform[y]); x++ {
			if platform[y][x] == 'O' {
				displaceRockNorth(x, y)
			}
		}
	}
}

// displaceRockNorth moves a the rock at (x, y) to the North until it hits an obstacle
func displaceRockNorth(x, y int) {
	for y > 0 && platform[y-1][x] == '.' {
		platform[y][x], platform[y-1][x] = platform[y-1][x], platform[y][x]
		y--
	}
}
