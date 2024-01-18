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

	// Load input
	fmt.Println("Calculating ...")
	var sum int
	var first bool = true
	for scn.Scan() {
		l := scn.Text()
		f := strings.Split(l, " ")
		n := strings.Split(f[1], ",")
		var tmp []int
		for i := 0; i < len(n); i++ {
			x, _ := strconv.Atoi(n[i])
			tmp = append(tmp, x)
		}
		numA := calculateArrangements([]byte(f[0]), tmp)
		sum += numA
		if first {
			fmt.Printf("%d", numA)
			first = false
		} else {
			fmt.Printf(" + %d", numA)
		}
	}
	fmt.Printf(" = \033[1m%d\033[0m.\n", sum)
}

// calculateArrangements calls itselve recursively and calculates how many
// possible arrangements are there in the given slice of springs (spr) that
// satisfy the records (rec)
func calculateArrangements(spr []byte, rec []int) int {
	//fmt.Println("")
	//fmt.Printf("%s --> %v %v\n", spr, counter(spr), valid(counter(spr), rec))

	// If the spr is not valid, we just return 0
	if !(valid(counter(spr), rec)) {
		return 0
	}

	// We search for the first '?'
	i := strings.Index(string(spr), "?")

	// The '?' is not found
	if i == -1 {
		if same(counter(spr), rec) {
			return 1
		} else {
			return 0
		}
	}

	// We create a copy of spr for each call.
	newSpr1 := make([]byte, len(spr))
	newSpr2 := make([]byte, len(spr))
	copy(newSpr1, spr)
	copy(newSpr2, spr)
	newSpr1[i] = '.'
	newSpr2[i] = '#'

	// Recursive call to find all combinations
	return calculateArrangements(newSpr1, rec) + calculateArrangements(newSpr2, rec)

}

// counter reads the slice of springs (spr) and returns a slice of
// ints as a report
func counter(spr []byte) []int {
	var r []int
	c := 0
	old := spr[0]
	if old == '?' {
		r = append(r, c)
		return r
	}
	if old == '#' {
		c++
	}
	for i := 1; i < len(spr); i++ {
		if spr[i] == '?' {
			if c > 0 {
				r = append(r, c)
			}
			return r
		} else if spr[i] == '.' {
			if old == '#' {
				r = append(r, c)
				c = 0
				old = spr[i]
			}
		} else if spr[i] == '#' {
			c++
			old = spr[i]
		}
	}
	if c > 0 {
		r = append(r, c)
	}
	return r
}

// valid returns true if the record (a) will be compatible with record (b)
func valid(a, b []int) bool {
	if len(a) == 0 {
		return true
	}
	if len(a) > len(b) {
		return false
	}

	for i := 0; i < len(a)-1; i++ {
		if a[i] != b[i] {
			return false
		}
	}

	i := len(a) - 1
	return a[i] <= b[i]
}

// same returns true if both int slices are equal
func same(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
