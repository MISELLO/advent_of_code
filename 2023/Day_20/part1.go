package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ButtonPushes is the number of times we push the "Button Module"
const ButtonPushes int = 1000

var lowPulses, highPulses int

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

	// Execution loop
	fmt.Println("Pushing button ...")
	for i := 0; i < ButtonPushes; i++ {
		pushButton()
		for len(messages) > 0 {
			processSignal(popMessage())
		}
	}

	// Result
	fmt.Printf("After %d button pushes, we got %d low pulses and %d high pulses. The product is: \033[1m%d\033[0m\n", ButtonPushes, lowPulses, highPulses, lowPulses*highPulses)
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
func processSignal(s tSignal) {

	// Count the pulse
	if s.highPulse {
		highPulses++
	} else {
		lowPulses++
	}

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
