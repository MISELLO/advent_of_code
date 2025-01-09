package main

import (
	"bufio"
	"fmt"
	"os"
)

// tPos represents a position or a direction
// on a 2D board
type tPos struct {
	x, y int
}

// tLab represents a lab.
// A lap contains a map and a guard
// The map is a 2D slice being '#' obstacles.
// The guard has a position and a direction.
type tLab struct {
	board    [][]byte
	guardPos tPos
	guardDir tPos
}

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

	// Our lab (with it's map (board) and it's guard)
	var lab tLab

	// Positions visited
	visited := make(map[tPos]struct{})

	// Load map (board)
	fmt.Println("Loading map ...")
	for scn.Scan() {
		l := scn.Text()
		lab.board = append(lab.board, []byte(l))
	}

	// Initialize guard
	fmt.Println("Positioning guard ...")
	lab.InitGuard()

	// Processing
	fmt.Println("Guard patroling ...")
	for lab.GuardInArea() {
		visited[lab.guardPos] = struct{}{}
		if lab.GuardFacingObstacle() {
			lab.GuardTurnRight()
		} else {
			lab.GuardAdvance()
		}
	}

	// fmt.Printf(" Pos: %+v\n Dir: %+v\n", lab.guardPos, lab.guardDir)

	// Result
	fmt.Printf("The guard will visit \033[1m%d\033[0m before leaving the mapped area.\n", len(visited))
}

// InitGuard initializes the position and direction of the guard
func (l *tLab) InitGuard() {
	for y := 0; y < len(l.board); y++ {
		for x := 0; x < len(l.board[y]); x++ {
			if l.board[y][x] != '.' && l.board[y][x] != '#' {
				// Can only be '^', facing North (would be more complex if it could start facing East, South or West)
				l.guardPos = tPos{x: x, y: y}
				l.guardDir = tPos{x: 0, y: -1}
				return // There is only one guard
			}
		}
	}
}

// Returns true if the guard is still in the mapped area
func (l *tLab) GuardInArea() bool {
	return l.guardPos.x >= 0 && l.guardPos.y >= 0 && l.guardPos.x < len(l.board[0]) && l.guardPos.y < len(l.board)
}

// guardAdvance makes the guard advance in the direction he is facing
// Does not consider the obstacles or map limits
func (l *tLab) GuardAdvance() {
	l.guardPos = tPos{x: l.guardPos.x + l.guardDir.x, y: l.guardPos.y + l.guardDir.y}
}

// guardFacingObstacle returns true if there will be an obstacle in the next step of the guard
// returns false if no obstacles or outside mapped area
func (l *tLab) GuardFacingObstacle() bool {
	// Next position
	p := l.guardPos
	p.x += l.guardDir.x
	p.y += l.guardDir.y

	// Outside mapped area
	if p.x < 0 || p.y < 0 || p.x >= len(l.board[0]) || p.y >= len(l.board) {
		return false
	}

	// Obstacle
	return l.board[p.y][p.x] == '#'
}

// guardTurnRight makes the guard to change the direction to his right
func (l *tLab) GuardTurnRight() {
	switch l.guardDir {
	case tPos{0, -1}:
		l.guardDir = tPos{x: 1, y: 0}
	case tPos{1, 0}:
		l.guardDir = tPos{x: 0, y: 1}
	case tPos{0, 1}:
		l.guardDir = tPos{x: -1, y: 0}
	case tPos{-1, 0}:
		l.guardDir = tPos{x: 0, y: -1}
	}
}
