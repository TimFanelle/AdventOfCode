package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"regexp"
)

func main(){
	//p1()
	p2()
}

func p1(){
	var total int = 0;
	file, _ := os.Open("2024/3-input.txt")
	defer file.Close()

	re := regexp.MustCompile("mul\\([0-9]+,[0-9]+\\)")
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		line := scanner.Text()
		found := re.FindAllString(line, -1)
		for i := 0; i < len(found); i++{
			total += interpretMult(found[i])
		}
	}
	fmt.Println(total)
}
func p2(){
	var total int = 0;
	var do_yes bool = true;
	file, _ := os.Open("2024/3-input.txt")
	defer file.Close()

	re := regexp.MustCompile("(do\\(\\))|(mul\\([0-9]+,[0-9]+\\))|(don't\\(\\))")
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		line := scanner.Text()
		found := re.FindAllString(line, -1)
		for i := 0; i < len(found); i++{
			if len(found[i]) == 4{
				do_yes = true
			} else if found[i][:5] == "don't" {
				do_yes = false
			} else {
				if do_yes {
					total += interpretMult(found[i])
				}
			}	
		}
	}
	fmt.Println(total)
}

func valToInt(s string) int {
	out, _ := strconv.Atoi(s)
	return out
}

func interpretMult(input string) int {
	working := input[4:len(input)-1]
	temp := strings.Split(working, ",")
	first := valToInt(temp[0])
	second := valToInt(temp[1])
	return first*second
}