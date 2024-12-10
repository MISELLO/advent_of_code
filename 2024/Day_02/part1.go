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

	var reports [][]int

	// Load input
	fmt.Println("Loading reports ...")
	for scn.Scan() {
		l := scn.Text()
		levels := strings.Split(l, " ")
		r := []int{}
		for i := 0; i < len(levels); i++ {
			n, err := strconv.Atoi(levels[i])
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
			r = append(r, n)
		}
		reports = append(reports, r)
	}

	//print(reports)

	// Calculate
	fmt.Println(" Counting the number of safe reports...")
	numSafe := 0
	for i := 0; i < len(reports); i++ {
		if isSafe(reports[i]) {
			numSafe++
		}
	}

	// Result
	fmt.Printf("We have \033[1m%d\033[0m safe reports.\n", numSafe)
}

func print(r [][]int) {
	for i := 0; i < len(r); i++ {
		fmt.Println(r[i], isSafe(r[i]))
	}
}

// isSafe returns true if a report is only ascending or only descending by a difference of 1, 2 or 3
func isSafe(r []int) bool {
	oldDif := 0
	for i := 1; i < len(r); i++ {
		dif := r[i] - r[i-1]
		if dif > 3 || dif < -3 || dif == 0 {
			return false
		}
		if oldDif < 0 && dif > 0 {
			return false
		}
		if oldDif > 0 && dif < 0 {
			return false
		}
		oldDif = dif
	}
	return true
}