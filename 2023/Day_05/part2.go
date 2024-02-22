package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

type statusT int

const (
	ini statusT = iota
	seeds
	soil
	fertilizer
	water
	light
	temperature
	humidity
	location
	end
)

type mapT struct {
	dst int
	src int
	ran int
}

var toSoil, toFertilizer, toWater, toLight, toTemperature, toHumidity, toLocation []mapT

var seedsList []int

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

	fmt.Println("Loading almanac")
	load(scn)
	fmt.Println("Almanac loaded")

	lowLocList := make([]int, len(seedsList)/2)
	var wg sync.WaitGroup
	for i := 0; i < len(seedsList); i += 2 {
		wg.Add(1)
		go concurrentCalculation(i/2, seedsList[i], seedsList[i]+seedsList[i+1], &lowLocList[i/2], &wg)
	}
	wg.Wait()
	var lowestLocation int
	for i := 0; i < len(lowLocList); i++ {
		if lowestLocation == 0 || lowLocList[i] < lowestLocation {
			lowestLocation = lowLocList[i]
		}
	}
	fmt.Println("")
	fmt.Println("The lowest location number is ", lowestLocation)
}

func convert(x int, mt []mapT) int {
	for i := 0; i < len(mt); i++ {
		if x >= mt[i].src && x < mt[i].src+mt[i].ran {
			return mt[i].dst + (x - mt[i].src)
		}
	}
	return x
}

func concurrentCalculation(id, sStart, sEnd int, w *int, wg *sync.WaitGroup) {
	defer wg.Done()
	var lowLoc int
	for i := sStart; i < sEnd; i++ {
		s := i
		t := convert(s, toSoil)
		u := convert(t, toFertilizer)
		v := convert(u, toWater)
		w := convert(v, toLight)
		x := convert(w, toTemperature)
		y := convert(x, toHumidity)
		z := convert(y, toLocation)
		if lowLoc == 0 || z < lowLoc {
			lowLoc = z
		}
	}
	*w = lowLoc
	fmt.Println("Done", id)
}

// load loads all input data
func load(scn *bufio.Scanner) {
	var s statusT = ini
	for scn.Scan() {
		l := scn.Text()
		if s == ini {
			// Seeds
			fmt.Println(" Loading seeds")
			loadSeeds(l)
			s = soil
		} else if s == soil {
			// Soil
			loadSeedToSoil(scn)
			s = fertilizer
		} else if s == fertilizer {
			// Fertilizer
			loadSoilToFertilizer(scn)
			s = water
		} else if s == water {
			// Water
			loadFertilizertoWater(scn)
			s = light
		} else if s == light {
			// Light
			loadWaterToLight(scn)
			s = temperature
		} else if s == temperature {
			// Temperature
			loadLightToTemperature(scn)
			s = humidity
		} else if s == humidity {
			// Humidity
			loadTemperatureToHumidity(scn)
			s = location
		} else if s == location {
			// Location
			loadHumidityToLocation(scn)
			s = end
		}
	}
}

// loadSeeds loads the seeds into the variable seedsList
func loadSeeds(l string) {
	parts := strings.Split(l, ":")
	seedNums := strings.Fields(parts[1])
	for i := 0; i < len(seedNums); i++ {
		n, _ := strconv.Atoi(seedNums[i])
		seedsList = append(seedsList, n)
	}
}

// loadSeedToSoil
func loadSeedToSoil(scn *bufio.Scanner) {
	scn.Scan()
	scn.Text()
	fmt.Println(" Loading seed-to-soil map")
	scn.Scan()
	l := scn.Text()
	for l != "" {
		nums := strings.Fields(l)
		var m mapT
		m.dst, _ = strconv.Atoi(nums[0])
		m.src, _ = strconv.Atoi(nums[1])
		m.ran, _ = strconv.Atoi(nums[2])
		toSoil = append(toSoil, m)
		scn.Scan()
		l = scn.Text()
	}
}

// loadSoilToFertilizer
func loadSoilToFertilizer(scn *bufio.Scanner) {
	scn.Scan()
	scn.Text()
	fmt.Println(" Loading soil-to-fertilizer map")
	scn.Scan()
	l := scn.Text()
	for l != "" {
		nums := strings.Fields(l)
		var m mapT
		m.dst, _ = strconv.Atoi(nums[0])
		m.src, _ = strconv.Atoi(nums[1])
		m.ran, _ = strconv.Atoi(nums[2])
		toFertilizer = append(toFertilizer, m)
		scn.Scan()
		l = scn.Text()
	}
}

// loadFertilizertoWater
func loadFertilizertoWater(scn *bufio.Scanner) {
	scn.Scan()
	scn.Text()
	fmt.Println(" Loading fertilizer-to-water map")
	scn.Scan()
	l := scn.Text()
	for l != "" {
		nums := strings.Fields(l)
		var m mapT
		m.dst, _ = strconv.Atoi(nums[0])
		m.src, _ = strconv.Atoi(nums[1])
		m.ran, _ = strconv.Atoi(nums[2])
		toWater = append(toWater, m)
		scn.Scan()
		l = scn.Text()
	}
}

// loadWaterToLight
func loadWaterToLight(scn *bufio.Scanner) {
	scn.Scan()
	scn.Text()
	fmt.Println(" Loading water-to-light map")
	scn.Scan()
	l := scn.Text()
	for l != "" {
		nums := strings.Fields(l)
		var m mapT
		m.dst, _ = strconv.Atoi(nums[0])
		m.src, _ = strconv.Atoi(nums[1])
		m.ran, _ = strconv.Atoi(nums[2])
		toLight = append(toLight, m)
		scn.Scan()
		l = scn.Text()
	}
}

// loadLightToTemperature
func loadLightToTemperature(scn *bufio.Scanner) {
	scn.Scan()
	scn.Text()
	fmt.Println(" Loading light-to-temperature map")
	scn.Scan()
	l := scn.Text()
	for l != "" {
		nums := strings.Fields(l)
		var m mapT
		m.dst, _ = strconv.Atoi(nums[0])
		m.src, _ = strconv.Atoi(nums[1])
		m.ran, _ = strconv.Atoi(nums[2])
		toTemperature = append(toTemperature, m)
		scn.Scan()
		l = scn.Text()
	}
}

// loadTemperatureToHumidity
func loadTemperatureToHumidity(scn *bufio.Scanner) {
	scn.Scan()
	scn.Text()
	fmt.Println(" Loading temperature-to-humidity map")
	scn.Scan()
	l := scn.Text()
	for l != "" {
		nums := strings.Fields(l)
		var m mapT
		m.dst, _ = strconv.Atoi(nums[0])
		m.src, _ = strconv.Atoi(nums[1])
		m.ran, _ = strconv.Atoi(nums[2])
		toHumidity = append(toHumidity, m)
		scn.Scan()
		l = scn.Text()
	}
}

// loadHumidityToLocation
func loadHumidityToLocation(scn *bufio.Scanner) {
	scn.Scan()
	scn.Text()
	fmt.Println(" Loading humidity-to-location map")
	scn.Scan()
	l := scn.Text()
	for l != "" {
		nums := strings.Fields(l)
		var m mapT
		m.dst, _ = strconv.Atoi(nums[0])
		m.src, _ = strconv.Atoi(nums[1])
		m.ran, _ = strconv.Atoi(nums[2])
		toLocation = append(toLocation, m)
		scn.Scan()
		l = scn.Text()
	}
}
