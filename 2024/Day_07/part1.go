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
		panic(err)
	}
	defer f.Close()

	fmt.Println("File", os.Args[1], "opened")

	scn := bufio.NewScanner(f)
	sum := 0

	// Check te equations one by one
	fmt.Println("Checking equations ...")
	for scn.Scan() {
		eq := scn.Text()
		solStr, numStr, found := strings.Cut(eq, ": ")
		if !found {
			panic(fmt.Errorf("Separator \": \" not found on equation \"%s\"", eq))
		}

		sol, err := strconv.Atoi(solStr)
		if err != nil {
			panic(err)
		}

		numStrSlice := strings.Split(numStr, " ")

		if anySolution(sol, numStrSlice, []int{}) {
			fmt.Println(" Solution found on", numStr, "=", sol, "☺")
			sum += sol
		} else {
			fmt.Println(" No solution found on", numStr, "=", sol)
		}
	}

	// Result
	fmt.Printf("The total calibration result is \033[1m%d\033[0m.\n", sum)
}

// anySolution returns true if adding "+" and "*" between the numbers in ns
// we can find the result r
// This is done using the temporary results tmp
func anySolution(r int, ns []string, tmp []int) bool {

	// No initialized numbers slice, error.
	if ns == nil {
		panic(fmt.Errorf("Slice passed to anySolution is empty or nil"))
	}

	// No more numbers, check if one is the result
	if len(ns) == 0 {
		for _, x := range tmp {
			if x == r {
				return true
			}
		}
		return false
	}

	// Operate next number
	n, err := strconv.Atoi(ns[0])
	if err != nil {
		panic(err)
	}

	// We create the new tmp slice of results
	var newTmp []int
	if len(tmp) == 0 {
		newTmp = append(newTmp, n)
	} else {
		for _, x := range tmp {
			newX1 := x + n
			newX2 := x * n
			newTmp = append(newTmp, newX1)
			newTmp = append(newTmp, newX2)
		}
	}

	return anySolution(r, ns[1:], newTmp)

}
