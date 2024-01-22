package main

import (
	"bufio"
	"fmt"
	"hash/crc32"
	"os"
	"strconv"
)

const cicles int = 1000000000

var platform [][]byte

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
		platform = append(platform, []byte(l))
	}
	fmt.Println("Platform loaded.")
	//printPlatform()

	// Tilding
	var e bool
	var t int
	var exitIndex int
	mapIndex := make(map[string]int) // Hash --> Index
	mapLoad := make(map[int]int)     // Index --> Load
	l := load()
	h := hash()
	fmt.Printf("Cicle: %3d, load: %6d, hash: %s\n", 0, l, h)
	mapIndex[h] = 0
	mapLoad[0] = l
	for i := 1; i <= cicles; i++ {
		moveRocksNorth()
		moveRocksWest()
		moveRocksSouth()
		moveRocksEast()
		l = load()
		h = hash()
		fmt.Printf("Cicle: %3d, load: %6d, hash: %s\n", i, l, h)
		t, e = mapIndex[h]
		if e {
			fmt.Printf("\n Hash \"%s\" was already found on cicle %d.\n This means we are now into a loop.\n", h, t)
			exitIndex = i
			break
		} else {
			mapIndex[h] = i
			mapLoad[i] = l
		}
	}

	if e {
		fmt.Printf(" We ended at %d because it goes back to %d.\n", exitIndex, t)
		fmt.Printf(" The loop is %d - %d = %d units length.\n", exitIndex, t, exitIndex-t)
		fmt.Printf(" %d total cicles - %d (starting loop cicle) is %d pending iterations from the start of the loop.\n", cicles, t, cicles-t)
		resultIndex := (cicles-t)%(exitIndex-t) + t
		fmt.Printf(" %d modus %d is %d + the starting loop cicle (%d) is %d.\n", cicles-t, exitIndex-t, (cicles-t)%(exitIndex-t), t, resultIndex)
		fmt.Printf(" The load on the cicle %d will be the same to the cicle %d.\n", cicles, resultIndex)
		fmt.Printf(" The load on cicle %d is \033[1m%d\033[0m\n", resultIndex, mapLoad[resultIndex])
	} else {
		fmt.Println("I can't believe we did", cicles, "iterations.")
	}
}

// printPlatform is a debug function that prints on the screen the current state of the platform
func printPlatform() {
	for i := 0; i < len(platform); i++ {
		fmt.Println(string(platform[i]))
	}
}

// moveRocksNorth moves all the rocks to the North
func moveRocksNorth() {
	for y := 1; y < len(platform); y++ {
		for x := 0; x < len(platform[y]); x++ {
			if platform[y][x] == 'O' {
				displaceRockNorth(x, y)
			}
		}
	}
}

// moveRocksWest moves all the rocks to the West
func moveRocksWest() {
	for x := 1; x < len(platform[0]); x++ {
		for y := 0; y < len(platform); y++ {
			if platform[y][x] == 'O' {
				displaceRockWest(x, y)
			}
		}
	}
}

// moveRocksSouth moves all the rocks to the South
func moveRocksSouth() {
	for y := len(platform) - 2; y >= 0; y-- {
		for x := 0; x < len(platform[y]); x++ {
			if platform[y][x] == 'O' {
				displaceRockSouth(x, y)
			}
		}
	}
}

// moveRocksEast moves all the rocks to the East
func moveRocksEast() {
	for x := len(platform[0]) - 2; x >= 0; x-- {
		for y := 0; y < len(platform); y++ {
			if platform[y][x] == 'O' {
				displaceRockEast(x, y)
			}
		}
	}
}

// displaceRockNorth moves a the rock at (x, y) to the North until it hits an obstacle
func displaceRockNorth(x, y int) {
	for y > 0 && platform[y-1][x] == '.' {
		platform[y][x], platform[y-1][x] = platform[y-1][x], platform[y][x]
		y--
	}
}

// displaceRockWest moves a the rock at (x, y) to the West until it hits an obstacle
func displaceRockWest(x, y int) {
	for x > 0 && platform[y][x-1] == '.' {
		platform[y][x], platform[y][x-1] = platform[y][x-1], platform[y][x]
		x--
	}
}

// displaceRockSouth moves a the rock at (x, y) to the South until it hits an obstacle
func displaceRockSouth(x, y int) {
	for y < len(platform)-1 && platform[y+1][x] == '.' {
		platform[y][x], platform[y+1][x] = platform[y+1][x], platform[y][x]
		y++
	}
}

// displaceRockEast moves a the rock at (x, y) to the East until it hits an obstacle
func displaceRockEast(x, y int) {
	for x < len(platform[y])-1 && platform[y][x+1] == '.' {
		platform[y][x], platform[y][x+1] = platform[y][x+1], platform[y][x]
		x++
	}
}

// load returns the sum of each rock times it's row.
func load() int {
	var sum int
	multiplier := len(platform)
	for y := 0; y < len(platform); y++ {
		for x := 0; x < len(platform[y]); x++ {
			if platform[y][x] == 'O' {
				sum += multiplier
			}
		}
		multiplier--
	}
	return sum
}

// hash returns a string that identifies a unique platform
func hash() string {
	h := crc32.NewIEEE()
	var p []byte
	for i := 0; i < len(platform); i++ {
		p = append(p, platform[i]...)
	}
	h.Write(p)
	v := h.Sum32()
	s := strconv.FormatInt(int64(v), 16)
	for len(s) < 8 {
		s = "0" + s
	}
	return s
}
