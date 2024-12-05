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
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer f.Close()

	fmt.Println("File", os.Args[1], "opened")

	scn := bufio.NewScanner(f)

	var l1 []int
	// Key is an element of the right list
	// value is the number of times it appears
	occurrences := make(map[int]int)

	// Load input
	fmt.Println("Loading components ...")
	for scn.Scan() {
		l := scn.Text()
		parts := strings.Split(l, "   ")

		n1, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		n2, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		l1 = append(l1, n1)
		//fmt.Println(l1)
		occurrences[n2]++
		//fmt.Println(occurrences)
	}

	// Calculate
	fmt.Println("Adding the lowest numbers of each list ...")
	sum := 0

	for i := 0; i < len(l1); i++ {
		sum += l1[i] * occurrences[l1[i]]
		fmt.Printf("%d * %d = %d (Sum = %d)\n", l1[i], occurrences[l1[i]], l1[i]*occurrences[l1[i]], sum)
	}

	// Result
	fmt.Printf("The sum of the distances is \033[1m%d\033[0m.\n", sum)
}
