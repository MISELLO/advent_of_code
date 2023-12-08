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
	seedsEnd
	soil
	soilEnd
	fertilizer
	fertilizerEnd
	water
	waterEnd
	light
	lightEnd
	temperature
	temperatureEnd
	humidity
	humidityEnd
	location
	locationEnd
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
			s = seeds
			fmt.Println(" Loading seeds")
			parts := strings.Split(l, ":")
			seedNums := strings.Fields(parts[1])
			for i := 0; i < len(seedNums); i++ {
				n, _ := strconv.Atoi(seedNums[i])
				seedsList = append(seedsList, n)
			}
			s = seedsEnd
		} else if s == seedsEnd || s == soil {
			// Soil
			if s == seedsEnd {
				if strings.Contains(l, "seed-to-soil") {
					s = soil
					fmt.Println(" Loading seed-to-soil map")
				}
			} else if s == soil {
				if l != "" {
					nums := strings.Fields(l)
					var m mapT
					m.dst, _ = strconv.Atoi(nums[0])
					m.src, _ = strconv.Atoi(nums[1])
					m.ran, _ = strconv.Atoi(nums[2])
					toSoil = append(toSoil, m)
				} else {
					s = soilEnd
				}
			}
		} else if s == soilEnd || s == fertilizer {
			// Fertilizer
			if s == soilEnd {
				if strings.Contains(l, "soil-to-fertilizer") {
					s = fertilizer
					fmt.Println(" Loading soil-to-fertilizer map")
				}
			} else if s == fertilizer {
				if l != "" {
					nums := strings.Fields(l)
					var m mapT
					m.dst, _ = strconv.Atoi(nums[0])
					m.src, _ = strconv.Atoi(nums[1])
					m.ran, _ = strconv.Atoi(nums[2])
					toFertilizer = append(toFertilizer, m)
				} else {
					s = fertilizerEnd
				}
			}
		} else if s == fertilizerEnd || s == water {
			// Water
			if s == fertilizerEnd {
				if strings.Contains(l, "fertilizer-to-water") {
					s = water
					fmt.Println(" Loading fertilizer-to-water map")
				}
			} else if s == water {
				if l != "" {
					nums := strings.Fields(l)
					var m mapT
					m.dst, _ = strconv.Atoi(nums[0])
					m.src, _ = strconv.Atoi(nums[1])
					m.ran, _ = strconv.Atoi(nums[2])
					toWater = append(toWater, m)
				} else {
					s = waterEnd
				}
			}
		} else if s == waterEnd || s == light {
			// Light
			if s == waterEnd {
				if strings.Contains(l, "water-to-light") {
					s = light
					fmt.Println(" Loading water-to-light map")
				}
			} else if s == light {
				if l != "" {
					nums := strings.Fields(l)
					var m mapT
					m.dst, _ = strconv.Atoi(nums[0])
					m.src, _ = strconv.Atoi(nums[1])
					m.ran, _ = strconv.Atoi(nums[2])
					toLight = append(toLight, m)
				} else {
					s = lightEnd
				}
			}
		} else if s == lightEnd || s == temperature {
			// Temperature
			if s == lightEnd {
				if strings.Contains(l, "light-to-temperature") {
					s = temperature
					fmt.Println(" Loading light-to-temperature map")
				}
			} else if s == temperature {
				if l != "" {
					nums := strings.Fields(l)
					var m mapT
					m.dst, _ = strconv.Atoi(nums[0])
					m.src, _ = strconv.Atoi(nums[1])
					m.ran, _ = strconv.Atoi(nums[2])
					toTemperature = append(toTemperature, m)
				} else {
					s = temperatureEnd
				}
			}
		} else if s == temperatureEnd || s == humidity {
			// Humidity
			if s == temperatureEnd {
				if strings.Contains(l, "temperature-to-humidity") {
					s = humidity
					fmt.Println(" Loading temperature-to-humidity map")
				}
			} else if s == humidity {
				if l != "" {
					nums := strings.Fields(l)
					var m mapT
					m.dst, _ = strconv.Atoi(nums[0])
					m.src, _ = strconv.Atoi(nums[1])
					m.ran, _ = strconv.Atoi(nums[2])
					toHumidity = append(toHumidity, m)
				} else {
					s = humidityEnd
				}
			}
		} else if s == humidityEnd || s == location {
			// Location
			if s == humidityEnd {
				if strings.Contains(l, "humidity-to-location") {
					s = location
					fmt.Println(" Loading humidity-to-location map")
				}
			} else if s == location {
				if l != "" {
					nums := strings.Fields(l)
					var m mapT
					m.dst, _ = strconv.Atoi(nums[0])
					m.src, _ = strconv.Atoi(nums[1])
					m.ran, _ = strconv.Atoi(nums[2])
					toLocation = append(toLocation, m)
				} else {
					s = locationEnd
				}
			}
		}
	}
	fmt.Println("Almanac loaded")
	fmt.Println(" - The seeds are:", seedsList)
	fmt.Println(" - The toSoil map is:", toSoil)
	fmt.Println(" - The toFertilizer map is:", toFertilizer)
	fmt.Println(" - The toWater map is:", toWater)
	fmt.Println(" - The toLight map is:", toLight)
	fmt.Println(" - The toTemperature map is:", toTemperature)
	fmt.Println(" - The toHumidity map is:", toHumidity)
	fmt.Println(" - The toLocation map is:", toLocation)
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
