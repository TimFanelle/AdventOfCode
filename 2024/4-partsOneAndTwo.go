package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main(){
	//p1()
	p2()
}

func p1(){
	var grid [][]string;
	file, _ := os.Open("2024/4-input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		line := scanner.Text()
		temp := strings.Split(line, "");
		//fmt.Println(temp)
		grid = append(grid, temp)
	}

	var total int = checkHorizontal(grid) + checkVertical(grid) + checkDiagonal(grid)

	fmt.Println(total)
}

func p2(){
	var grid [][]string;
	file, _ := os.Open("2024/4-input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		line := scanner.Text()
		temp := strings.Split(line, "");
		grid = append(grid, temp)
	}

	var total int = findXmasCount(grid)
	fmt.Println(total)
}

func checkHorizontal(g [][]string) int {
	var found int = 0
	for i := 0; i < len(g); i++{
		for j := 0; j < len(g[i])-3; j++{
			letters := strings.Join(g[i][j:j+4], "")
			if letters == "XMAS" || letters == "SAMX"{
				found += 1
			}
		}
	}
	return found
}

func checkVertical(g [][]string) int {
	var found int = 0
	for col := 0; col < len(g[0]); col++{
		for startRow := 0; startRow < len(g)-3; startRow ++{
			var letters []string;
			for i := 0; i < 4 ; i++{
				letters = append(letters, g[startRow+i][col])
			}
			word := strings.Join(letters, "")
			if word == "XMAS" || word == "SAMX"{
				found += 1
			}
		}
	}
	return found
}

func checkDiagonal(g [][]string) int {
	var found int = 0
	//up to the right
	for col := 3; col < len(g[0]); col++{
		for startRow := 0; startRow < len(g)-3; startRow++{
			var letters []string;
			for i := 0; i < 4; i++{
				letters = append(letters, g[startRow+i][col-i])
			}
			word := strings.Join(letters, "")
			if word == "XMAS" || word == "SAMX"{
				found += 1
			}
		}
	}
	//up to the left
	for col := 0; col < len(g[0])-3; col++{
		for startRow := 0; startRow < len(g)-3; startRow++{
			var letters []string;
			for i := 0; i < 4; i++{
				letters = append(letters, g[startRow+i][col+i])
			}
			word := strings.Join(letters, "")
			if word == "XMAS" || word == "SAMX"{
				found += 1
			}
		}
	}
	return found
}

func findXmasCount(g [][]string) int {
	var found int = 0
	for startCol := 0; startCol <= len(g[0])-3; startCol++{
		for startRow := 0; startRow <= len(g)-3; startRow++{
			var goRight []string;
			var goLeft []string;
			for i := 0; i < 3; i++{
				goRight = append(goRight, g[startRow+i][startCol+i])
				goLeft = append(goLeft, g[startRow+2-i][startCol+i])
			}
			wordRight := strings.Join(goRight, "")
			wordLeft := strings.Join(goLeft, "")
			if (wordRight == "MAS" || wordRight == "SAM") && (wordLeft == "MAS" || wordLeft == "SAM"){
				found += 1
			}
		}
	}
	return found
}