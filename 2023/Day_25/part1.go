package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type tNodeInfo struct {
	group bool
	conn  []string
}

var nodes map[string]tNodeInfo

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
	fmt.Println("Loading components ...")
	nodes = make(map[string]tNodeInfo)
	for scn.Scan() {
		l := scn.Text()
		parts := strings.Split(l, ": ")
		a := parts[0]
		links := strings.Split(parts[1], " ")
		for _, b := range links {
			addConnection(a, b)
		}
	}

	// Calculate
	fmt.Println("Counting connections and changing groups ...")
	fmt.Println("We have", len(nodes), "nodes.")
	var x int
	for n, _ := range nodes {
		// We will assign each node to a group with no criteria
		tmp := nodes[n]
		tmp.group = x%2 == 0
		nodes[n] = tmp
		x++
	}
	fmt.Println("Applying the algorithm of Kernighan-Lin.")
	var separationLinks, oldSL, iteration int
	forceOneChange := false
	separationLinks = sepCon()
	for separationLinks > 3 {
		fmt.Println("Iteration", iteration, "(we have ", separationLinks, "separation links)")
		for n, _ := range nodes {
			if forceOneChange && noDifference(n) {
				forceOneChange = false
				tmp := nodes[n]
				tmp.group = !tmp.group
				nodes[n] = tmp
				separationLinks = sepCon()
			} else if worthChangeGroup(n) {
				tmp := nodes[n]
				tmp.group = !tmp.group
				nodes[n] = tmp
				separationLinks = sepCon()
			}
		}
		iteration++
		if oldSL == separationLinks {
			forceOneChange = true
		} else {
			oldSL = separationLinks
		}
	}
	fmt.Println("End of iterations, we now have", separationLinks, "separation links.")

	g1Count, g2Count := groupCount()
	if g1Count == 0 || g2Count == 0 {
		fmt.Println("Seems all nodes have ended in the same group. Please, \033[1mtry again\033[0m.")
	}

	// Result
	fmt.Printf("One group has %d components, and the other group has %d components.\n", g1Count, g2Count)
	fmt.Printf("The product of this two groups is \033[1m%d\033[0m.\n", g1Count*g2Count)
}

// addConnection updates the proper data structures and adds a new connection between two nodes
func addConnection(a, b string) {
	var tmp tNodeInfo
	tmp = nodes[a]
	tmp.conn = append(nodes[a].conn, b)
	nodes[a] = tmp
	tmp = nodes[b]
	tmp.conn = append(nodes[b].conn, a)
	nodes[b] = tmp
}

// groupCount returns the number of elements of both groups
func groupCount() (int, int) {
	x := 0
	for _, n := range nodes {
		if n.group {
			x++
		}
	}
	return x, len(nodes) - x
}

// sepCon returns the number of connections that separate the 2 groups
func sepCon() int {
	x := 0
	for _, n := range nodes {
		if n.group {
			continue
		}
		// For sure n.group == false
		for _, c := range n.conn {
			if nodes[c].group {
				x++
			}
		}
	}
	return x
}

// worthChangeGroup tells if it is worth to change the node (n) to the
// other group, it tells if there will be less connections to the
// other group if the node (n) changes groups.
func worthChangeGroup(n string) bool {
	x := 0
	for _, c := range nodes[n].conn {
		if nodes[c].group == nodes[n].group {
			x++
		}
	}
	return len(nodes[n].conn)-x > x
}

// noDifference is similar to worthChangeGroup, but it returns true when
// there will be no difference in the number of connections to the
// other group. It is useful when in one iteration the number of
// connections has not changed but we still have too many.
func noDifference(n string) bool {
	x := 0
	for _, c := range nodes[n].conn {
		if nodes[c].group == nodes[n].group {
			x++
		}
	}
	return len(nodes[n].conn)-x == x
}
