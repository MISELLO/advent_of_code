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

	fmt.Println("Moving files ...")
	blocks := representBlocs(dm)
	sortFiles(blocks)

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

// sortFiles fills the empty spaces on the left with the files on the right
func sortFiles(b []tBlock) {
	i := 0
	j := len(b) - 1
	movedFiles := make(map[int]bool)
	movedFiles[-1] = true // we count the free space as a moved file

	for i < j {
		fileId := b[j].id

		if movedFiles[fileId] {
			j = nextFile(b, j)
			continue
		}
		fileSize := fileSize(b, j)

		i, err := findSpot(b, fileSize, j)
		if err != nil {
			//fmt.Println(err.Error())
			movedFiles[fileId] = true // Only one attempt to move per file.
			continue
		}

		i, j = moveFile(b, i, j, fileSize)
		movedFiles[fileId] = true
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

// skipFile moves the index i to the next free space
func skipFile(b []tBlock, i int) int {
	for i < len(b) && b[i].id > -1 {
		i++
	}
	return i
}

// spacesAhead returns the number of spaces starting from position i
func spacesAhead(b []tBlock, i int) int {
	count := 0
	for i < len(b) && b[i].id == -1 {
		count++
		i++
	}
	return count
}

// i, err = findSpot(b, fileSize, j)
func findSpot(b []tBlock, size, stop int) (int, error) {
	for i := 0; i < stop; i++ {
		if spacesAhead(b, i) >= size {
			return i, nil
		}
	}
	return 0, fmt.Errorf("No spots found of size %d before index %d", size, stop)
}

// nextFile returns the file Id that comes next and the new j.
// This is the rightmost file on the left of j
func nextFile(b []tBlock, j int) int {
	oldId := b[j].id
	for j > 0 && b[j].id == oldId {
		j--
	}
	return j
}

// fileSize returns the size of the file that is on position j
// We assume j is the rightmost position of the file
func fileSize(b []tBlock, j int) int {
	id := b[j].id
	count := 0
	for j > 0 && b[j].id == id {
		count++
		j--
	}
	return count
}

// moveFile moves the file of size fs that ends at j to the starting position i
func moveFile(b []tBlock, i, j, fs int) (int, int) {
	for k := 0; k < fs; k++ {
		b[i], b[j] = b[j], b[i]
		i++
		j--
	}
	return i, j
}

// DEBUF function print
func print(b []tBlock) {
	for _, block := range b {
		if block.id < 0 {
			fmt.Printf("%s", ".")
		} else {
			fmt.Printf("%d", block.id)
		}
	}
	fmt.Println()
}
