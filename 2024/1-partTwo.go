package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)
func count[T any](slice []T, f func(T) bool) int{
	count := 0
	for _, s := range slice {
		if f(s){
			count++
		}
	}
	return count
}

func main(){
	var left []int;
	var right []int;

	file, err := os.Open("2024/1-input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		temp := scanner.Text()
		line := strings.Split(temp, "   ")
		
		t, e := strconv.Atoi(line[0])
		if e == nil{
			left = append(left, t)
		}
		
		t, e = strconv.Atoi(line[1])
		if e == nil{
			right = append(right, t)
		}
	}
	sort.Ints(left)
	sort.Ints(right)
	
	for i := 0; i<len(left); i++{
		left[i] *= count(right, func(num int) bool {
			return num == left[i]
		})
	}

	var endVal int = 0;

	for i := 0; i < len(left); i++{
		endVal += left[i]
	}
	fmt.Println(endVal)
}