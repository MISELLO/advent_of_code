package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type opt struct {
	L string
	R string
}

type nodeSteps struct {
	node  string
	steps int
	index int
}

// Instr represent the instructions Left (L) and Right (R) that must be followed
var Instr string

// Maps represents the documents labeled "maps" where the key is node where you are
// and the value are the two options you can choose (L and R).
var Maps map[string]opt

// Cicles tells us the next node and the steps it takes to reach it given a state
// and the index of the instructions.
var Cicles map[nodeSteps]nodeSteps

// Nodes are the different starting nodes and their corresponding steps
var Nodes []nodeSteps

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

	// First line is for the instructions
	fmt.Println("Loading instructions")
	scn.Scan()
	Instr = scn.Text()
	//fmt.Println("Instructions are:", Instr)

	// The rest of the lines are for the maps
	fmt.Println("Loading map nodes")
	Maps = make(map[string]opt)
	Cicles = make(map[nodeSteps]nodeSteps)
	for scn.Scan() {
		l := scn.Text()
		l = strings.ReplaceAll(l, "=", "")
		l = strings.ReplaceAll(l, "(", "")
		l = strings.ReplaceAll(l, ",", "")
		l = strings.ReplaceAll(l, ")", "")
		f := strings.Fields(l)
		if len(f) == 3 {
			Maps[f[0]] = opt{L: f[1], R: f[2]}
			if endA(f[0]) {
				Nodes = append(Nodes, nodeSteps{node: f[0], steps: 0, index: 0})
			}
		}
	}
	fmt.Println("Map nodes loaded")
	//fmt.Println("We have", len(Nodes), "starting points.")
	//fmt.Println(Nodes)
	// Note: We have 6 starting points on the input.

	// Populate the Cicles map
	fmt.Println("Populating the cicles map")

	for i := 0; i < len(Nodes); i++ {
		n := Nodes[i]
		old := n
		for !endZ(n.node) {
			n.Next()
		}

		n.steps = n.steps - old.steps // We need the diference
		old.steps = 0                 // key should have 0 steps for proper indexing
		//fmt.Printf("Cicles[%s-%d-%d] = %s-%d-%d\n", old.node, old.steps, old.index, n.node, n.steps, n.index)
		Cicles[old] = n

		old = n
		_, e := Cicles[old]
		n.Next()
		for !e {
			for !endZ(n.node) || n.steps == 0 {
				n.Next()
			}

			n.steps = n.steps - old.steps
			if n.steps == 0 {
				n.steps = old.steps
			}
			old.steps = 0
			//fmt.Printf("Cicles[%s-%d-%d] = %s-%d-%d\n", old.node, old.steps, old.index, n.node, n.steps, n.index)
			Cicles[old] = n
			old = n
			_, e = Cicles[nodeSteps{node: old.node, steps: 0, index: old.index}]
		}
	}

	//fmt.Println(Cicles)

	// NOTE: Seems that each initial node goes to an end node ending
	// with the instruction index at 0. Also, each of this ending nodes
	// complete a cicle as long as the length of the initial node to
	// reach that same end node.
	// This could be a coincidence or made in purpose, but if this was
	// known from the start this problem could have ben solved just
	// applying the LCM (Lesser Common Multiplier) algorithm.
	// As applying LCM would make it much clear, faster and easyer I
	// will do it both ways, just in case the LCM does not work properly.

	// Calculating the steps to get all paths from ??A to ??Z
	fmt.Println("Calculating steps using LCM")
	result := calculateWithLCM()
	fmt.Println("Our first approach tells us the number of steps are", result)

	fmt.Println("Calculating steps doing displacements, this might take a while... (CTRL + C to cancel)")
	init := Nodes[0]
	Nodes[0] = Cicles[init]
	var steps int
	for !allStepsSame(Nodes) {
		steps = highestSteps(Nodes)
		for i := 0; i < len(Nodes); i++ {
			if Nodes[i].steps < steps {
				k := nodeSteps{node: Nodes[i].node, steps: 0, index: Nodes[i].index}
				Nodes[i].node = Cicles[k].node
				Nodes[i].steps = Nodes[i].steps + Cicles[k].steps
				Nodes[i].index = Cicles[k].index
			}
		}
		fmt.Printf("\b\b\b\b\b\b\b%05.2f%%", (100.0 * float64(steps) / float64(result)))
	}
	fmt.Println("")

	// Result
	fmt.Println("We went from all nodes ending with A to all nodes ending with Z in", steps, "steps.")
}

// endA returns true if the give string ends with the character 'A'
func endA(s string) bool {
	l := len(s)
	if l >= 1 {
		return s[l-1:l] == "A"
	}
	return false
}

// endZ returns true if the give string ends with the character 'Z'
func endZ(s string) bool {
	l := len(s)
	if l >= 1 {
		return s[l-1:l] == "Z"
	}
	return false
}

// allStepsSame returns true if all the steps of the list are the same and > 0
func allStepsSame(sn []nodeSteps) bool {
	if len(sn) == 0 {
		return false
	}
	n := sn[0].steps
	if n == 0 {
		return false
	}
	for i := 0; i < len(sn); i++ {
		if n != sn[i].steps {
			return false
		}
	}
	return true
}

// highestSteps checks a list of nodeSteps and returns the highest steps
// included in it.
// This is used to move the other elements to that number of steps.
func highestSteps(sn []nodeSteps) int {
	max := 0
	for i := 0; i < len(sn); i++ {
		if sn[i].steps > max {
			max = sn[i].steps
		}
	}
	return max
}

// Next moves the node one step further
func (n *nodeSteps) Next() {
	n.steps++
	if Instr[n.index:n.index+1] == "L" {
		n.node = Maps[n.node].L
	} else { // Instr[i:i+1] == "R"
		n.node = Maps[n.node].R
	}
	n.index = (n.index + 1) % len(Instr)
}

// EuclidesGCD is the Euclides algorithm for the Greater Common Divisor
func EuclidesGCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// LCM is the Lesser Common Divisor algorithm following the formula
// LCM(a,b) = a * b / GCD(a,b)
func LCM(a, b int) int {
	return a * b / EuclidesGCD(a, b)
}

// calculateWithLCM
func calculateWithLCM() int {
	var result int
	var first bool = true
	for _, ns := range Cicles {
		if first {
			result = ns.steps
			first = false
		} else {
			result = LCM(result, ns.steps)
		}
	}
	return result
}
