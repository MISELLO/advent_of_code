package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type tInterval struct {
	xMin, xMax, mMin, mMax, aMin, aMax, sMin, sMax int
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
	var count int

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
			// No need to process pieces
		}
	}

	initInterval := tInterval{
		xMin: 1, xMax: 4000,
		mMin: 1, mMax: 4000,
		aMin: 1, aMax: 4000,
		sMin: 1, sMax: 4000,
	}
	count = processLabel(initInterval, "in")

	// Result
	fmt.Printf("\nIf each property can go from 1 to 4000 we will have \033[1m%d\033[0m accepted combinations\n", count)
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

// processLabel interprets the label (l) and returns the combinations
// that satisfy the conditions brought by that label
func processLabel(itv tInterval, l string) int {

	if l == "B" {
		return 0
	}

	if l == "A" {
		intX := itv.xMax - itv.xMin + 1
		intM := itv.mMax - itv.mMin + 1
		intA := itv.aMax - itv.aMin + 1
		intS := itv.sMax - itv.sMin + 1
		return intX * intM * intA * intS
	}

	// We have to go deeper with that label
	w := workflows[l]
	var sum int
	sat, noSat, remain := itv, itv, itv
	for i := 0; i < len(w); i++ {
		if hasOp(w[i]) {
			sat, noSat = ifRule(remain, w[i])
			sum += processLabel(sat, w[i].dst)
			remain = noSat
		} else { // has no operation (last rule of the workflow)
			sum += processLabel(remain, w[i].dst)
		}
	}
	return sum
}

// hasOp has Operand returns true if the rule (h) has an operand ("<" or ">")
func hasOp(r tRule) bool {
	return r.op == "<" || r.op == ">"
}

// ifRule takes the rule (r) and it applies its condition to the interval (itv).
// The result is 2 intervals, the one that satisfies the condition and
// the one that doesn't (sat & noSat).
func ifRule(itv tInterval, r tRule) (sat, noSat tInterval) {
	sat, noSat = itv, itv
	switch r.prop {
	case "x":
		if r.op == "<" {
			sat.xMax = r.val - 1
			noSat.xMin = r.val
		} else { // r.op == ">"
			sat.xMin = r.val + 1
			noSat.xMax = r.val
		}
	case "m":
		if r.op == "<" {
			sat.mMax = r.val - 1
			noSat.mMin = r.val
		} else { // r.op == ">"
			sat.mMin = r.val + 1
			noSat.mMax = r.val
		}
	case "a":
		if r.op == "<" {
			sat.aMax = r.val - 1
			noSat.aMin = r.val
		} else { // r.op == ">"
			sat.aMin = r.val + 1
			noSat.aMax = r.val
		}
	case "s":
		if r.op == "<" {
			sat.sMax = r.val - 1
			noSat.sMin = r.val
		} else { // r.op == ">"
			sat.sMin = r.val + 1
			noSat.sMax = r.val
		}
	}
	return sat, noSat
}
