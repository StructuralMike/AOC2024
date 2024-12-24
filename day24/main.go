package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Gate struct {
	operator string
	inputs   inputWires
}

type Wire struct {
	outputGate *Gate
	value      string
}

type inputWires []*Wire

func solvePart1(inputFile string) uint64 {
	wires := loadDayData(inputFile)

	var output [65]int
	for name, wire := range wires {
		if !strings.HasPrefix(name, "z") {
			continue
		}
		bitStr := strings.Split(name, "z")[1]
		bitNum, err := strconv.Atoi(bitStr)
		if err != nil {
			log.Fatal(err)
		}

		setWireValue(wire)
		bitVal, err := strconv.Atoi(wire.value)
		if err != nil {
			log.Fatal(err)
		}

		output[64-bitNum] = bitVal
	}

	var result uint64 = 0
	for _, bit := range output {
		result = result<<1 | uint64(bit)
	}

	return result
}

func solvePart2(inputFile string) uint64 {
	wires := loadDayData(inputFile)

	var output [65]int
	for name, wire := range wires {
		if !strings.HasPrefix(name, "z") {
			continue
		}
		bitStr := strings.Split(name, "z")[1]
		bitNum, err := strconv.Atoi(bitStr)
		if err != nil {
			log.Fatal(err)
		}

		setWireValue(wire)
		bitVal, err := strconv.Atoi(wire.value)
		if err != nil {
			log.Fatal(err)
		}

		output[64-bitNum] = bitVal
	}

	var result uint64 = 0
	for _, bit := range output {
		result = result<<1 | uint64(bit)
	}

	return result
}

func setWireValue(wire *Wire) {
	if wire.value == "1" || wire.value == "0" {
		return
	}
	gate := wire.outputGate
	for _, inputWire := range gate.inputs {
		setWireValue(inputWire)
	}
	wire1 := gate.inputs[0]
	wire2 := gate.inputs[1]
	var outputState bool
	switch gate.operator {
	case "OR":
		outputState = wire1.value == "1" || wire2.value == "1"
	case "XOR":
		outputState = wire1.value != wire2.value
	case "AND":
		outputState = wire1.value == "1" && wire2.value == "1"
	}
	if outputState {
		wire.value = "1"
	} else {
		wire.value = "0"
	}
}

func loadDayData(inputFile string) map[string]*Wire {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	var wires = make(map[string]*Wire)
	isInputs := true
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) < 2 {
			isInputs = false
			continue
		}

		if isInputs {
			lineSplit := strings.Split(strings.Replace(line, " ", "", -1), ":")
			wireName := lineSplit[0]
			wireState := lineSplit[1]
			wires[wireName] = &Wire{value: wireState}
		}

		if !isInputs {
			lineSplit := strings.Split(strings.Replace(line, "-> ", "", -1), " ")
			wire1Name := lineSplit[0]
			gateType := lineSplit[1]
			wire2Name := lineSplit[2]
			outputWireName := lineSplit[3]

			newGate := Gate{operator: gateType}

			_, ok := wires[wire1Name]
			if !ok {
				wires[wire1Name] = &Wire{}
			}

			_, ok = wires[wire2Name]
			if !ok {
				wires[wire2Name] = &Wire{}
			}

			_, ok = wires[outputWireName]
			if !ok {
				wires[outputWireName] = &Wire{outputGate: &newGate}
			} else {
				wires[outputWireName].outputGate = &newGate
			}

			newGate.inputs = append(newGate.inputs, wires[wire1Name])
			newGate.inputs = append(newGate.inputs, wires[wire2Name])
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return wires
}

func main() {
	start := time.Now()

	// fmt.Println(solvePart1("sample_input.txt"))
	// fmt.Println(solvePart1("input1.txt"))
	fmt.Println(solvePart2("input1.txt"))

	fmt.Println(time.Since(start))
}
