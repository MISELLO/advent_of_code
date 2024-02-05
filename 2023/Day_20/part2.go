package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var buttonPushes int
var end bool

type tModuletype int

const (
	flip_flop tModuletype = iota
	conjunction
	broadcaster
)

type tModule struct {
	mType  tModuletype
	input  map[string]bool
	output []string
}

var moduleList []tModule

var moduleIndex map[string]int
var moduleCicle map[string][]int

type tSignal struct {
	src, dst  string
	highPulse bool
}

var messages []tSignal

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
	fmt.Println("Loading modules ...")
	moduleIndex = make(map[string]int)

	for scn.Scan() {
		l := scn.Text()
		parseModuleConfiguration(l)
	}

	lastModuleName, lastModule := revealLastModule()
	fmt.Printf("The module that feeds \"%s\" is \"%s\".\n", "rx", lastModuleName)

	fmt.Printf("\"%s\" will send a low pulse to \"rx\" when all of it's inputs are high.\n", lastModuleName)
	list := getInputLabels(lastModule.input)
	fmt.Println("This inputs are:", list)
	fmt.Println("We can assume they output a high pulse once in a cicle.")
	fmt.Println("If we find the cicle of each of them we will be able to calculate when all output a high pulse at the same time.")
	moduleCicle = make(map[string][]int)

	// Search loop
	fmt.Println("Searching for cicles ...")
	for !end {
		pushButton()
		buttonPushes++
		for len(messages) > 0 {
			processSignal(lastModule.input, popMessage())
		}
		end = ciclesFound(lastModule.input)
	}

	// Result
	res := computeWhenAllCiclesMatch()
	fmt.Printf("After \033[1m%d\033[0m button pushes, we will get a low pulse on the \"rx\" module.\n", res)
}

// parseModuleConfiguration reads a string and stores the module configuration in our structures
func parseModuleConfiguration(s string) {
	parts := strings.Split(s, " -> ")
	moduleRaw := parts[0]
	outputsRaw := parts[1]

	// What name and type?
	var m tModule
	var moduleLabel string
	if moduleRaw[0] == '%' {
		m.mType = flip_flop
		moduleLabel = moduleRaw[1:]
	} else if moduleRaw[0] == '&' {
		m.mType = conjunction
		moduleLabel = moduleRaw[1:]
	} else {
		m.mType = broadcaster
		moduleLabel = moduleRaw
	}

	// Does it exist? At what index?
	existingI, exists := moduleIndex[moduleLabel]

	// Assign mType
	if exists {
		moduleList[existingI].mType = m.mType
	}

	// input
	if exists && m.mType == flip_flop {
		moduleList[existingI].input = make(map[string]bool)
		moduleList[existingI].input["_STATUS_"] = false
	} else if !exists && m.mType == flip_flop {
		m.input = make(map[string]bool)
		m.input["_STATUS_"] = false
	} else if !exists && m.mType == conjunction {
		m.input = make(map[string]bool)
	}

	// output
	if exists {
		moduleList[existingI].output = strings.Split(outputsRaw, ", ")
	} else {
		m.output = strings.Split(outputsRaw, ", ")
	}

	// Adding module
	if !exists {
		moduleList = append(moduleList, m)
		moduleIndex[moduleLabel] = len(moduleList) - 1
	}

	// Adding output modules to the module list (and setting this current module as their input)
	for i := 0; i < len(m.output); i++ {
		modInd, e := moduleIndex[m.output[i]]
		if !e {
			var mod tModule
			mod.input = make(map[string]bool)
			mod.input[moduleLabel] = false
			moduleList = append(moduleList, mod)
			moduleIndex[m.output[i]] = len(moduleList) - 1
		} else {
			moduleList[modInd].input[moduleLabel] = false
		}
	}
}

// pushButton sends a low pulse to the broadcaster module
func pushButton() {
	signal := tSignal{src: "_BUTTON_", dst: "broadcaster", highPulse: false}
	messages = append(messages, signal)
}

// popMessage takes out the oldest signal from the messages list
// and it returns it's value.
func popMessage() tSignal {
	s := messages[0]
	messages = messages[1:]
	return s
}

// processSignal simulates the efect on the modules that will have a given signal
// including putting more signals on the messages queue.
// Also, checks if the signal comes from a module inside a given list (moduleInput)
// and stores the button pushes done when it outputs a high pulse
func processSignal(moduleInput map[string]bool, s tSignal) {

	// Checks if the signal comes from one of the list.
	for mName, _ := range moduleInput {
		if mName == s.src && s.highPulse {
			moduleCicle[s.src] = append(moduleCicle[s.src], buttonPushes)
		}
	}

	// The module we will work on
	mIndex := moduleIndex[s.dst]
	m := moduleList[mIndex]

	// Broadcaster
	if m.mType == broadcaster {
		for i := 0; i < len(m.output); i++ {
			signal := tSignal{src: s.dst, dst: m.output[i], highPulse: s.highPulse}
			messages = append(messages, signal)
		}
	}

	// Flip-flop (%)
	if m.mType == flip_flop {
		if !s.highPulse {
			// Low pulse
			m.input["_STATUS_"] = !m.input["_STATUS_"]
			for i := 0; i < len(m.output); i++ {
				signal := tSignal{src: s.dst, dst: m.output[i], highPulse: m.input["_STATUS_"]}
				messages = append(messages, signal)
			}
		}
	}

	// Conjunction (&)
	if m.mType == conjunction {

		// Update memory
		m.input[s.src] = s.highPulse

		// Check if all are high pulses
		allHighPulses := true
		for _, v := range m.input {
			if v == false {
				allHighPulses = false
				break
			}
		}
		var answer bool
		if allHighPulses {
			answer = false
		} else {
			answer = true
		}

		// Sends the answer to it's output
		for i := 0; i < len(m.output); i++ {
			signal := tSignal{src: s.dst, dst: m.output[i], highPulse: answer}
			messages = append(messages, signal)
		}
	}
}

// revealLastModule checks, reveals and returns the last module of the graph
// (the one that feeds the "rx" module)
func revealLastModule() (string, tModule) {
	// We know this will end once a low pulse reaches the "rx" module
	// After checking the input file we notice only a conjunction module can send a pulse to "rx"
	// The following code searches this module.
	for label, index := range moduleIndex {
		m := moduleList[index]
		for i := 0; i < len(m.output); i++ {
			// Technically the input list we are looking for
			// only has 1 element.
			if m.output[i] == "rx" {
				return label, m
			}
		}
	}
	return "", tModule{}
}

// getInputLabels gets a module input as parameter and returns
// a list (string) with the labels of all the input modules
func getInputLabels(moduleInput map[string]bool) string {
	var l string
	for k, _ := range moduleInput {
		if len(l) != 0 {
			l += ", " + k
		} else {
			l = k
		}
	}
	return l
}

// ciclesFound checks if we have enough data to know the cicles
// of the needed modules
func ciclesFound(input map[string]bool) bool {
	if len(moduleCicle) != len(input) {
		return false
	}
	for _, l := range moduleCicle {
		if len(l) < 2 {
			return false
		}
	}
	return true
}

// computeWhenAllCiclesMatch
func computeWhenAllCiclesMatch() int {
	var cicles []int
	for name, l := range moduleCicle {
		n := len(l) - 1
		preN := len(l) - 2
		fmt.Printf("The cicle of \"%s\" is %d.\n", name, l[n]-l[preN])
		cicles = append(cicles, l[n]-l[preN])
	}

	var result int
	result = cicles[0]
	for i := 1; i < len(cicles); i++ {
		result = LCM(result, cicles[i])
	}

	return result
}

// EuclidesGCD is the Euclides algorithm for the Greater Common Divisor
func EuclidesGCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// LCM computes the Lesser Common Divisor algorithm following the formula
// LCM(a,b) = a * b / GCD(a,b)
func LCM(a, b int) int {
	return a * b / EuclidesGCD(a, b)
}
