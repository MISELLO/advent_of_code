package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type tOperation int

const (
	assign tOperation = iota
	remove
)

type tCommand struct {
	label string
	op    tOperation
	focal int
}

type tLens struct {
	label string
	focal int
}

var boxes [256][]tLens

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

	// We loop the commands and execute them
	for i := 0; i < len(steps); i++ {
		com := readCommand(steps[i])
		com.executeCommand()
	}
	//printBoxes()

	// Result
	var totalFocusPower int
	for i := 0; i < len(boxes); i++ {
		totalFocusPower += focusPower(i)
	}
	fmt.Printf("The focusing power of the resulting configuration is \033[1m%d\033[0m\n", totalFocusPower)
}

// printBoxes is a debug function to know the content of the boxes that are not empty.
func printBoxes() {
	for i := 0; i < len(boxes); i++ {
		if len(boxes[i]) != 0 {
			fmt.Println("Box", i, "-->", boxes[i])
		}
	}
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

// readCommand reads a string (s) and returns the parsed command
func readCommand(s string) tCommand {
	var c tCommand
	if strings.Contains(s, "=") {
		c.op = assign
		parts := strings.Split(s, "=")
		c.label = parts[0]
		c.focal, _ = strconv.Atoi(parts[1])
	} else { // Contains "-"
		c.op = remove
		c.label = strings.TrimSuffix(s, "-")
	}
	return c
}

// executeCommand performs the necessary change on the proper box from the 256 ones
func (c *tCommand) executeCommand() {
	index := HolidayAsciiStringHelperAlgorithm(c.label)
	b := boxes[index]
	if c.op == assign {
		found := false
		for i := 0; i < len(b); i++ {
			if b[i].label == c.label {
				b[i].focal = c.focal
				found = true
			}
		}
		if !found {
			b = append(b, tLens{c.label, c.focal})
		}
	} else { // c.op == remove
		var newB []tLens
		for i := 0; i < len(b); i++ {
			if b[i].label != c.label {
				newB = append(newB, b[i])
			}
		}
		b = newB
	}
	boxes[index] = b
}

// focusPower returns the focus power of the box (boxNum)
func focusPower(boxNum int) int {
	boxMultiplier := boxNum + 1
	boxPower := 0
	for i := 0; i < len(boxes[boxNum]); i++ {
		slotMultiplier := i + 1
		boxPower += boxes[boxNum][i].focal * slotMultiplier * boxMultiplier
	}
	return boxPower
}
