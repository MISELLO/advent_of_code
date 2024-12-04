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

	var l1, l2 []int

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

		l1 = addOrdered(l1, n1)
		l2 = addOrdered(l2, n2)
	}

	// Calculate
	fmt.Println("Adding the lowest numbers of each list ...")
	sum := 0

	for i := 0; i < len(l1); i++ {
		dif := dist(l1[i], l2[i])
		sum += dif
		fmt.Printf("|%d - %d| = %d (Sum = %d)\n", l1[i], l2[i], dif, sum)
	}

	// Result
	fmt.Printf("The sum of the distances is \033[1m%d\033[0m.\n", sum)
}

// addOrdered adds the integer n to the already ordered list l (lower numbers first)
func addOrdered(l []int, n int) []int {
	swap := func(i, j int) {
		l[i], l[j] = l[j], l[i]
	}

	l = append(l, n)
	for i := len(l) - 1; i > 0; i-- {
		if l[i] < l[i-1] {
			swap(i, i-1)
		}
	}

	return l
}

// dist returns the distance between a and b
func dist(a, b int) int {
	d := a - b
	if d < 0 {
		return -d
	}
	return d
}
