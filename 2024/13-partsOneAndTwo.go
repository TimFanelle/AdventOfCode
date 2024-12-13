package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"gonum.org/v1/gonum/mat"
)

func main(){
	p1()
	p2()
}

func p1(){
	buttonsAndPrizes := readInput_13("13-input.txt")
	tokens := 0
	for _, v := range buttonsAndPrizes{
		checkButtons := solveSystem(v)

		if roundFloat(checkButtons[0], 4) == math.Round(checkButtons[0]) && roundFloat(checkButtons[1], 4) == math.Round(checkButtons[1]){
			tokens += int(math.Round(checkButtons[0]))*3 + int(math.Round(checkButtons[1]))*1
		}
	}
	fmt.Println(tokens)
}
func p2(){
	buttonsAndPrizes := readInput_13("13-input.txt")
	delta := 10000000000000
	result := 0
	for _, game := range buttonsAndPrizes {
		result += getPriceByEquation(game, delta)
	}
	fmt.Println(result)
}

func readInput_13(fileName string) []buttonPrizes {
	var buttonsAndPrizes []buttonPrizes;

	file, _ := os.Open(fileName)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string;
	numLines := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" && line != " " && line != "\n"{
			lines = append(lines, line)
			numLines++
			if numLines == 3 {
				numLines = 0
				lineA := strings.Split(strings.Split(lines[0], ": ")[1], ", ")
				lineB := strings.Split(strings.Split(lines[1], ": ")[1], ", ")
				lineC := strings.Split(strings.Split(lines[2], ": ")[1], ", ")

				bp := buttonPrizes{aButton: []int{convString2Int(lineA[0][1:]),convString2Int(lineA[1][1:])}, bButton: []int{convString2Int(lineB[0][1:]),convString2Int(lineB[1][1:])}, prizeLoc: []int{convString2Int(lineC[0][2:]),convString2Int(lineC[1][2:])}}
				buttonsAndPrizes = append(buttonsAndPrizes, bp)
				lines = []string{}
			}
		}
	}
	return buttonsAndPrizes
}

type buttonPrizes struct {
	aButton []int
	bButton []int
	prizeLoc []int
}

func convString2Int(s string) int {
	val, _ := strconv.Atoi(s)
	return val
}

func solveSystem(buttonVals buttonPrizes) []float64 {
	A := mat.NewDense(2, 2, []float64{
		float64(buttonVals.aButton[0]), float64(buttonVals.bButton[0]),
		float64(buttonVals.aButton[1]), float64(buttonVals.bButton[1]),
	})

	B := mat.NewVecDense(2,[]float64{
		float64(buttonVals.prizeLoc[0]),
		float64(buttonVals.prizeLoc[1]),
	})
	
	var x mat.VecDense
	if err := x.SolveVec(A, B); err != nil {
		fmt.Println(err)
	}
	return x.RawVector().Data
}

func roundFloat(val float64, precision uint) float64 {
    ratio := math.Pow(10, float64(precision))
    return math.Round(val*ratio) / ratio
}

func getPriceByEquation(game buttonPrizes, delta int) int {
	// This function pulled from https://github.com/omotto/AdventOfCode2024/blob/main/src/day13/main.go
	/*
		2x2 System Linear Equations
		a * aX + b * bX = pX
		a * aY + b * bY = pY
	*/
	pX := game.prizeLoc[0] + delta
	pY := game.prizeLoc[1] + delta
	aX := game.aButton[0]
	aY := game.aButton[1]
	bX := game.bButton[0]
	bY := game.bButton[1]

	a := float64(pX*bY-pY*bX) / float64(aX*bY-aY*bX)
	b := float64(pY*aX-pX*aY) / float64(aX*bY-aY*bX)

	// if there is no decimals is valid
	if a == math.Trunc(a) && b == math.Trunc(b) {
		return int(a*3 + b)
	}
	return 0
}