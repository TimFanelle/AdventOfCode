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
	rep, th := readInput_10("2024/10-input.txt")
	count := 0
	for _, val := range th {
		temp := make(map[loc]bool)
		endLocs := determineUniqueTrails(val, rep)
		for _, v := range endLocs {
			temp[v] = true
		}
		count += len(temp)
	}
	fmt.Println(count)
}

func p2(){
	rep, th := readInput_10("2024/10-input.txt")
	count := 0
	for _, val := range th {
		count += determineNumTrails(val, rep)
	}
	fmt.Println(count)
}

type loc struct {
	x int
	y int
	value int
}

func readInput_10(fileName string) ([][]int, []loc) {
	var representation [][]int;
	var th []loc;

	file, _ := os.Open(fileName)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		temp := convListString2Int(strings.Split(line, ""))
		representation = append(representation, temp)
	}

	for i := 0; i < len(representation); i++{
		for j := 0; j < len(representation[i]); j++{
			if representation[i][j] == 0 {
				th = append(th, loc{x:j, y:i, value: 0})
			}
		}
	}

	return representation, th
}

func determineNumTrails(curLoc loc, representation [][]int) int {
	options := getOptions(curLoc, representation)
	if len(options) == 0 {
		if curLoc.value == 9 {
			return 1
		}
		return 0
	} else {
		count := 0
		for _, op := range options {
			count += determineNumTrails(op, representation)
		}
		return count
	}
}

func determineUniqueTrails(curLoc loc, representation [][]int) []loc {
	options := getOptions(curLoc, representation)
	if len(options) == 0 {
		if curLoc.value == 9 {
			return []loc{curLoc}
		}
		return []loc{}
	}
	var validLocs []loc
	for _, op := range options {
		validLocs = append(validLocs, determineUniqueTrails(op, representation)...)
	}
	return validLocs
}

func getOptions(curLoc loc, representation [][]int) []loc {
	var validLocs []loc;

	locChoices := []loc{
		loc{x:curLoc.x+1, y:curLoc.y},
		loc{x:curLoc.x-1, y:curLoc.y},
		loc{x:curLoc.x, y:curLoc.y+1},
		loc{x:curLoc.x, y:curLoc.y-1},
	}


	for _, l := range locChoices {
		if l.x >=0 && l.x < len(representation[0]) && l.y >=0 && l.y < len(representation){
			l.value = representation[l.y][l.x]
			if l.value == curLoc.value +1 {
				validLocs = append(validLocs, l)
			}
		}
	}
	return validLocs
}

func convListString2Int(valIn []string) []int {
	var out []int;
	for i := 0; i < len(valIn); i++{
		out = append(out, convString2Int(valIn[i]))
	}
	return out
}

func convString2Int(s string) int {
	val, _ := strconv.Atoi(s)
	return val
}