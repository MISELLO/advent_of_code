package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type tPiece struct {
	x, m, a, s int
}

type tRule struct {
	prop string // Property of the piece ("s", "m", "a" or "x")
	op   string // Operand ("<" or ">")
	val  int    // Value (an integer)
	dst  string // Destination ("A", "R" or the label to another workflow)
}

var workflows map[string][]tRule

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
	fmt.Println("Loading instructions ...")

	workflows = make(map[string][]tRule)
	workflowLoad := true
	var sum int

	for scn.Scan() {
		l := scn.Text()
		if workflowLoad && l != "" {
			// Load workflows
			label, ruleList := parseWorkflow(l)
			workflows[label] = ruleList
		} else if l == "" {
			// Empty line
			workflowLoad = false
			fmt.Println("Processing the pieces ...")
		} else {
			// Process pieces
			p := parsePiece(l)
			fmt.Printf(" - {x=%d, m=%d, a=%d, s=%d}: ", p.x, p.m, p.a, p.s)
			add := evaluatePiece(p, "in")
			fmt.Printf(" (%d + %d = %d)\n", sum, add, sum+add)
			sum += add
		}
	}

	// Result
	fmt.Printf("The sum of the ratings of all accepted pieces is \033[1m%d\033[0m\n", sum)
}

// parseWorkflow parses a string (s) and returns a label and a list of rules
// error is not an option
func parseWorkflow(s string) (string, []tRule) {
	var resultingList []tRule
	var tmp []string
	tmp = strings.Split(s, "{")
	label := tmp[0]
	rest := tmp[1]
	rest, _ = strings.CutSuffix(rest, "}")
	rulesRawList := strings.Split(rest, ",")
	for i := 0; i < len(rulesRawList); i++ {
		var r tRule
		tmp2 := strings.Split(rulesRawList[i], ":")
		r.dst = tmp2[len(tmp2)-1]
		r.op = "nop"
		if len(tmp2) > 1 {
			rest = tmp2[0]
			r.prop = rest[:1]
			r.op = rest[1:2]
			r.val, _ = strconv.Atoi(rest[2:])
		}
		resultingList = append(resultingList, r)
	}
	return label, resultingList
}

// parsePiece returns the resulting piece from the input string
// error is not an option
func parsePiece(s string) tPiece {
	var parts [4]int
	s, _ = strings.CutSuffix(s, "}")
	tmp := strings.Split(s, ",")
	for i := 0; i < len(tmp); i++ {
		tmp2 := strings.Split(tmp[i], "=")
		parts[i], _ = strconv.Atoi(tmp2[1])
	}
	return tPiece{x: parts[0], m: parts[1], a: parts[2], s: parts[3]}
}

// evaluatePiece applies the label (l) to the piece (p) and
// returns the sum of it's properties if it is accepted (A)
// or 0 if it's rejected (R)
func evaluatePiece(p tPiece, l string) int {
	fmt.Printf("%s --> ", l)
	w := workflows[l]
	var dst string
	for i := 0; i < len(w); i++ {
		var val int
		switch w[i].prop {
		case "x":
			val = p.x
		case "m":
			val = p.m
		case "a":
			val = p.a
		case "s":
			val = p.s
		default:
			dst = w[i].dst
		}
		if (w[i].op == "<" && val < w[i].val) || (w[i].op == ">" && val > w[i].val) || (w[i].op != "<" && w[i].op != ">") {
			dst = w[i].dst
			break
		}
	}

	switch dst {
	case "A":
		fmt.Printf("A")
		return p.x + p.m + p.a + p.s
	case "R":
		fmt.Printf("R")
		return 0
	default:
		return evaluatePiece(p, dst)
	}
}
