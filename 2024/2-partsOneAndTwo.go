package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main(){
	//p1()
	p2()
}

func p1(){
	var reports [][]int
	var numSafe int = 0
	file, err := os.Open("2024/2-input.txt")
	
	if err != nil{
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		temp := scanner.Text()
		line := strings.Split(temp, " ")

		iLine := sList2iList(line)
		reports = append(reports, iLine)
		if checkSafe(iLine){
			numSafe += 1
		}

	}
	fmt.Println(numSafe)
}
func p2(){
	var reports [][]int
	var numSafe int = 0
	file, err := os.Open("2024/2-input.txt")
	
	if err != nil{
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		temp := scanner.Text()
		line := strings.Split(temp, " ")

		iLine := sList2iList(line)
		reports = append(reports, iLine)
		if checkSafeDampered(iLine){
			numSafe += 1
		}

	}
	fmt.Println(numSafe)
}


func Abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
func listWRemoved(slice []int, index int) []int{
	var outList []int;
	for i := 0; i <len(slice); i++{
		if i != index{
			outList = append(outList, slice[i])
		}
	}
	return outList
}

func sList2iList(strIn []string) []int {
	var outList []int;
	for i := 0; i < len(strIn); i++{
		t, _ := strconv.Atoi(strIn[i])
		outList = append(outList, t)
	}
	return outList
}

func checkSafe(report []int) bool {
	var safe bool = checkIncreaseDecrease(report) && checkSkipSize(report)
	return safe
}

func checkSafeDampered(report []int) bool {
	var safe bool = checkIncreaseDecrease(report) && checkSkipSize(report)
	if safe {
		return true
	} else {
		var lCheck []int; 
		for j := 0; j < len(report); j++{
			lCheck = listWRemoved(report, j)
			if checkSafe(lCheck){
				return true
			}
		}
		return false
	}
}

func checkIncreaseDecrease(report []int) bool{
	var allIncrease bool = true
	var allDecrease bool = true
	for i := 1; i <len(report); i++{
		if report[i] < report[i-1]{
			allIncrease = false
		}
		if report[i] > report[i-1]{
			allDecrease = false
		}
	}
	return allIncrease || allDecrease
}

func checkSkipSize(report []int) bool {
	var safe bool = true
	for i := 1; i < len(report); i++{
		if Abs(report[i]-report[i-1]) > 3 || Abs(report[i]-report[i-1]) < 1{
			safe = false
		}
	}
	return safe
}