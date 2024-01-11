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
	var result int

	// Processing lines
	for scn.Scan() {
		l := scn.Text()
		f := strings.Fields(l)
		var n []int
		for i := 0; i < len(f); i++ {
			x, _ := strconv.Atoi(f[i])
			n = append(n, x)
		}
		oldResult := result
		e := extrapolate(n)
		result = oldResult + e
		fmt.Println(oldResult, "+", e, "=", result)
	}

	// Result
	fmt.Println("The sum of the extrapolated values from the left is", result)
}

// extrapolate returns the extrapolated value of a given history result
func extrapolate(s []int) int {
	if allZeros(s) {
		return 0
	}

	var aux []int
	for i := 1; i < len(s); i++ {
		aux = append(aux, s[i]-s[i-1])
	}

	return s[0] - extrapolate(aux)
}

// allZeros returns true if all the values of the slice are 0
func allZeros(s []int) bool {
	for i := 0; i < len(s); i++ {
		if s[i] != 0 {
			return false
		}
	}
	return true
}
