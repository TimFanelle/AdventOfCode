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
	stones := readInput_11("2024/11-input.txt")
	fmt.Println(stones)
	curStones := stones
	for i := 0; i < 25; i++ {
		curStones = blink(curStones)
	}
	fmt.Println(len(curStones))
}

func p2() {
	stones := readInput_11("2024/11-input.txt")
	fmt.Println(stones)
	curStones := convListStringToMap(stones)
	for i := 0; i < 75; i++ {
		curStones = p2Blink(curStones)
	}

	count := 0
	for _, v := range curStones {
		count += v
	}
	fmt.Println(count)
}

func readInput_11(fileName string) []string {
	var stones []string;

	file, _ := os.Open(fileName)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		temp := strings.Split(line, " ")
		stones = append(stones, temp...)
	}
	return stones
}

func blink(curStones []string) []string {
	var stonesOut []string

	for _, val := range curStones {
		if len(val) % 2 == 0 {
			temp := []string{val[:len(val)/2], val[len(val)/2:]}
			stonesOut = append(stonesOut, temp...)
		} else if val == "0" {
			stonesOut = append(stonesOut, "1")
		} else {
			stonesOut = append(stonesOut, convString2IntAndBackWMult(val, 2024))
		}
	}
	stonesOut = correctZeroes(stonesOut)
	return stonesOut
}

func p2Blink(curStones map[string]int) map[string]int {
	nextStones := make(map[string]int) 

	for k := range curStones {
		if len(k) % 2 == 0 {
			temp := []string{k[:len(k)/2], k[len(k)/2:]}
			for _, j := range temp {
				for len(j) > 1 && j[0:1] == "0"{
					j = j[1:]
				}
				nextStones[j] += 1*curStones[k]
			}
		} else if k == "0" {
			nextStones["1"] += 1*curStones[k]
		} else {
			j := convString2IntAndBackWMult(k, 2024)
			for len(j) > 1 && j[0:1] == "0"{
				j = j[1:]
			}
			nextStones[j] += 1*curStones[k]
		}
	}
	return nextStones
}

func convString2IntAndBackWMult(s string, multVal int) string {
	v, _ := strconv.Atoi(s)
	t := v*multVal
	out := strconv.Itoa(t)
	return out
}

func correctZeroes(lIn []string) []string {
	var out []string
	for _, val := range lIn {
		temp := val
		for len(temp) > 1 && temp[0:1] == "0"{
			temp = temp[1:]
		}
		
		out = append(out, temp)
	}
	return out
}

func convListStringToMap(s []string) map[string]int {
	out := make(map[string]int)

	for _, val := range s {
		out[val] += 1
	}
	return out
}