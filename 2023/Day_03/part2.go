package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	//"strings"
)

var schematic []string

var gearIDcount int = 1

type number struct {
	startX, startY, endX, endY int
	// Note: As numbers are always horizontal endX will always be the same as startX
	hasGear      bool
	gearX, gearY int
}

var numberList []number

func (n *number) print() {
	fmt.Println(schematic[n.startX][n.startY : n.endY+1])
}

func (n *number) value() int {
	v, _ := strconv.Atoi(schematic[n.startX][n.startY : n.endY+1])
	return v
}

func (n *number) updateGears() {
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

	fmt.Print("N:", north, "S:", south, "E:", east, "W:", west)
	n.hasGear = false

	// Checking north
	for i, c := range north {
		if c == '*' {
			n.hasGear = true
			fmt.Print(" -->Tn<--")
			n.gearX = n.startX - 1
			if n.startY > 0 {
				n.gearY = n.startY - 1 + i
			} else {
				n.gearY = n.startY + i
			}
		}
	}

	// Checking south
	for i, c := range south {
		if c == '*' {
			n.hasGear = true
			fmt.Print(" -->Ts<--")
			n.gearX = n.startX + 1
			if n.startY > 0 {
				n.gearY = n.startY - 1 + i
			} else {
				n.gearY = n.startY + i
			}
		}
	}

	// Checking east
	for _, c := range east {
		if c == '*' {
			n.hasGear = true
			fmt.Print(" -->Te<--")
			n.gearX = n.startX
			n.gearY = n.endY + 1
		}
	}

	// Checking west
	for _, c := range west {
		if c == '*' {
			n.hasGear = true
			fmt.Print(" -->Tw<--")
			n.gearX = n.startX
			n.gearY = n.startY - 1
		}
	}
	fmt.Println("")
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

	// Gears
	fmt.Println("Updating Gears")
	for i := 0; i < len(numberList); i++ {
		fmt.Print(numberList[i].value(), "* -")
		numberList[i].updateGears()
	}
	fmt.Println("Gears updated")

	var sum int
	fmt.Println("Getting gear pairs")
	for i := 0; i < len(numberList); i++ {
		if numberList[i].hasGear {
			fmt.Println(numberList[i].value(), numberList[i].hasGear, "(", numberList[i].gearX, ",", numberList[i].gearY, ")")
			gX := numberList[i].gearX
			gY := numberList[i].gearY
			for j := i + 1; j < len(numberList); j++ {
				if numberList[j].hasGear && numberList[j].gearX == gX && numberList[j].gearY == gY {
					// Same gear, different numbers
					fmt.Println(numberList[i].value(), "*", numberList[j].value())
					sum += numberList[i].value() * numberList[j].value()
				}
			}
		}
	}

	fmt.Println("The sum of the gear ratios is", sum)
}
