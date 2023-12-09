package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

	fmt.Println("Checking games")

	var sum int

	for scn.Scan() {
		l := scn.Text()

		// Game ID
		parts := strings.Split(l, ":")

		var minRed, minGreen, minBlue int

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
					if red > minRed {
						minRed = red
					}
				} else if f[1] == "green" {
					green, _ = strconv.Atoi(f[0])
					if green > minGreen {
						minGreen = green
					}
				} else { // blue
					blue, _ = strconv.Atoi(f[0])
					if blue > minBlue {
						minBlue = blue
					}
				}
			}
			fmt.Println("-> red:", red, "green:", green, "blue:", blue, "<-")
		}
		pow := minRed * minGreen * minBlue
		sum += pow
	}

	fmt.Println("Games checked")
	fmt.Println("The sum of the IDs is", sum)
}
