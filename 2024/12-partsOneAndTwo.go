package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main(){
	//p1()
	p2()
}

func p1(){
	garden := readInput_12("2024/12-input.txt")
	regions := defineRegions(garden)
	perimetersAndAreas := []int{}
	for _, v := range regions{
		perimetersAndAreas = append(perimetersAndAreas, getPerimetersAndAreas(v, false)...)
	}
	price := 0
	for i := 0; i<len(perimetersAndAreas); i += 2 {
		price += perimetersAndAreas[i]*perimetersAndAreas[i+1]
	}
	fmt.Println(price)
}

func p2(){
	garden := readInput_12("2024/12-input.txt")
	regions := defineRegions(garden)
	perimetersAndAreas := []int{}
	for _, v := range regions{
		perimetersAndAreas = append(perimetersAndAreas, getPerimetersAndAreas(v, true)...)
	}
	price := 0
	for i := 0; i<len(perimetersAndAreas); i += 2 {
		price += perimetersAndAreas[i]*perimetersAndAreas[i+1]
	}
	fmt.Println(price)
}


func readInput_12(fileName string) [][]string {
	var garden [][]string;

	file, _ := os.Open(fileName)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		row := strings.Split(line, "")
		garden = append(garden, row)
	}
	return garden
}

type plotLoc struct {
	x int
	y int
}

func defineRegions(garden [][]string) map[string][]plotLoc {
	regions := make(map[string][]plotLoc)

	for i := 0; i < len(garden); i++{
		for j := 0; j < len(garden[i]); j++{
			plant := garden[i][j]
			regions[plant] = append(regions[plant], plotLoc{x: j, y: i})
		}
	}
	return regions
}

func getPerimetersAndAreas(region []plotLoc, partTwo bool) []int {
	used := []plotLoc{}
	var r [][]plotLoc;
	for len(used) < len(region){
		var cur plotLoc;
		for i := 0; i<len(region); i++{
			if !slices.Contains(used, region[i]){
				cur = region[i]
				used = append(used, cur)
				break
			}
		}
		found := getConnected(cur, region)
		found = dedupe(append(found, cur))
		used = dedupe(append(used, found...))
		r = append(r, found)
	}
	var out []int;
	for i := 0; i< len(r); i++{
		perimeter := 0
		if !partTwo{
			perimeter = getPerimeter(r[i])
		} else {
			perimeter = getSides(r[i])
		}
		area := len(r[i])
		out = append(out, []int{perimeter, area}...)
	}
	return out
}

func areConnected(cur plotLoc, avail []plotLoc) bool {
	options := []plotLoc{plotLoc{x:cur.x+1, y:cur.y},plotLoc{x:cur.x-1, y:cur.y},plotLoc{x:cur.x, y:cur.y+1},plotLoc{x:cur.x, y:cur.y-1}}

	for _, val := range options {
		if slices.Contains(avail, val){
			return true
		}
	}
	return false
}

func getConnected(cur plotLoc, region []plotLoc) []plotLoc {
	changed := true
	items := []plotLoc{cur}
	for changed {
		changed = false
		for i := 0; i < len(items); i++ {
			current := items[i]
			options := []plotLoc{plotLoc{x:current.x+1, y:current.y},plotLoc{x:current.x-1, y:current.y},plotLoc{x:current.x, y:current.y+1},plotLoc{x:current.x, y:current.y-1}}
			for _, op := range options {
				if slices.Contains(region, op) && !slices.Contains(items, op){
					changed = true
					items = append(items, op)
				}
			}
		}
	}
	return items
}

func getPerimeter(pl []plotLoc) int {
	fences := 0

	for i := 0; i<len(pl); i++{
		cur := pl[i]
		options := []plotLoc{plotLoc{x:cur.x+1, y:cur.y},plotLoc{x:cur.x-1, y:cur.y},plotLoc{x:cur.x, y:cur.y+1},plotLoc{x:cur.x, y:cur.y-1}}
		for _, op := range options{
			if !slices.Contains(pl, op){
				fences += 1
			}
		}
	}
	return fences
}

func getSides(pl []plotLoc) int {
	curSides := 0
	for _, val := range pl {
		if !slices.Contains(pl, plotLoc{x: val.x, y:val.y+1}){
			if !slices.Contains(pl, plotLoc{x: val.x-1, y:val.y}){
				curSides += 1
			}
			if !slices.Contains(pl, plotLoc{x: val.x+1, y:val.y}){
				curSides += 1
			}
			if slices.Contains(pl, plotLoc{x: val.x+1, y:val.y+1}) && (slices.Contains(pl, plotLoc{x: val.x+1, y:val.y})){
				curSides += 1
			}
			if slices.Contains(pl, plotLoc{x: val.x-1, y:val.y+1}) && slices.Contains(pl, plotLoc{x: val.x-1, y:val.y}){
				curSides += 1
			}

		} 
		if !slices.Contains(pl, plotLoc{x:val.x, y:val.y-1}){
			if !slices.Contains(pl, plotLoc{x: val.x-1, y:val.y}){
				curSides += 1
			}
			if !slices.Contains(pl, plotLoc{x: val.x+1, y:val.y}){
				curSides += 1
			}
			if slices.Contains(pl, plotLoc{x: val.x+1, y:val.y-1}) && slices.Contains(pl, plotLoc{x: val.x+1, y:val.y}){
				curSides += 1
			}
			if slices.Contains(pl, plotLoc{x: val.x-1, y:val.y-1}) && slices.Contains(pl, plotLoc{x: val.x-1, y:val.y}){
				curSides += 1
			}
		}
	}
	return curSides
}

func dedupe(input []plotLoc) []plotLoc {
	out := []plotLoc{}

	for _, v := range input{
		if !slices.Contains(out, v){
			out = append(out, v)
		}
	}
	return out
}