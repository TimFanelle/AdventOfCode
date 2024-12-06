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

func p1() {
	mapImg := readInput("2024/6-input.txt")
	gl := determineStartingLoc(mapImg)
	nStep := moveGuardAndCount(mapImg, gl)
	fmt.Println(nStep)
}

func p2() {
	mapImg := readInput("2024/6-input.txt")
	gl := determineStartingLoc(mapImg)
	nObs := determineObstacles(mapImg, gl)
	fmt.Println(nObs)
}

func readInput(input string) [][]string {
	var grid [][]string;
	file, _ := os.Open(input)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan(){
		line := scanner.Text()
		temp := strings.Split(line, "")
		grid = append(grid, temp)
	}

	return grid
}

type guardLoc struct {
	row int
	col int
	direction int
}
type loc struct {
	x int
	y int
	facing_dir int
}

func determineStartingLoc(input [][]string) guardLoc {
	for row := 0; row < len(input); row++ {
		for col := 0; col < len(input[row]); col++ {
			if input[row][col] == "^"{
				return guardLoc{row:row, col:col, direction:1}
			}
		}
	}
	return guardLoc{}
}

func moveGuardAndCount(input [][]string, curLoc guardLoc) int {
	validLoc := curLoc.col < len(input[0]) && curLoc.col >= 0 && curLoc.row >= 0 && curLoc.row < len(input)
	var allLocs map[loc]bool = make(map[loc]bool);
	allLocs[loc{x: curLoc.col, y: curLoc.row}] = true
	for validLoc {
		switch curLoc.direction {
			case 4:
				curLoc.col -= 1
				if curLoc.col -1 >= 0 && input[curLoc.row][curLoc.col-1] == "#"{
					curLoc.direction = 1
				}
			case 2:
				curLoc.col += 1
				if curLoc.col +1 < len(input[curLoc.row]) && input[curLoc.row][curLoc.col+1] == "#"{
					curLoc.direction = 3
				}
			case 1:
				curLoc.row -= 1
				if curLoc.row-1 >=0 && input[curLoc.row-1][curLoc.col] == "#" {
					curLoc.direction = 2
				}
			
			case 3:
				curLoc.row += 1
				if curLoc.row + 1 < len(input) && input[curLoc.row+1][curLoc.col] == "#"{
					curLoc.direction = 4
				}
			
		}
		validLoc = curLoc.col < len(input[0]) && curLoc.col >= 0 && curLoc.row >= 0 && curLoc.row < len(input)
		if validLoc{
			_, ok := allLocs[loc{x: curLoc.col, y: curLoc.row, facing_dir: 1}]
			if !ok {
				allLocs[loc{x: curLoc.col, y: curLoc.row}] = true
			}
		}
	}
	numMoves := len(allLocs)
	return numMoves
}

func makeMapCopy(input [][]string, r int, c int) [][]string {
	duplicate := make([][]string, len(input))
	for i := range input {
		duplicate[i] = make([]string, len(input[i]))
		copy(duplicate[i], input[i])
	}
	duplicate[r][c] = "#"
	return duplicate

}

func determineObstacles(input [][]string, curLoc guardLoc) int {
	var numObs int = 0
	for row := 0; row < len(input); row++{
		for col := 0; col < len(input[row]); col++{
			if input[row][col] == "." {
				temp := makeMapCopy(input, row, col)
				if determineIfLoop(temp, curLoc){
					numObs++
				}
			}
		}
	}
	return numObs
}

func move(curLoc guardLoc) guardLoc {
	switch curLoc.direction {
		case 4:
			curLoc.col -= 1
		case 2:
			curLoc.col += 1
		case 1:
			curLoc.row -= 1
		case 3:
			curLoc.row += 1
	}
	return curLoc
}

func nextIsUnobstructed(input [][]string, curLoc guardLoc) bool {
	switch curLoc.direction {
		case 4:
			if curLoc.col -1 >= 0 && input[curLoc.row][curLoc.col-1] != "#"{
				return true
			} else if curLoc.col - 1 < 0 {
				return true
			}
		case 2:
			if curLoc.col +1 < len(input[curLoc.row]) && input[curLoc.row][curLoc.col+1] != "#"{
				return true
			}else if curLoc.col + 1 >= len(input[curLoc.row]) {
				return true
			}
		case 1:
			if curLoc.row-1 >=0 && input[curLoc.row-1][curLoc.col] != "#" {
				return true
			} else if curLoc.row - 1 < 0 {
				return true
			}
		
		case 3:
			if curLoc.row + 1 < len(input) && input[curLoc.row+1][curLoc.col] != "#"{
				return true
			} else if curLoc.row + 1 >= len(input) {
				return true
			}
		default:
			return false
	}
	return false
}

func turn(curLoc guardLoc) int {
	switch curLoc.direction {
		case 4:
			return 1
		case 2:
			return 3
		case 1:
			return 2
		case 3:
			return 4
	}
	return 0
}

func determineIfLoop(input [][]string, curLoc guardLoc) bool {
	allLocs := make(map[loc]bool);
	
	validLoc := curLoc.col < len(input[0]) && curLoc.col >= 0 && curLoc.row >= 0 && curLoc.row < len(input)
	allLocs[loc{x: curLoc.col, y: curLoc.row, facing_dir: curLoc.direction}] = true
	
	for validLoc {
		nextValid := nextIsUnobstructed(input, curLoc)
		if nextValid {
			curLoc = move(curLoc)
		} else {
			curLoc.direction = turn(curLoc)
		}
		_, ok := allLocs[loc{x: curLoc.col, y: curLoc.row, facing_dir: curLoc.direction}]
		if ok {
			return true
		} else {
			allLocs[loc{x: curLoc.col, y: curLoc.row, facing_dir: curLoc.direction}] = true
		}
		validLoc = curLoc.col < len(input[0]) && curLoc.col >= 0 && curLoc.row >= 0 && curLoc.row < len(input)
	}
	return false
}