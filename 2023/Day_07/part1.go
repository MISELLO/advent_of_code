package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type hand struct {
	cards string
	bet   int
}

var handList []hand

// Hand strength
type handStr int

const (
	High_card handStr = iota
	One_pair
	Two_pair
	Three_of_a_kind
	Full_house
	Four_of_a_kind
	Five_of_a_kind
)

func computeStr(h hand) handStr {

	m := make(map[string]int)

	for i := 0; i < len(h.cards); i++ {
		m[h.cards[i:i+1]]++
	}
	if len(m) == 1 {
		// All cards are the same
		return Five_of_a_kind
	} else if len(m) == 2 {
		// Two types of cards, could be Four_of_a_kind or Full_house
		for _, n := range m {
			if n == 4 || n == 1 {
				return Four_of_a_kind
			} else {
				return Full_house
			}
		}
	} else if len(m) == 3 {
		// Three types of cards, could be Three_of_a_kind or Two_pair
		for _, n := range m {
			if n == 3 {
				return Three_of_a_kind
			}
			if n == 2 {
				return Two_pair
			}
		}
	} else if len(m) == 4 {
		// Four types of cards and 5 positions, ther must be One_pair
		return One_pair
	}
	// All cards are different
	return High_card
}

func handCmp(a, b hand) int {
	aStr, bStr := computeStr(a), computeStr(b)
	if aStr > bStr {
		return 1
	} else if aStr < bStr {
		return -1
	} else {
		i := 0
		for i < len(a.cards) && a.cards[i:i+1] == b.cards[i:i+1] {
			i++
		}
		if i == len(a.cards) {
			return 0
		}
		cardLabelList := "AKQJT98765432"
		aIndex := strings.Index(cardLabelList, a.cards[i:i+1])
		bIndex := strings.Index(cardLabelList, b.cards[i:i+1])
		if aIndex > bIndex {
			return -1
		} else if aIndex < bIndex {
			return 1
		}
	}
	return 0
}

func main() {

	// Open file
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

	// Read hands
	scn := bufio.NewScanner(f)

	fmt.Println("Loading hands")
	for scn.Scan() {
		l := scn.Text()

		var h hand
		parts := strings.Split(l, " ")
		h.cards = parts[0]
		h.bet, err = strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}

		handList = append(handList, h)
	}
	fmt.Println("Hands loaded")
	fmt.Println("")
	//fmt.Println("hands list:", handList)

	// Sort hands
	slices.SortFunc(handList, handCmp)
	//fmt.Println("hands list:", handList)

	// Compute winnings
	var total int
	for i := 0; i < len(handList); i++ {
		total += (i + 1) * handList[i].bet
	}
	fmt.Println("The total winning is ", total)
}
