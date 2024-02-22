package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	//"strings"
)

var schematic []string

type number struct {
	startX, startY, endX, endY int
	// Note: As numbers are always horizontal endX will always be the same as startX
}

var numberList []number

func (n *number) value() int {
	v, _ := strconv.Atoi(schematic[n.startX][n.startY : n.endY+1])
	return v
}

func (n *number) isAdjacentToSymbol() bool {
	var north, south, east, west string

	// north
	if n.startX > 0 {
		var i0, in int
		if n.startY > 0 {
			i0 = n.startY - 1
		} else {
			i0 = n.startY
		}
		if n.endY < len(schematic[n.startX])-1 {
			in = n.endY + 2
		} else {
			in = n.endY + 1
		}
		north = schematic[n.startX-1][i0:in]
	}

	// south
	if n.endX < len(schematic)-1 {
		var i0, in int
		if n.startY > 0 {
			i0 = n.startY - 1
		} else {
			i0 = n.startY
		}
		if n.endY < len(schematic[n.startX])-1 {
			in = n.endY + 2
		} else {
			in = n.endY + 1
		}
		south = schematic[n.startX+1][i0:in]
	}

	// east
	if n.endY < len(schematic[n.startX])-1 {
		east = schematic[n.startX][n.endY+1 : n.endY+2]
	}

	// west
	if n.startY > 0 {
		west = schematic[n.startX][n.startY-1 : n.startY]
	}

	// Check directions
	return checkDirections(north, south, east, west)
}

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

	// Schematic
	fmt.Println("Loading schematic")
	for scn.Scan() {
		l := scn.Text()

		schematic = append(schematic, l)
	}
	fmt.Println("Schematic loaded")

	// Numbers
	for i, v := range schematic {
		numFound := false
		var n number
		for j, w := range v {
			if w >= '0' && w <= '9' { // We have a digit
				if !numFound {
					n.startX = i
					n.startY = j
				}
				numFound = true
				//fmt.Println("(", i, j, ") ->", string(w))
			} else { // We have something that is not a digit
				if numFound {
					n.endX = i
					n.endY = j - 1
					numberList = append(numberList, n)
				}
				numFound = false
			}
			if j == len(v)-1 && numFound { // We reached the end we had a digit
				n.endX = i
				n.endY = j
				numberList = append(numberList, n)
			}
		}
	}
	var sum int
	//fmt.Println("Schematic:", schematic)
	fmt.Println("NumberList:")
	for _, n := range numberList {
		//n.print()
		fmt.Printf("^- %d (%v)\n", n.value(), n.isAdjacentToSymbol())
		if n.isAdjacentToSymbol() {
			sum += n.value()
		}
	}
	fmt.Println("The sum of the part numbers is", sum)
}

// checkDirections
func checkDirections(north, south, east, west string) bool {

	// Checking north
	for _, c := range north {
		if c != '.' {
			return true
		}
	}

	// Checking south
	for _, c := range south {
		if c != '.' {
			return true
		}
	}

	// Checking east
	for _, c := range east {
		if c != '.' {
			return true
		}
	}

	// Checking west
	for _, c := range west {
		if c != '.' {
			return true
		}
	}

	return false
}
