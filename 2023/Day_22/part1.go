package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const groundLevel int = 1

type tPosition struct {
	x, y, z int
}

type tBlock struct {
	index                 int
	bHead, bTail          tPosition
	supportedBy, supports []int
	inPlace               bool
}

// blockList is the list of all blocks in order of appearance.
var blockList []*tBlock

// blockHeight gives you a list of indexes of blocks that are on the
// same height (the key is the max height that reaches the block).
var blockHeight map[int][]int

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
	fmt.Println("Loading and making fall the bricks ...")
	for scn.Scan() {
		l := scn.Text()
		var b tBlock
		b.New(l)
	}

	// Blocks are not in order, so we have to first store them and then make them fall
	// from the lowest one to the highest one
	for i := 0; i < len(blockList); i++ {
		b := GetLowerBlockNotInPlace()
		b.Fall()
	}

	// Result
	var res int
	for _, b := range blockList {
		if b.IsDesintegrable() {
			res++
		}
	}
	fmt.Printf("There are \033[1m%d\033[0m bricks that can be choosen as the one to get desintegrated.\n", res)
}

// New is the constructor of blocks.
// Initializes the block from a string and places it on the blockList slice.
func (b *tBlock) New(s string) {

	if blockHeight == nil {
		blockHeight = make(map[int][]int)
	}

	parts := strings.Split(s, "~")
	headCoord := strings.Split(parts[0], ",")
	b.bHead.x, _ = strconv.Atoi(headCoord[0])
	b.bHead.y, _ = strconv.Atoi(headCoord[1])
	b.bHead.z, _ = strconv.Atoi(headCoord[2])
	tailCoord := strings.Split(parts[1], ",")
	b.bTail.x, _ = strconv.Atoi(tailCoord[0])
	b.bTail.y, _ = strconv.Atoi(tailCoord[1])
	b.bTail.z, _ = strconv.Atoi(tailCoord[2])
	b.index = len(blockList)
	b.inPlace = false

	blockList = append(blockList, b)
}

// Fall makes the block to fall until it reaches the ground or another block
func (b *tBlock) Fall() {
	for !willCollide(b) {
		b.Down()
	}
	b.inPlace = true

	h := b.HighestZ()
	blockHeight[h] = append(blockHeight[h], b.index)
}

// willCollide returns true if the block will collide with the floor or
// another block if it goes down one more unit.
// If it will collide it updates the lists supports and supported by
// of the affected blocks (could be more than 2).
func willCollide(b *tBlock) bool {
	minZ := b.LowestZ()
	if minZ <= groundLevel {
		return true
	}

	itWill := false
	l := blockHeight[minZ-1]
	for _, id := range l {
		a := blockList[id]
		pointsA := a.Points()
		pointsB := b.Points()
		for _, pB := range pointsB {
			for _, pA := range pointsA {
				if pA.x == pB.x && pA.y == pB.y { // z will be the same
					b.Over(a)
					itWill = true
				}
			}
		}
	}

	return itWill
}

// LowestZ returns the lowest height (z) of the block
func (b *tBlock) LowestZ() int {
	if b.bHead.z < b.bTail.z {
		return b.bHead.z
	}
	return b.bTail.z
}

// HighestZ returns the highest height (z) of the block
func (b *tBlock) HighestZ() int {
	if b.bHead.z > b.bTail.z {
		return b.bHead.z
	}
	return b.bTail.z
}

// Down moves the block one unit down
func (b *tBlock) Down() {
	b.bHead.z--
	b.bTail.z--
}

// Points returns a slice of all the points that contain the block
// A block that starts at 0, 0, 0 and ends at 2, 0, 0 will return
// {0, 0, 0}, {1, 0, ,} and {2, 0, 0}
func (b *tBlock) Points() []tPosition {
	var l []tPosition
	// Note: The coordinates of the tail are always higher thant the head
	// ie.: Head = {a, b, c}; Tail = {x, y, z}; x >= a, y >= b, z >= c
	for z := b.bHead.z; z <= b.bTail.z; z++ {
		for y := b.bHead.y; y <= b.bTail.y; y++ {
			for x := b.bHead.x; x <= b.bTail.x; x++ {
				l = append(l, tPosition{x, y, z})
			}
		}
	}
	return l
}

// Over puts the referenced block (b) over the one from parameters (a)
func (b *tBlock) Over(a *tBlock) {
	// First we check if (b) is already over (a)
	for _, i := range b.supportedBy {
		if i == a.index {
			return
		}
	}

	// Seems that (b) is not over (a) yet.
	// And, consequently, (a) does not suppor (b).
	b.supportedBy = append(b.supportedBy, a.index)
	a.supports = append(a.supports, b.index)
}

// IsDesintegrable returns true if after removing it, no other blocks will fall
func (b *tBlock) IsDesintegrable() bool {

	itIs := true

	// For each block it supports
	for _, x := range b.supports {
		if len(blockList[x].supportedBy) == 1 {
			// At least one supported block is only supported by our block
			// It is NOT Desintegrable
			itIs = false
		}
	}

	return itIs
}

// GetLowerBlockNotInPlace returns the block that is not yet in place and is the lowest one
func GetLowerBlockNotInPlace() *tBlock {
	first := true
	var lowest *tBlock
	for _, b := range blockList {

		// Already in place
		if b.inPlace {
			continue
		}

		// First case
		if first {
			lowest = b
			first = false
		}

		// Regular cases
		if b.bHead.z < lowest.bHead.z {
			lowest = b
		}
	}

	return lowest
}
