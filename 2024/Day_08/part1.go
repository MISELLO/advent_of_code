package main

import (
	"bufio"
	"fmt"
	"os"
)

type tPos struct {
	x, y int
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please, provide just one file to analize.")
		os.Exit(0)
	}
	fmt.Println("Opening file", os.Args[1])

	f, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fmt.Println("File", os.Args[1], "opened")

	scn := bufio.NewScanner(f)

	// Check the anthena map: We don't need to save it, we are just interested on the anthena name and position.
	fmt.Println("Checking the anthena map ...")
	anthenaPos := make(map[byte][]tPos)
	y := 0
	var maxX, maxY int
	for scn.Scan() {
		l := scn.Text()
		maxX = len(l)
		for x := 0; x < len(l); x++ {
			if l[x] != '.' {
				posList, _ := anthenaPos[l[x]]
				posList = append(posList, tPos{x, y})
				anthenaPos[l[x]] = posList
			}
		}
		y++
	}
	maxY = y

	// Generating antinodes
	fmt.Println("Generating antinodes ...")
	antiNodeMap := make(map[tPos]struct{})
	for a, list := range anthenaPos {
		for i := 0; i < len(list); i++ {
			for j := i + 1; j < len(list); j++ {
				var antiNode tPos
				antiNode = tPos{
					x: list[i].x + (list[i].x - list[j].x),
					y: list[i].y + (list[i].y - list[j].y),
				}
				if isInside(antiNode, maxX, maxY) {
					antiNodeMap[antiNode] = struct{}{}
				}

				// We do the same for the other anthena
				antiNode = tPos{
					x: list[j].x + (list[j].x - list[i].x),
					y: list[j].y + (list[j].y - list[i].y),
				}
				if isInside(antiNode, maxX, maxY) {
					antiNodeMap[antiNode] = struct{}{}
				}
			}
		}
		fmt.Println(" Antinodes type", string(a), "done.")
	}

	// Result
	fmt.Printf("There are \033[1m%d\033[0m unique anitnodes locations.\n", len(antiNodeMap))
}

// isInside returns true if the position p is inside the box formed by the points (0, 0) & (maxX, maxY)
func isInside(p tPos, maxX, maxY int) bool {
	return p.x >= 0 && p.y >= 0 && p.x < maxX && p.y < maxY
}
