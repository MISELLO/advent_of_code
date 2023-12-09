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

	// Cards
	fmt.Println("Loading cards")
	var totalPoints int
	for scn.Scan() {
		l := scn.Text()

		var points int

		halfs := strings.Split(l, ":")
		lists := strings.Split(halfs[1], "|")
		nums1 := strings.Fields(lists[0])
		nums2 := strings.Fields(lists[1])

		for i := 0; i < len(nums1); i++ {
			x, _ := strconv.Atoi(nums1[i])
			for j := 0; j < len(nums2); j++ {
				y, _ := strconv.Atoi(nums2[j])
				if x == y {
					if points == 0 {
						points = 1
					} else {
						points = points * 2
					}
				}
			}
		}
		totalPoints += points
	}
	fmt.Println("Cards loaded")
	fmt.Println("All cards have a value of", totalPoints)
}
