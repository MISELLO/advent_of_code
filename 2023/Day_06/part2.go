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
	time, err := strconv.Atoi(strings.Join(timeList[1:], ""))
	if err != nil {
		panic(err)
	}

	// Distance
	fmt.Println("Loading distance list")
	scn.Scan()
	l = scn.Text()
	distanceList := strings.Fields(l)
	dist, err := strconv.Atoi(strings.Join(distanceList[1:], ""))
	if err != nil {
		panic(err)
	}

	fmt.Println("Time:    ", time)
	fmt.Println("Distance:", dist)

	// There is no need to check all combinations.
	// In the range of 1 to time-1 we will have 3 cases:
	//  1.- Pressed the button for too short a time
	//  2.- Pressed the button enough time
	//  3.- Pressed the button for too much time
	// We don't need to find all combinations 2 is satisfied,
	// we just need to find the range 2 is satisfied and
	// perform the difference.

	var lower, upper int

	// Lower
	for lower = 1; lower*(time-lower) < dist; lower++ {
	}

	// Upper
	for upper = time - 1; upper*(time-upper) < dist; upper-- {
	}

	fmt.Println("Lower:", lower, "Upper:", upper)

	fmt.Println("The number of ways I can beat it is", upper-lower+1)
}
