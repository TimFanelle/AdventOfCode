package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

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
	
	var difVal float64 = 0;

	for i := 0; i < len(left); i++{
		j := float64(left[i]-right[i])
		difVal += math.Abs(j)
	}
	fmt.Println(difVal)
}