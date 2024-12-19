package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var memo = make(map[string]bool)
var memoCount = make(map[string]int)

func main(){
	//p1()
	p2()
}

func p1(){
	availPats, desiredPats := readInput_19("2024/19-input.txt")
	count := 0
	for _, v := range desiredPats{
		if attemptToMakePattern(v, availPats) {
			count += 1
		}
	}
	fmt.Println(count)
}
func p2(){
	availPats, desiredPats := readInput_19("2024/19-input.txt")
	count := 0
	for _, v := range desiredPats{
		count += countToMakePattern(v, availPats)
	}
	fmt.Println(count)
}

func readInput_19(fileName string) ([]string, []string) {
	avail := []string{}
	desired := []string{}

	file, _ := os.Open(fileName)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	all_lines := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 && line != " " && line != "\n"{
			all_lines = append(all_lines, line)
		}
	}
	avail = strings.Split(all_lines[0], ", ")
	desired = all_lines[1:]

	return avail, desired
}

func attemptToMakePattern(desired string, avail []string) bool {
	if val, found := memo[desired]; found {
		return val
	}
	for _, av := range avail{
		if len(desired) >= len(av) && desired[:len(av)] == av{
			if len(desired) == len(av){
				memo[desired] = true
				return memo[desired]
			} else {
				check := attemptToMakePattern(desired[len(av):], avail)
				if check {
					memo[desired] = true
					return memo[desired]
				}
			}
		}
	}
	memo[desired] = false
	return memo[desired]
}

func countToMakePattern(desired string, avail []string) int {
	if val, found := memoCount[desired]; found {
		return val
	}
	poss := 0
	for _, av := range avail {
		if len(desired) >= len(av) && desired[:len(av)] == av{
			if len(desired) == len(av){
				poss += 1
			} else {
				poss += countToMakePattern(desired[len(av):], avail)
			}
		}

	}
	memoCount[desired] = poss
	return poss
}