package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	//p1()
	p2()
}

func p1(){
	var total int = 0;
	configs := readInput_7("2024/7-input.txt")
	for i := 0; i<len(configs); i++{
		if canMakeNum(configs[i][0], configs[i][1], configs[i][2:], false){
			total += configs[i][0]
		}
	}
	
	fmt.Println(total)
}

func p2(){
	var total int = 0;
	configs := readInput_7("2024/7-input.txt")
	for i := 0; i<len(configs); i++{
		if canMakeNum(configs[i][0], configs[i][1], configs[i][2:], true){
			total += configs[i][0]
		}
	}
	
	fmt.Println(total)
}


func canMakeNum(targetVal int, curVal int, valIn []int, part2 bool) bool {
	if len(valIn) == 0 {
		return curVal == targetVal
	}
	if curVal > targetVal {
		return false
	}
	if canMakeNum(targetVal, performOp(curVal, valIn[0], "+"), valIn[1:], part2){
		return true
	}
	if canMakeNum(targetVal, performOp(curVal, valIn[0], "*"), valIn[1:], part2){
		return true
	}
	if part2 && canMakeNum(targetVal, performOp(curVal, valIn[0], "||"), valIn[1:], part2){
		return true
	}
	return false
}

func performOp(a int, b int, operation string) int {
	switch operation {
		case "+":
			return a + b
		case "*":
			return a*b
		case "||":
			s_a := strconv.Itoa(a)
			s_b := strconv.Itoa(b)
			return convToInt(s_a+s_b)
	}
	return 0
}


func readInput_7(inputName string) [][]int{
	var configs [][]int;
	file, _ := os.Open(inputName)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lineBreak := strings.Split(line, ": ")
		configNum := convToInt(lineBreak[0])
		v := strings.Split(lineBreak[1], " ")
		var temp []int;
		temp = append(temp, configNum)
		for i := 0; i < len(v); i++ {
			temp = append(temp, convToInt(v[i]))
		}
		configs = append(configs, temp)
	}
	return configs
}

func convToInt(input string) int {
	val, _ := strconv.Atoi(input)
	return val
}