package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Number of times we will try the first approach before trying the second one.
const TryTimes int = 10

// if we can't find the exact position and speed we will vary errorUnits the result we get
const errorUnits int = 4

type tHailstone struct {
	px, py, pz, vx, vy, vz int
}

type tPos struct {
	x, y, z int
}

var hailList []tHailstone

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
	fmt.Println("Loading hailstones ...")
	for scn.Scan() {
		var h tHailstone
		l := scn.Text()
		l = strings.ReplaceAll(l, "  ", " ") // Example had a double space.
		// This turns every double space into a single space.
		part := strings.Split(l, " @ ")
		position := strings.Split(part[0], ", ")
		speed := strings.Split(part[1], ", ")
		h.px, _ = strconv.Atoi(position[0])
		h.py, _ = strconv.Atoi(position[1])
		h.pz, _ = strconv.Atoi(position[2])
		h.vx, _ = strconv.Atoi(speed[0])
		h.vy, _ = strconv.Atoi(speed[1])
		h.vz, _ = strconv.Atoi(speed[2])
		hailList = append(hailList, h)
	}

	// Calculate
	fmt.Println("There are two approaches for this and both need 3 hailstones.")
	fmt.Println(" - A simple approach that need ceratin assumptions and bruteforce.")
	fmt.Println(" - Another approach that needs less assumptions and no bruteforce, but with more formulas and calculations.")
	fmt.Println()
	fmt.Println("We will try the first approach", TryTimes, "times before trying the second one.")

	var res, i int
	for n := 0; res == 0 && n < TryTimes; n++ {
		fmt.Println(" First approach attempt", n)
		firstApproach(&res, &i)
	}

	if res == 0 {
		fmt.Println()
		fmt.Println(" Calculating the second approach...")
		secondApproach(&res)
	}

	// Result
	fmt.Println()
	fmt.Printf("If we add the starting coordinates of our rock we will get \033[1m%d\033[0m.\n", res)
}

// firstApproach does a first attempt to resolve the problem
// r is a pointer to the result
// i is the index of the hailstones we will use
func firstApproach(r, i *int) {
	// We will work on times. From 0 to #hailstones
	// We assume that, in this interval of time, our trajectory will hit 2 hailstones.
	// We will take the first hailstone at ta = 0 and create a trajectory to the second hailstone at tb = 1.
	// If this trajectory it does not hit a third hailstone, then we increase the tb.
	// Once checked all possible tb we will increase ta and repeat the process.
	// If all combinations have been done, then we have to change strategies.

	ha := hailList[*i]
	hb := hailList[*i+1]
	hc := hailList[*i+2]
	*i += 3
	found := false
	for ta := 0; !found && ta <= len(hailList); ta++ {
		for tb := 0; !found && tb <= len(hailList); tb++ {
			if ta == tb {
				// The rock can't hit 2 hailstones at the same time
				continue
			}
			pa := timePos(ha, ta)
			pb := timePos(hb, tb)

			rock := computeTrajectory(pa, pb, ta, tb)
			if willCollide(rock, hc) {
				*r = rock.px + rock.py + rock.pz
				found = true
			}
		}
	}
}

// secondApproach does not use bruteforce, but needs 3 hailstones and more equations to get the answer.
// r is a pointer to the result
func secondApproach(r *int) {
	// The position of each coordinade can be expressed as:
	//	x = p + t*s
	//	Where x is the final position, p the initial position, t the time and s the speed.
	//	We don't know the position or speed of the rock, but we know the position and speed of every hailstone and
	//	we know that in any time (t) the position of the rock will be the same as any hailstone.
	// 		Rp + (t * Rs) = Hp + (t * Hs)
	//		(t * Rs) - (t * Hs) = Hp - Rp
	//		t * (Rs - Hs) = Hp - Rp
	//		t = (Hp - Rp) / (Rs - Hs)

	// We can generate 3 equations for each hailstone. And each hailstone introduces a new unknown variable (the interception time)
	// Step 0.- We have 6 unknown variables (R.px, R.py, R.pz, R.vx, R.vy, R.vz)
	// 	Note: "R" stands for Rock, "p" stands for position and "v" for velocity (then we have the xyz coordinades).
	// Step 1.- We add 1 Hailstone (a).
	//	We have 1 more unknown valirable (ta).
	//	We have 3 equations, one for each coordinade.
	//		ta = (a.px - R.px) / (R.vx - a.vx)
	//		ta = (a.py - R.py) / (R.vy - a.vy)
	//		ta = (a.pz - R.pz) / (R.vz - a.vz)
	// Step 2.- We add one more hailstone (b).
	//	We now have 1 more unknown variable (tb). Addint a total of 8.
	//	But we now have 6 equatins.
	//		tb = (b.px - R.px) / (R.vx - b.vx)
	//		...
	// Step 3.- Adding one more hailstone (c).
	//	We now have a total of 9 unknown variables and 9 equations. The problem can be solved with lineal algebra.

	// We take our 3 hailstones
	a := hailList[0]
	b := hailList[1]
	c := hailList[2]

	// Step 4.- We have to play and combine the equations in order to solve it.
	//	I will first work only with coordinades x and y of hailstone (a).
	//		 - As time is the same we can equal both equations (this way we get rid of times)
	//		ta = (a.px - R.px) / (R.vx - a.vx) = (a.py - R.py) / (R.vy - a.vy)
	//		(a.px - R.px) / (R.vx - a.vx) = (a.py - R.py) / (R.vy - a.vy)
	//		 - We transform the denominators to products on the other side
	//		(a.px - R.px) * (R.vy - a.vy) = (a.py - R.py) * (R.vx - a.vx)
	//		 - We apply the distributive property for products of sums.
	//		a.px * R.vy - a.px * a.vy - R.px * R.vy + R.px * a.vy = a.py * R.vx - a.py * a.vx - R.py * R.vx + R.py * a.vx
	//		 - We move to the LHS the purely unknown variables. Note that this LHS does not depend on the hailstone we use.
	//		R.py * R.vx - R.px * R.vy = a.py * R.vx - a.py * a.vx + R.py * a.vx - a.px * R.vy + a.px * a.vy - R.px * a.vy
	// Step 5.- We can use the RHS for hailstone (a) from above and equal it to it's homologue for hailstone (b)
	//		  a.py * R.vx - a.py * a.vx + R.py * a.vx - a.px * R.vy + a.px * a.vy - R.px * a.vy =
	//		= b.py * R.vx - b.py * b.vx + R.py * b.vx - b.px * R.vy + b.px * b.vy - R.px * b.vy
	//		 - And group it
	//		 R.px * b.vy - R.px * a.vy + R.py * a.vx - R.py * b.vx + a.py * R.vx - b.py * R.vx + b.px * R.vy - a.px * R.vy =
	//		= a.py * a.vx - a.px * a.vy - b.py * b.vx + b.px * b.vy
	//	->	R.px * (b.vy - a.vy) + R.py * (a.vx - b.vx) + R.vx * (a.py - b.py) + R.vy * (b.px - a.px) = a.py * a.vx - a.px * a.vy - b.py * b.vx + b.px * b.vy
	//		 - This was the first formula for X and Y with the pair of hailstones (a) and (b).
	//		 - Now we will do the same for Y and Z with the same pairs (a) and (b).
	//	->	R.py * (b.vz - a.vz) + R.pz * (a.vy - b.vy) + R.vy * (a.pz - b.pz) + R.vz * (b.py - a.py) = a.pz * a.vy - a.py * a.vz - b.pz * b.vy + b.py * b.vz
	//		 - And now for Z and X with the same pairs (a) and (b).
	//	->	R.pz * (b.vx - a.vx) + R.px * (a.vz - b.vz) + R.vz * (a.px - b.px) + R.vx * (b.pz - a.pz) = a.px * a.vz - a.pz * a.vx - b.px * b.vz + b.pz * b.vx
	//		 - We repeat the same for X, Y and Z but for the pairs (b) and (c) instead.
	//	->	R.py * (c.vz - b.vz) + R.pz * (b.vy - c.vy) + R.vy * (b.pz - c.pz) + R.vz * (c.py - b.py) = b.pz * b.vy - b.py * b.vz - c.pz * c.vy + c.py * c.vz
	//	->	R.pz * (c.vx - b.vx) + R.px * (b.vz - c.vz) + R.vz * (b.px - c.px) + R.vx * (c.pz - b.pz) = b.px * b.vz - b.pz * b.vx - c.px * c.vz + c.pz * c.vx
	//	->	R.px * (c.vy - b.vy) + R.py * (b.vx - c.vx) + R.vx * (b.py - c.py) + R.vy * (c.px - b.px) = b.py * b.vx - b.px * b.vy - c.py * c.vx + c.px * c.vy
	// With this 6 equations and 6 unknown variables we can create the expanded matrix
	// But before that we will order the LHS a bit, so it will be easier to extract the coeficients.
	//	>>	R.px * (b.vy - a.vy) + R.py * (a.vx - b.vx) + R.vx * (a.py - b.py) + R.vy * (b.px - a.px) = a.py * a.vx - a.px * a.vy - b.py * b.vx + b.px * b.vy
	//	>>	R.py * (b.vz - a.vz) + R.pz * (a.vy - b.vy) + R.vy * (a.pz - b.pz) + R.vz * (b.py - a.py) = a.pz * a.vy - a.py * a.vz - b.pz * b.vy + b.py * b.vz
	//	>>	R.px * (a.vz - b.vz) + R.pz * (b.vx - a.vx) + R.vx * (b.pz - a.pz) + R.vz * (a.px - b.px) = a.px * a.vz - a.pz * a.vx - b.px * b.vz + b.pz * b.vx
	//	>>	R.py * (c.vz - b.vz) + R.pz * (b.vy - c.vy) + R.vy * (b.pz - c.pz) + R.vz * (c.py - b.py) = b.pz * b.vy - b.py * b.vz - c.pz * c.vy + c.py * c.vz
	//	>>	R.px * (b.vz - c.vz) + R.pz * (c.vx - b.vx) + R.vx * (c.pz - b.pz) + R.vz * (b.px - c.px) = b.px * b.vz - b.pz * b.vx - c.px * c.vz + c.pz * c.vx
	//	>	R.px * (c.vy - b.vy) + R.py * (b.vx - c.vx) + R.vx * (b.py - c.py) + R.vy * (c.px - b.px) = b.py * b.vx - b.px * b.vy - c.py * c.vx + c.px * c.vy

	// The expanded matrix that represent the 6 equations and 6 unknown variables (turned into 64 bit float)
	var expMatrix = [][]float64{
		{float64(b.vy - a.vy), float64(a.vx - b.vx), 0.0, float64(a.py - b.py), float64(b.px - a.px), 0.0, float64(a.py*a.vx - a.px*a.vy - b.py*b.vx + b.px*b.vy)},
		{0.0, float64(b.vz - a.vz), float64(a.vy - b.vy), 0.0, float64(a.pz - b.pz), float64(b.py - a.py), float64(a.pz*a.vy - a.py*a.vz - b.pz*b.vy + b.py*b.vz)},
		{float64(a.vz - b.vz), 0.0, float64(b.vx - a.vx), float64(b.pz - a.pz), 0.0, float64(a.px - b.px), float64(a.px*a.vz - a.pz*a.vx - b.px*b.vz + b.pz*b.vx)},
		{0.0, float64(c.vz - b.vz), float64(b.vy - c.vy), 0.0, float64(b.pz - c.pz), float64(c.py - b.py), float64(b.pz*b.vy - b.py*b.vz - c.pz*c.vy + c.py*c.vz)},
		{float64(b.vz - c.vz), 0.0, float64(c.vx - b.vx), float64(c.pz - b.pz), 0.0, float64(b.px - c.px), float64(b.px*b.vz - b.pz*b.vx - c.px*c.vz + c.pz*c.vx)},
		{float64(c.vy - b.vy), float64(b.vx - c.vx), 0.0, float64(b.py - c.py), float64(c.px - b.px), 0.0, float64(b.py*b.vx - b.px*b.vy - c.py*c.vx + c.px*c.vy)},
	}

	rock := solveExpandedMatrixLinealAlgebra(expMatrix)

	// We could have lost some precission.
	// We will check if we collide at least to the first hailstone
	if willCollide(hailList[0], rock) {
		// We found our rock
		*r = rock.px + rock.py + rock.pz
	} else {
		fmt.Println("Seems there are some precission errors and the rock will not hit the hailstones.")
		fmt.Println("Now that we have an approximate solution will try to find by bruteforce the exact position and speed.")
		for vz := rock.vz - errorUnits; vz <= rock.vz+errorUnits; vz++ {
			for vy := rock.vy - errorUnits; vy <= rock.vy+errorUnits; vy++ {
				for vx := rock.vx - errorUnits; vx <= rock.vx+errorUnits; vx++ {
					for pz := rock.pz - errorUnits; pz <= rock.pz+errorUnits; pz++ {
						for py := rock.py - errorUnits; py <= rock.py+errorUnits; py++ {
							for px := rock.px - errorUnits; px <= rock.px+errorUnits; px++ {
								if willCollide(hailList[0], tHailstone{px, py, pz, vx, vy, vz}) {
									fmt.Println("Found it!")
									*r = px + py + pz
									return
								}
							}
						}
					}
				}
			}
		}
		fmt.Println("Unfortunatelly, we could not find the exact position and speed direction of the rock.")
	}
}

// timePos returns the position of a hailstone (h) at a given time (t)
func timePos(h tHailstone, t int) tPos {
	var p tPos
	p.x = h.px + (h.vx * t)
	p.y = h.py + (h.vy * t)
	p.z = h.pz + (h.vz * t)
	return p
}

// computeTrajectory returns a position and speed in the form of a hailstone
// given 2 positions (pa & pb) and 2 times (ta & tb).
func computeTrajectory(pa, pb tPos, ta, tb int) tHailstone {
	var r tHailstone
	r.vx = (pb.x - pa.x) / (tb - ta)
	r.vy = (pb.y - pa.y) / (tb - ta)
	r.vz = (pb.z - pa.z) / (tb - ta)
	r.px = pa.x - (ta * r.vx)
	r.py = pa.y - (ta * r.vy)
	r.pz = pa.z - (ta * r.vz)
	return r
}

// willCollide returns true if the objects (a) and (b) have a collision trajectory.
func willCollide(a, b tHailstone) bool {
	if float64(a.vx)/float64(b.vx) == float64(a.vy)/float64(b.vy) && float64(a.vx)/float64(b.vx) == float64(a.vz)/float64(b.vz) {
		// Lines are parallel
		return false
	}

	dx := a.vx - b.vx
	dy := a.vy - b.vy
	dz := a.vz - b.vz

	if dx == 0 || dy == 0 || dz == 0 {
		// We can't divide by zero
		return false
	}

	tx := float64(b.px-a.px) / float64(dx)
	ty := float64(b.py-a.py) / float64(dy)
	tz := float64(b.pz-a.pz) / float64(dz)

	if tx != ty || tx != tz {
		// Trajectories cross, but not at the same time
		return false
	}

	return true
}

// solveExpandedMatrixLinealAlgebra solves the expanded matrix that represents a system
// of equations that contains N equations and N unknown variables.
// It supposes it exists one solution (can be solved with lineal algebra).
// m is the expanded matrix where the rows are the equations and the columns are
// the values of each unknown variable, being this: px, py, pz, vx, vy, vz, (and the independent value)
// function returns a type Hailstone because it can be used as the rock (it has a position in space and
// a direction in each coordinate)
func solveExpandedMatrixLinealAlgebra(m [][]float64) tHailstone {

	transform := func(a, b []float64, v float64) {
		for i := 0; i < len(a); i++ {
			a[i] = a[i] - (v * b[i])
		}
	}

	//printMatrix(m)

	for y := 0; y < len(m); y++ { // For each line
		// Make sure value m[y][y] == 1
		if m[y][y] == 0.0 {
			// We search for another line and we swap them
			swapped := false
			for i := y + 1; !swapped && i < len(m); i++ {
				if m[i][y] != 0.0 {
					m[y], m[i] = m[i], m[y]
					swapped = true
				}
			}
			if !swapped {
				printMatrix(m)
				fmt.Println("Could not solve the expanded matrix :(")
				return tHailstone{0, 0, 0, 0, 0, 0}
			}
		}
		if m[y][y] != 1 {
			d := m[y][y]
			for i := 0; i < len(m[y]); i++ {
				m[y][i] = m[y][i] / d
			}
		}

		// Make sure the column below [y][y] are all 0's
		for i := y + 1; i < len(m); i++ {
			v := m[i][y]
			if v != 0 {
				transform(m[i], m[y], v)
			}
		}
	}

	// Now the extended matrix should be solved and we should know all unknown variables.

	//printMatrix(m)
	vz := m[5][6]
	vy := m[4][6] - (vz * m[4][5])
	vx := m[3][6] - (vz * m[3][5]) - (vy * m[3][4])
	pz := m[2][6] - (vz * m[2][5]) - (vy * m[2][4]) - (vx * m[2][3])
	py := m[1][6] - (vz * m[1][5]) - (vy * m[1][4]) - (vx * m[1][3]) - (pz * m[1][2])
	px := m[0][6] - (vz * m[0][5]) - (vy * m[0][4]) - (vx * m[0][3]) - (pz * m[0][2]) - (py * m[0][1])

	var r tHailstone
	//fmt.Printf("px=%0.3f, py=%0.3f, pz=%0.3f, vx=%0.3f, vy=%0.3f, vz=%0.3f, TOTAL=%0.5f\n", px, py, pz, vx, vy, vz, px+py+pz)
	r.px = int(px)
	r.py = int(py)
	r.pz = int(pz)
	r.vx = int(vx)
	r.vy = int(vy)
	r.vz = int(vz)

	//fmt.Println(r)

	return r
}

// printMatrix prints the current extended matrix
func printMatrix(m [][]float64) {
	for y := 0; y < len(m); y++ {
		for x := 0; x < len(m[y]); x++ {
			if x == len(m[y])-1 {
				fmt.Printf(" |")
			}
			if x >= 6 {
				fmt.Printf(" %22.2f", m[y][x])
			} else if x >= 5 {
				fmt.Printf(" %19.2f", m[y][x])
			} else if x >= 3 {
				fmt.Printf(" %18.2f", m[y][x])
			} else {
				fmt.Printf(" %7.2f", m[y][x])
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

// collideAll checks if the rock collides all the hailstones
func collideAll(r tHailstone) bool {
	for i, h := range hailList {
		if !willCollide(r, h) {
			fmt.Println(i, "Compared:", r, "and", h)
			return false
		}
	}
	return true
}
