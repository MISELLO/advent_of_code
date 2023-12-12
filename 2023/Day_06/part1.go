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
	var l string

	// Time
	fmt.Println("Loading time list")
	scn.Scan()
	l = scn.Text()
	timeList := strings.Fields(l)

	// Distance
	fmt.Println("Loading distance list")
	scn.Scan()
	l = scn.Text()
	distanceList := strings.Fields(l)

	fmt.Println("Time:    ", timeList[1:])
	fmt.Println("Distance:", distanceList[1:])

	if len(timeList) != len(distanceList) {
		fmt.Println("Time list and distance list must have the same lenght.")
		os.Exit(0)
	}

	var solList []int

	for i := 1; i < len(timeList); i++ {
		time, err := strconv.Atoi(timeList[i])
		if err != nil {
			panic(err)
		}

		dist, err := strconv.Atoi(distanceList[i])
		if err != nil {
			panic(err)
		}

		var numSol int

		for t := 1; t < time; t++ {
			d := t * (time - t)
			if d > dist {
				numSol++
			}
		}
		solList = append(solList, numSol)
	}

	var numWays int = solList[0]
	for i := 1; i < len(solList); i++ {
		numWays *= solList[i]
	}

	fmt.Println("The solution list is", solList, "and the number of ways I can beat it is", numWays)
}
