package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	//p1()
	p2()
}

func p1() {
	mappy, moves := readInput_15("2024/15-input.txt")
	
	curLoc := findStartingLoc(mappy)
	mappy[curLoc[1]][curLoc[0]] = "."
	
	for i := 0; i<len(moves); i++ {
		curLoc, mappy = attemptMove(curLoc, mappy, moves[i:i+1])
	}
	
	for _, r := range mappy {
		fmt.Println(r)
	}
	count := getGPSsum(mappy, false)
	fmt.Println(count)
}

func p2(){
	mappy, moves := readInput_15("2024/15-input.txt")
	mappy = updateMapForP2(mappy)
	
	for _, r := range mappy {
		fmt.Println(r)
	}
	fmt.Println(moves)
	
	curLoc := findStartingLoc(mappy)
	mappy[curLoc[1]][curLoc[0]] = "."

	for i := 0; i<len(moves); i++ {
		curLoc, mappy = attemptMoveP2(curLoc, mappy, moves[i:i+1])
	}
	
	mappy[curLoc[1]][curLoc[0]] = "@"
	for _, r := range mappy {
		fmt.Println(r)
	}
	count := getGPSsum(mappy, true)
	fmt.Println(count)
}

func readInput_15(fileName string) ([][]string, string) {
	var mappy [][]string;
	moves := ""

	file, _ := os.Open(fileName)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	movements := false
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" || line == "\n" || line == " " {
			movements = true
		} else {
			if movements {
				moves = line
			} else {
				c := strings.Split(line, "")
				mappy = append(mappy, c)
			}
		}
	}
	return mappy, moves
}

func findStartingLoc(m [][]string) [2]int {
	for i := 0; i<len(m); i++{
		for j := 0; j<len(m[i]); j++{
			if m[i][j] == "@"{
				return [2]int{j, i}
			}
		}
	}
	return [2]int{0,0}
}

func attemptMove(curLoc [2]int, m [][]string, move string) ([2]int, [][]string){
	nextLoc := [2]int{0, 0}
	switch move {
	case "<":
		nextLoc = [2]int{curLoc[0]-1,curLoc[1]}
	case ">":
		nextLoc = [2]int{curLoc[0]+1,curLoc[1]}
	case "^":
		nextLoc = [2]int{curLoc[0], curLoc[1]-1}
	case "v":
		nextLoc = [2]int{curLoc[0], curLoc[1]+1}
	default:
		nextLoc = curLoc
	}

	if nextLoc[0] < len(m[0]) && nextLoc[1] < len(m){
		nextSpot := m[nextLoc[1]][nextLoc[0]]
		if nextSpot == "."{
			return nextLoc, m
		} else if nextSpot == "#" {
			return curLoc, m
		} else if nextSpot == "O" {
			endLoc, m := moveBlocks(curLoc, m, move)
			return endLoc, m
		}
	}
	return curLoc, m
}

func attemptMoveP2(curLoc [2]int, m [][]string, move string) ([2]int, [][]string){
	nextLoc := [2]int{0, 0}
	switch move {
	case "<":
		nextLoc = [2]int{curLoc[0]-1,curLoc[1]}
	case ">":
		nextLoc = [2]int{curLoc[0]+1,curLoc[1]}
	case "^":
		nextLoc = [2]int{curLoc[0], curLoc[1]-1}
	case "v":
		nextLoc = [2]int{curLoc[0], curLoc[1]+1}
	default:
		nextLoc = curLoc
	}
	if nextLoc[0] < len(m[0]) && nextLoc[1] < len(m){
		nextSpot := m[nextLoc[1]][nextLoc[0]]
		if nextSpot == "."{
			return nextLoc, m
		} else if nextSpot == "#" {
			return curLoc, m
		} else if nextSpot == "[" || nextSpot == "]" {
			endLoc, m := moveBlocksP2(curLoc, m, move)
			return endLoc, m
		}
	}
	return curLoc, m
}

func moveBlocks(curLoc [2]int, m [][]string, move string) ([2]int, [][]string) {
	mC := make([][]string, len(m))
    for i := range m {
        mC[i] = make([]string, len(m[i]))
        copy(mC[i], m[i])
    }
	
	switch move {
		case "<":
			finalBlock := [2]int{-1, curLoc[1]}
			for i := curLoc[0]-1; i >= 0; i--{
				if mC[curLoc[1]][i] == "." {
					finalBlock[0] = i
					break
				} else if mC[curLoc[1]][i] == "#"{
					break
				}
			}
			if finalBlock[0] != -1 {
				if m[finalBlock[1]][finalBlock[0]] == "."{
					for i := finalBlock[0]; i < curLoc[0]-1; i++{
						mC[curLoc[1]][i] = "O"
						mC[curLoc[1]][i+1] = "."
					}
					curLoc[0] = curLoc[0]-1
				}
			}
			
		case ">":
			finalBlock := [2]int{-1, curLoc[1]}
			for i := curLoc[0]+1; i < len(m); i++{
				if mC[curLoc[1]][i] == "." {
					finalBlock[0] = i
					break
				} else if mC[curLoc[1]][i] == "#"{
					break
				}
			}
			if finalBlock[0] != -1{
				if m[finalBlock[1]][finalBlock[0]] == "."{
					for i := finalBlock[0]; i > curLoc[0]+1; i--{
						mC[curLoc[1]][i] = "O"
						mC[curLoc[1]][i-1] = "."
					}
					curLoc[0] = curLoc[0]+1
				}
			}
			
		case "^":
			finalBlock := [2]int{curLoc[0], -1}
			for i := curLoc[1]-1; i >= 0; i--{
				if mC[i][curLoc[0]] == "." {
					finalBlock[1] = i
					break
				} else if mC[i][curLoc[0]] == "#"{
					break
				}
			}
			if finalBlock[1] != -1 {
				if mC[finalBlock[1]][finalBlock[0]] == "."{
					for i := finalBlock[1]; i < curLoc[1]-1; i++{
						mC[i][curLoc[0]] = "O"
						mC[i+1][curLoc[0]] = "."
					}
					curLoc[1] = curLoc[1]-1
				}
			}
			
		case "v":
			finalBlock := [2]int{curLoc[0], -1}
			for i := curLoc[1]+1; i < len(m[0]); i++{
				if mC[i][curLoc[0]] == "."{
					finalBlock[1] = i
					break
				} else if mC[i][curLoc[0]] == "#"{
					break
				}
			}
			if finalBlock[1] != -1 {
				if mC[finalBlock[1]][finalBlock[0]] == "."{
					for i := finalBlock[1]; i > curLoc[1]+1; i--{
						mC[i][curLoc[0]] = "O"
						mC[i-1][curLoc[0]] = "."
					}
					curLoc[1] = curLoc[1]+1
				}
			}
			
		default:
			return curLoc, m
	}
	return curLoc, mC
}
func moveBlocksP2(curLoc [2]int, m [][]string, move string) ([2]int, [][]string) {
	mC := make([][]string, len(m))
    for i := range m {
        mC[i] = make([]string, len(m[i]))
        copy(mC[i], m[i])
    }
	
	switch move {
		case "<":
			finalBlock := [2]int{-1, curLoc[1]}
			for i := curLoc[0]-1; i >= 0; i--{
				if mC[curLoc[1]][i] == "." {
					finalBlock[0] = i
					break
				} else if mC[curLoc[1]][i] == "#"{
					break
				}
			}
			if finalBlock[0] != -1 {
				if m[finalBlock[1]][finalBlock[0]] == "."{
					for i := finalBlock[0]; i < curLoc[0]-1; i++{
						mC[curLoc[1]][i] = mC[curLoc[1]][i+1]
						mC[curLoc[1]][i+1] = mC[curLoc[1]][i+2]
					}
					curLoc[0] = curLoc[0]-1
				}
			}
			
		case ">":
			finalBlock := [2]int{-1, curLoc[1]}
			for i := curLoc[0]+1; i < len(m[0]); i++{
				if mC[curLoc[1]][i] == "." {
					finalBlock[0] = i
					break
				} else if mC[curLoc[1]][i] == "#"{
					break
				}
			}
			if finalBlock[0] != -1{
				if m[finalBlock[1]][finalBlock[0]] == "."{
					for i := finalBlock[0]; i > curLoc[0]+1; i--{
						mC[curLoc[1]][i] = mC[curLoc[1]][i-1]
						mC[curLoc[1]][i-1] = mC[curLoc[1]][i-2]
					}
					curLoc[0] = curLoc[0]+1
				}
			}
			
		case "^":
			connected := findAllConnected(curLoc, mC, move)
			for i := 0; i<curLoc[1]; i++{
				for _, val := range connected {
					if val[1] == i {
						mC[val[1]-1][val[0]] = mC[val[1]][val[0]]
						mC[val[1]][val[0]] = "." 
					}
				}
			}
			if len(connected) > 0 {
				curLoc[1] = curLoc[1]-1
			}
			
		case "v":
			connected := findAllConnected(curLoc, mC, move)
			for i := len(mC)-1; i>curLoc[1]; i--{
				for _, val := range connected {
					if val[1] == i {
						mC[val[1]+1][val[0]] = mC[val[1]][val[0]]
						mC[val[1]][val[0]] = "." 
					}
				}
			}
			if len(connected) > 0 {
				curLoc[1] = curLoc[1] +1
			}
			
		default:
			return curLoc, m
	}
	return curLoc, mC
}

func getGPSsum(m [][]string, part2 bool) int {
	count := 0
	for i := 0; i< len(m); i++{
		for j := 0; j < len(m[i]); j++ {
			if part2 {
				if m[i][j] == "["{
					count += (100*i) + j
				}
			} else {
				if m[i][j] == "O" {
					count += (100*i) + j
				}
			}
			
		}
	}
	return count
}

func updateMapForP2(m [][]string) [][]string {
	out := [][]string{}

	for _, r := range m {
		row := []string{}
		for _, item := range r {
			switch item {
			case "#":
				row = append(row, []string{"#", "#"}...)
			case "O":
				row = append(row, []string{"[", "]"}...)
			case ".":
				row = append(row, []string{".", "."}...)
			case "@":
				row = append(row, []string{"@", "."}...)
			default:
				// do nothing
			}
		}
		out = append(out, row)
	}
	return out
}

func findAllConnected(curLoc [2]int, m [][]string, dir string) [][2]int {
	coords := [][2]int{}

	switch dir {
	case "^":
		coords = append(coords, [2]int{curLoc[0], curLoc[1]-1})
		if m[curLoc[1]-1][curLoc[0]] == "["{
			coords = append(coords, [2]int{curLoc[0]+1, curLoc[1]-1})
		} else {
			coords = append(coords, [2]int{curLoc[0]-1, curLoc[1]-1})
		}
		for i := curLoc[1]-1; i>=0; i--{
			add_to_coords := [][2]int{}
			for j := 0; j < len(m[0]); j++{
				for ind := range coords{
					if j == coords[ind][0] && coords[ind][1] == i+1 && (m[i][j] == "[" || m[i][j] == "]" || m[i][j] == "#"){
						add_to_coords = append(add_to_coords, [2]int{j, i})
					}
				}
			}
			others_to_add := [][2]int{}
			for j := 0; j<len(add_to_coords); j++ {
				if m[add_to_coords[j][1]][add_to_coords[j][0]] == "["{
					others_to_add = append(others_to_add, [2]int{add_to_coords[j][0]+1, add_to_coords[j][1]})
				} else if m[add_to_coords[j][1]][add_to_coords[j][0]] == "]"{
					others_to_add = append(others_to_add, [2]int{add_to_coords[j][0]-1, add_to_coords[j][1]})
				}
			}
			add_to_coords = removeDuplicates(append(add_to_coords, others_to_add...))
			coords = append(coords, add_to_coords...)
		}
	case "v":
		coords = append(coords, [2]int{curLoc[0], curLoc[1]+1})
		if m[curLoc[1]+1][curLoc[0]] == "["{
			coords = append(coords, [2]int{curLoc[0]+1, curLoc[1]+1})
		} else {
			coords = append(coords, [2]int{curLoc[0]-1, curLoc[1]+1})
		}
		for i := curLoc[1]+1; i<len(m); i++{
			add_to_coords := [][2]int{}
			for j := 0; j < len(m[0]); j++{
				for ind := range coords{
					if j == coords[ind][0] && coords[ind][1] == i-1 &&(m[i][j] == "[" || m[i][j] == "]" || m[i][j] == "#"){
						add_to_coords = append(add_to_coords, [2]int{j, i})
					}
				}
			}
			others_to_add := [][2]int{}
			for j := 0; j<len(add_to_coords); j++ {
				if m[add_to_coords[j][1]][add_to_coords[j][0]] == "["{
					others_to_add = append(others_to_add, [2]int{add_to_coords[j][0]+1, add_to_coords[j][1]})
				} else if m[add_to_coords[j][1]][add_to_coords[j][0]] == "]"{
					others_to_add = append(others_to_add, [2]int{add_to_coords[j][0]-1, add_to_coords[j][1]})
				}
			}
			add_to_coords = removeDuplicates(append(add_to_coords, others_to_add...))
			coords = append(coords, add_to_coords...)
		}
	default:
		// dont do anything
	}
	for _, c := range coords {
		if m[c[1]][c[0]] == "#" {
			coords = [][2]int{}
			break
		}
	}
	return coords
}

func removeDuplicates(inp [][2]int) [][2]int{
	out := [][2]int{}
	for _, v := range inp {
		if !slices.Contains(out, v) {
			out = append(out, v)
		}
	}
	return out
}