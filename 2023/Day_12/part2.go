package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var cache map[string]int

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

	cache = make(map[string]int)

	// Reading input
	fmt.Println("Calculating ...")
	var sum int
	var first bool = true
	for scn.Scan() {
		l := scn.Text()
		f := strings.Split(l, " ")
		n := strings.Split(f[1], ",")
		var nums []int
		for i := 0; i < len(n); i++ {
			x, _ := strconv.Atoi(n[i])
			nums = append(nums, x)
		}
		nums = expandI(nums)
		springs := expandB([]byte(f[0]))
		numA := calculateArrangements(springs, nums)
		sum += numA
		if first {
			fmt.Printf("%d", numA)
			first = false
		} else {
			fmt.Printf(" + %d", numA)
		}
	}
	fmt.Printf(" = \033[1m%d\033[0m.\n", sum)
	//fmt.Println(cache)
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

	// We get the key in order to know if we already know this solution
	key := getKey(spr, rec)
	result, exists := cache[key]
	//fmt.Printf("Key: %s, exists: %v, result: %d\n", key, exists, result)
	if exists {
		return result
	}

	// We create a copy of spr for each call.
	newSpr1 := make([]byte, len(spr))
	newSpr2 := make([]byte, len(spr))
	copy(newSpr1, spr)
	copy(newSpr2, spr)
	newSpr1[i] = '.'
	newSpr2[i] = '#'

	// Recursive call to find all combinations
	result = calculateArrangements(newSpr1, rec) + calculateArrangements(newSpr2, rec)
	//fmt.Println("spr:", string(spr), "key:", key, "result:", result)
	cache[key] = result
	return result
}

// counter reads the slice of springs (spr) and returns a slice of
// ints as a report
func counter(spr []byte) []int {
	var r []int
	c := 0
	old := spr[0]
	if old == '?' {
		return r
	}
	if old == '#' {
		c++
	}
	for i := 1; i < len(spr); i++ {
		if spr[i] == '?' {
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

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
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

// expandB returns a slice of bytes that contains 5 times
// the origial one separated by '?'
func expandB(s []byte) []byte {
	var r []byte
	r = append(r, s...)
	for i := 0; i < 4; i++ {
		r = append(r, '?')
		r = append(r, s...)
	}
	return r
}

// expandI returns a slice of integers that contains 5 times
// the original one.
func expandI(s []int) []int {
	var r []int
	for i := 0; i < 5; i++ {
		r = append(r, s...)
	}
	return r
}

// getKey gets a valid state that still has '?' in it
// returns a string that represents this state
func getKey(spr []byte, rec []int) string {

	i := strings.Index(string(spr), "#.")
	j := strings.Index(string(spr), "?")
	if i > -1 && j > i { // There is one group on the left
		return getKey(spr[i+2:], rec[1:])
	}

	s := string(spr)
	for i := 0; i < len(rec); i++ {
		s = s + "-" + strconv.Itoa(rec[i])
	}
	return s
}
