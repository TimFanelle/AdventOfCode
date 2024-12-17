package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	//p1()
	p2()
}

func p1() {
	reg, instruct := readInput_17("2024/17-input.txt")
	fmt.Println(reg)
	fmt.Println(instruct)

	outputs := runInstructions(instruct, reg)
	
	fmt.Println(outputs)
	outStr := ""
	for _, v := range outputs {
		outStr += strconv.Itoa(v) + ","
	}
	fmt.Println(outStr[:len(outStr)-1])
}

func p2(){
	_, instruct := readInput_17("2024/17-input.txt")
	fmt.Println(determineA(instruct))
	
}

func readInput_17(fileName string) (map[string]int, []int) {
	registers := make(map[string]int)
	
	file, _ := os.Open(fileName)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0{
			l := strings.Split(line, ": ")
			lines = append(lines, l[1])
		}
	}

	registers["A"] = convString2Int(lines[0])
	registers["B"] = convString2Int(lines[1])
	registers["C"] = convString2Int(lines[2])

	instructions := []int{}
	for _, v := range strings.Split(lines[len(lines)-1], ","){
		instructions = append(instructions, convString2Int(v))
	}
	return registers, instructions
}

func convString2Int(s string) int {
	val, _ := strconv.Atoi(s)
	return val
}

func doInstruction(pointer int, instructions []int, reg map[string]int) int {
	op := instructions[pointer]
	combo := instructions[pointer+1]

	switch op {
	case 0:
		outcome := 0.0
		switch {
			case combo <= 3:
				outcome = float64(reg["A"])/math.Pow(2, float64(combo))
			case combo == 4:
				outcome = float64(reg["A"])/math.Pow(2, float64(reg["A"]))
			case combo == 5:
				outcome = float64(reg["A"])/math.Pow(2, float64(reg["B"]))
			case combo == 6:
				outcome = float64(reg["A"])/math.Pow(2, float64(reg["C"]))
		}
		reg["A"] = int(math.Trunc(float64(outcome)))
		return -1
		
	case 1:
		outcome := reg["B"] ^ combo
		reg["B"] = outcome
		return -1

	case 2:
		outcome := 0
		switch {
			case combo <= 3:
				outcome = combo % 8
			case combo == 4:
				outcome = reg["A"] % 8
			case combo == 5:
				outcome = reg["B"] % 8
			case combo == 6:
				outcome = reg["C"] % 8
		}
		reg["B"] = outcome
		return -1
	case 3:
		if reg["A"] != 0 {
			return combo
		}
		return -1
	case 4:
		outcome := reg["B"] ^ reg["C"]
		reg["B"] = outcome
		return -1
	case 5:
		outcome := 0
		switch {
		case combo <= 3:
			outcome = combo % 8
		case combo == 4:
			outcome = reg["A"] % 8
		case combo == 5:
			outcome = reg["B"] % 8
		case combo == 6:
			outcome = reg["C"] % 8
		}
		return (outcome*-1) - 3
	case 6:
		outcome := 0.0
		switch {
			case combo <= 3:
				outcome = float64(reg["A"])/math.Pow(2, float64(combo))
			case combo == 4:
				outcome = float64(reg["A"])/math.Pow(2, float64(reg["A"]))
			case combo == 5:
				outcome = float64(reg["A"])/math.Pow(2, float64(reg["B"]))
			case combo == 6:
				outcome = float64(reg["A"])/math.Pow(2, float64(reg["C"]))
		}
		reg["B"] = int(math.Trunc(float64(outcome)))
		return -1
	case 7:
		outcome := 0.0
		switch {
			case combo <= 3:
				outcome = float64(reg["A"])/math.Pow(2, float64(combo))
			case combo == 4:
				outcome = float64(reg["A"])/math.Pow(2, float64(reg["A"]))
			case combo == 5:
				outcome = float64(reg["A"])/math.Pow(2, float64(reg["B"]))
			case combo == 6:
				outcome = float64(reg["A"])/math.Pow(2, float64(reg["C"]))
		}
		reg["C"] = int(math.Trunc(float64(outcome)))
		return -1
	default:
		// do nothing
	}
	return -1
}

func runInstructions(instruct []int, reg map[string]int ) []int {
	outputs := []int{}
	i := 0
	for i < len(instruct){
		nP := doInstruction(i, instruct, reg)
		if nP == -1 {
			i += 2
		} 
		if nP < -1 {
			outputs = append(outputs, (nP+3)*-1)
			i += 2
		}
		if nP > -1 {
			i = nP
		} 

	}
	return outputs
}

func determineA(instructs []int) int {
	a := 0
	working := make(map[string]int)
	for i := len(instructs)-1; i >=0; i--{
		a = a << 3
		working["A"] = a
		working["B"] = 0
		working["C"] = 0

		for !slices.Equal(runInstructions(instructs, working), instructs[i:]) {
			a += 1
			working["A"] = a
			working["B"] = 0
			working["C"] = 0

		}
	}
	return a
}
