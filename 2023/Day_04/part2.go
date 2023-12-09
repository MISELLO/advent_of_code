package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type card struct {
	nums1, nums2 []int
	multiplier   int
}

var games []card

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
	for scn.Scan() {
		l := scn.Text()

		halfs := strings.Split(l, ":")
		lists := strings.Split(halfs[1], "|")
		nums1 := strings.Fields(lists[0])
		nums2 := strings.Fields(lists[1])

		var c card
		c.multiplier = 1

		for i := 0; i < len(nums1); i++ {
			x, _ := strconv.Atoi(nums1[i])
			c.nums1 = append(c.nums1, x)
		}

		for i := 0; i < len(nums2); i++ {
			x, _ := strconv.Atoi(nums2[i])
			c.nums2 = append(c.nums2, x)
		}

		games = append(games, c)
	}
	fmt.Println("Cards loaded")
	fmt.Println("Processing prices")
	for i := 0; i < len(games); i++ {
		var match int
		for j := 0; j < len(games[i].nums1); j++ {
			for k := 0; k < len(games[i].nums2); k++ {
				if games[i].nums1[j] == games[i].nums2[k] {
					match++
				}
			}
		}
		fmt.Println("Card", i+1, "has", match, "matchings.")
		for l := 1; l <= match; l++ {
			games[i+l].multiplier += games[i].multiplier
		}
	}
	// Counting total cards
	var count int
	for i := 0; i < len(games); i++ {
		fmt.Println(games[i].multiplier, "instances of card", i+1)
		count += games[i].multiplier
	}
	fmt.Println("You get a total of", count, "cards.")
}
