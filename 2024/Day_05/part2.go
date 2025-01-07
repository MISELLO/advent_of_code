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

	rules := make(map[string]struct{})

	// Load rules
	fmt.Println("Loading rules ...")
	for scn.Scan() {
		l := scn.Text()
		// empty line means no more rules
		if l == "" {
			break
		}
		rules[l] = struct{}{}
	}

	//fmt.Println(rules)
	sum := 0

	// Process updates
	fmt.Println("Processing updates ...")
	for scn.Scan() {
		l := scn.Text()
		pages := strings.Split(l, ",")
		fmt.Println(pages)
		if !rightOrder(pages, rules) {
			//mid := pages[len(pages)/2]
			mid := middleOrdered(pages, rules)
			fmt.Println(" Corrected, new middle page", mid, "added.")
			con, err := strconv.Atoi(mid)
			if err != nil {
				panic(err)
			}
			sum += con
		}
	}

	// Result
	fmt.Printf("The sum of the middle page numbers of the corrected updates is \033[1m%d\033[0m.\n", sum)
}

// rightOrder checks if the slice s complies all rules found in r
// by checking if each element violates any rule
func rightOrder(s []string, r map[string]struct{}) bool {
	for i := 0; i < len(s); i++ {
		for j := 0; j < i; j++ {
			if _, ok := r[s[i]+"|"+s[j]]; ok {
				fmt.Printf(" Rule %s|%s violated!\n", s[i], s[j])
				return false
			}
		}
		for j := len(s) - 1; j > i; j-- {
			if _, ok := r[s[j]+"|"+s[i]]; ok {
				fmt.Printf(" Rule %s|%s violated!\n", s[j], s[i])
				return false
			}
		}
	}
	return true
}

// middleOrdered fixes the order of the slice of pages s according
// to the rules r and returns the middle page of the update (as a string)
func middleOrdered(s []string, r map[string]struct{}) string {
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s); j++ {
			if _, ok := r[s[i]+"|"+s[j]]; ok {
				s[i], s[j] = s[j], s[i]
			}
		}
		for j := len(s) - 1; j > i; j-- {
			if _, ok := r[s[j]+"|"+s[i]]; ok {
				s[i], s[j] = s[j], s[i]
			}
		}
	}
	return s[len(s)/2]
}
