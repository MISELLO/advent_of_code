package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type opt struct {
	L string
	R string
}

// Instr represent the instructions Left (L) and Right (R) that must be followed
var Instr string

// M represents the documents labeled "maps" where the key is node where you are
// and the value are the two options you can choose (L and R).
var M map[string]opt

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

	// First line is for the instructions
	fmt.Println("Loading instructions")
	scn.Scan()
	Instr = scn.Text()

	// The rest of the lines are for the maps
	fmt.Println("Loading map nodes")
	M = make(map[string]opt)
	for scn.Scan() {
		l := scn.Text()
		l = strings.ReplaceAll(l, "=", "")
		l = strings.ReplaceAll(l, "(", "")
		l = strings.ReplaceAll(l, ",", "")
		l = strings.ReplaceAll(l, ")", "")
		f := strings.Fields(l)
		if len(f) == 3 {
			M[f[0]] = opt{L: f[1], R: f[2]}
		}
	}
	fmt.Println("Map nodes loaded")

	// Calculating the steps to get from AAA to ZZZ
	fmt.Println("Calculating steps")
	var steps int = 0
	var node string = "AAA"
	var i int = 0
	for node != "ZZZ" {
		if Instr[i:i+1] == "L" {
			node = M[node].L
		} else { // Instr[i:i+1] == "R"
			node = M[node].R
		}
		steps++
		i++
		i = i % len(Instr)
	}

	// Result
	fmt.Println("We went from AAA to ZZZ in", steps, "steps.")
}
