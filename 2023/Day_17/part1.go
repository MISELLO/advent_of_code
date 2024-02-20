package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type tPosition struct {
	x, y int
}

type tStatus struct {
	pos   tPosition
	dir   tPosition
	steps int
}

type tNode struct {
	status tStatus
	costF  int
	costH  int
}

// HEAP type
type tNodeHeap []tNode

// HEAP interface functions
func (h tNodeHeap) Len() int           { return len(h) }
func (h tNodeHeap) Less(i, j int) bool { return h[i].costF+h[i].costH < h[j].costF+h[i].costH }
func (h tNodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *tNodeHeap) Push(x any)        { *h = append(*h, x.(tNode)) }
func (h *tNodeHeap) Pop() any          { x := (*h)[len(*h)-1]; *h = (*h)[0 : len(*h)-1]; return x }

// consecutiveSteps is the maximum number of steps we can take without turning.
const consecutiveSteps int = 3

// heatLossMap is the input heat loss map
var heatLossMap [][]int

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
	fmt.Println("Loading ...")
	for scn.Scan() {
		l := scn.Text()
		m := strings.Split(l, "")
		var numLine []int
		for i := 0; i < len(m); i++ {
			n, _ := strconv.Atoi(m[i])
			numLine = append(numLine, n)
		}
		heatLossMap = append(heatLossMap, numLine)
	}
	fmt.Println("Grid loaded.")
	fmt.Println("Applying an A* algorithm ...")

	res := aStar()

	// Result
	fmt.Printf("The least head loss a route can incur is \033[1m%d\033[0m\n", res)
}

// aStar or A* is a pathfinding algorithm that works with an Heuristic (h).
func aStar() int {

	// available are the available nodes we can take in a heap.
	available := &tNodeHeap{}
	visited := make(map[tStatus]bool)
	pending := make(map[tStatus]int)

	startRight := tNode{tStatus{tPosition{1, 0}, tPosition{1, 0}, 1}, 0, 0}
	startRight.costF = heatLossMap[startRight.status.pos.y][startRight.status.pos.x]
	startRight.costH = h(startRight.status.pos)
	startDown := tNode{tStatus{tPosition{0, 1}, tPosition{0, 1}, 1}, 0, 0}
	startDown.costF = heatLossMap[startDown.status.pos.y][startDown.status.pos.x]
	startDown.costH = h(startDown.status.pos)

	heap.Init(available)
	heap.Push(available, startRight)
	pending[startRight.status] = startRight.costF
	heap.Push(available, startDown)
	pending[startDown.status] = startDown.costF

	for available.Len() > 0 {
		//fmt.Println("available nodes:", available.Len(), "visited nodes:", len(visited))

		// We get the lower node
		n := heap.Pop(available).(tNode)

		// Check if it is the last point
		if n.status.pos.x == len(heatLossMap[0])-1 && n.status.pos.y == len(heatLossMap)-1 {
			return n.costF
		}

		// Neighbours
		nl := turnLeft(n)
		if valid(nl) && !visited[nl.status] {
			x, e := pending[nl.status]
			if (e && x > nl.costF) || !e {
				heap.Push(available, nl)
				pending[nl.status] = nl.costF
			}
		}
		nr := turnRight(n)
		if valid(nr) && !visited[nr.status] {
			x, e := pending[nr.status]
			if (e && x > nr.costF) || !e {
				heap.Push(available, nr)
				pending[nr.status] = nr.costF
			}
		}
		ns := goStraight(n)
		if valid(ns) && !visited[ns.status] {
			x, e := pending[ns.status]
			if (e && x > ns.costF) || !e {
				heap.Push(available, ns)
				pending[ns.status] = ns.costF
			}
		}

		// Current node is now visited
		visited[n.status] = true
	}

	// We should never reach this point (no more available nodes)
	return -1
}

// h is an heuristic that returns an approximate cost to reach the goal
// more specificly it calculates two possible routes:
//
//	One that goes first East and then South
//	Anotherone that goes first South and then East
//
// The result is the lowest of this two routes
func h(s tPosition) int {
	var eastSouth, southEast int
	var x, y int

	// Route East South
	x, y = s.x, s.y
	for x < len(heatLossMap[0])-1 {
		x++
		eastSouth += heatLossMap[y][x]
	}
	for y < len(heatLossMap)-1 {
		y++
		eastSouth += heatLossMap[y][x]
	}

	// Route South East
	x, y = s.x, s.y
	for y < len(heatLossMap)-1 {
		y++
		southEast += heatLossMap[y][x]
	}
	for x < len(heatLossMap[0])-1 {
		x++
		southEast += heatLossMap[y][x]
	}

	if eastSouth < southEast {
		return eastSouth
	}
	return southEast
}

// turnLeft returns the node resulting of turning left
func turnLeft(n tNode) tNode {
	var dir tPosition = tPosition{n.status.dir.y, -n.status.dir.x}
	var pos tPosition = tPosition{n.status.pos.x + dir.x, n.status.pos.y + dir.y}
	var status tStatus = tStatus{pos, dir, 1}
	var node tNode = tNode{status, 0, 0}
	if valid(node) {
		node.costF = n.costF + heatLossMap[pos.y][pos.x]
		node.costH = h(pos)
	}
	return node
}

// turnRight returns the node resulting of turning right
func turnRight(n tNode) tNode {
	var dir tPosition = tPosition{-n.status.dir.y, n.status.dir.x}
	var pos tPosition = tPosition{n.status.pos.x + dir.x, n.status.pos.y + dir.y}
	var status tStatus = tStatus{pos, dir, 1}
	var node tNode = tNode{status, 0, 0}
	if valid(node) {
		node.costF = n.costF + heatLossMap[pos.y][pos.x]
		node.costH = h(pos)
	}
	return node
}

// goStraight returns the node resulting of going straight
func goStraight(n tNode) tNode {
	var pos tPosition = tPosition{n.status.pos.x + n.status.dir.x, n.status.pos.y + n.status.dir.y}
	var status tStatus = tStatus{pos, n.status.dir, n.status.steps + 1}
	var node tNode = tNode{status, 0, 0}
	if valid(node) {
		node.costF = n.costF + heatLossMap[pos.y][pos.x]
		node.costH = h(pos)
	}
	return node
}

// valid returns true if the node (n) has done no more than consecutiveSteps (3)
// and is inside the grid
func valid(n tNode) bool {

	if n.status.steps > consecutiveSteps {
		return false
	}

	var p tPosition = n.status.pos

	// Inbounds
	return p.x >= 0 && p.y >= 0 && p.x < len(heatLossMap[0]) && p.y < len(heatLossMap)
}
