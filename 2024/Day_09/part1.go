package main

import (
	"bufio"
	"fmt"
	"os"
)

// tBlock represents a block with an id.
// If the id is -1 then the block is free space
type tBlock struct {
	id int
}

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

	fmt.Println("Reading disk map ...")
	scn.Scan()
	dm := scn.Text()

	fmt.Println("Moving blocks ...")
	blocks := representBlocs(dm)
	sortBlocks(blocks)

	fmt.Println("Generating checksum ...")
	ck := checksum(blocks)

	// Result
	fmt.Printf("The resulting filesystem checksum is \033[1m%d\033[0m.\n", ck)
}

// representblocs translates a disk map into the blocks it represents.
// This is: File blocks, empty blocs repeatedly and the ID going from 0 to n
func representBlocs(dm string) []tBlock {
	var r []tBlock
	idCount := 0
	emptyBlock := false
	for _, p := range dm {
		n := int(p - '0')
		for i := 0; i < n; i++ {
			if emptyBlock {
				r = append(r, tBlock{-1})
			} else {
				r = append(r, tBlock{idCount})
			}
		}
		if !emptyBlock {
			idCount++
		}
		emptyBlock = !emptyBlock
	}
	return r
}

// sortBlocks fills the empty spaces on the left with the blocks on the right
func sortBlocks(b []tBlock) {
	i := 0
	j := len(b) - 1
	for i < j {
		for b[i].id > -1 && i < j {
			i++
		}
		for b[j].id == -1 && i < j {
			j--
		}
		if b[i].id == -1 && b[j].id != -1 {
			b[i], b[j] = b[j], b[i]
		}
	}
}

// checkum calculates the checksum of a slice of blocs by multiplying the position by the id
// and adding them all (ignoring the empty blocks)
func checksum(bs []tBlock) int {
	var sum int
	for p, b := range bs {
		if b.id < 0 {
			continue
		}
		sum += p * b.id
	}
	return sum
}
