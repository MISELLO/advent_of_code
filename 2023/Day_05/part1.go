package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	fmt.Println("Almanac loaded")
	fmt.Println(" - There are", len(seedsList), "seeds.")
	fmt.Println(" - The toSoil map has", len(toSoil), "elements")
	fmt.Println(" - The toFertilizer map has", len(toFertilizer), "elements")
	fmt.Println(" - The toWater map has", len(toWater), "elements")
	fmt.Println(" - The toLight map has", len(toLight), "elements")
	fmt.Println(" - The toTemperature map has", len(toTemperature), "elements")
	fmt.Println(" - The toHumidity map has", len(toHumidity), "elements")
	fmt.Println(" - The toLocation map has", len(toLocation), "elements")
	fmt.Println("")
	var lowestLocation int
	for i := 0; i < len(seedsList); i++ {
		s := seedsList[i]
		t := convert(s, toSoil)
		u := convert(t, toFertilizer)
		v := convert(u, toWater)
		w := convert(v, toLight)
		x := convert(w, toTemperature)
		y := convert(x, toHumidity)
		z := convert(y, toLocation)
		if lowestLocation == 0 || z < lowestLocation {
			lowestLocation = z
		}
		fmt.Printf(" — Seed %d, soil %d, fertilizer %d, water %d, light %d, temperature %d, humidity %d, location %d.\n", s, t, u, v, w, x, y, z)
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
