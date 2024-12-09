package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
)

func main(){
	//p1()
	p2()
}

func p1(){
	disk_map := readInput_9("2024/9-input.txt")
	layout := generate_layout(disk_map)
	layout = rearrange_layout(layout)
	checkSum := calculateChecksum(layout, false)
	fmt.Println(checkSum)
}

func p2(){
	disk_map := readInput_9("2024/9-input.txt")
	layout := generate_layout_p2(disk_map)
	//fmt.Println(layout)
	layout_0 := rearrange_layout_p2(layout)
	//fmt.Println(layout_0)
	checkSum := calculateChecksum(layout_0, true)
	fmt.Println(checkSum)
}

func readInput_9(fileName string) string{
	var s_out string;
	file, _ := os.Open(fileName)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan(){
		line := scanner.Text()
		s_out += line
	}
	return s_out
}

func str2int(in string) int {
	val, _ := strconv.Atoi(in)
	return val
}

func generate_layout(disk_map string) []string {
	id_num := 0
	var layout []string;

	for i := 0; i < len(disk_map); i++{
		var temp []string; 
		for j := 0; j < str2int(disk_map[i:i+1]); j++ {
			if i % 2 == 0 {
				temp = append(temp, strconv.Itoa(id_num))
			} else {
				temp = append(temp, ".")
			}
		}
		if i % 2 == 0 {
			id_num += 1
		}
		layout = append(layout, temp[:]...)
	}
	return layout
}

func generate_layout_p2(disk_map string) [][]string {
	id_num := 0
	var layout [][]string;

	for i := 0; i < len(disk_map); i++{
		var temp []string; 
		for j := 0; j < str2int(disk_map[i:i+1]); j++ {
			if i % 2 == 0 {
				temp = append(temp, strconv.Itoa(id_num))
			} else {
				temp = append(temp, ".")
			}
		}
		if i % 2 == 0 {
			id_num += 1
		}
		layout = append(layout, temp)
	}
	return layout
}

func rearrange_layout(layout []string) []string {
	var s_out []string
	lastValIndex := len(layout)-1
	for layout[lastValIndex] == "."{
		lastValIndex -= 1
	}

	for i := 0; i < len(layout); i++ {
		nextVal := layout[i]
		if i >= lastValIndex+1{
			break
		} else if nextVal != "."{
			s_out = append(s_out, layout[i])
		} else {
			s_out = append(s_out, layout[lastValIndex])
			lastValIndex -= 1
			for layout[lastValIndex] == "."{
				lastValIndex -= 1
			}
		}
	}
	return s_out
}

func rearrange_layout_p2(layout [][]string) []string {
	var s_out []string;
	var lastItemsChecked []string;

	indexLastItem := len(layout) -1
	lastItem := layout[indexLastItem]
	for len(lastItem) == 0 || lastItem[0] == "."{
		indexLastItem -= 1
		lastItem = layout[indexLastItem]
	}
	
	var allOpenIndices []int
	for i := 0; i < indexLastItem; i++ {
		if len(layout[i]) > 0 && layout[i][0] == "."{
			allOpenIndices = append(allOpenIndices, i)
		}
	}

	for len(allOpenIndices) > 0 && allOpenIndices[0] <= indexLastItem {
		valLast := layout[indexLastItem][0]
		for _, ind := range allOpenIndices {
			if len(layout[indexLastItem]) <= len(layout[ind]){
				len_valLast := len(layout[indexLastItem])
				t := append([][]string{layout[indexLastItem]}, layout[ind][len(layout[indexLastItem]):])
				u := append(t, layout[ind+1:]...)
				layout = append(layout[:ind], u...)
				lastItemsChecked = append(lastItemsChecked, valLast)
				temp := []string{}
				for j := 0; j < len_valLast; j++ {
					temp = append(temp, ".")
				}
				layout[indexLastItem+1] = temp
				break
			}
		}
		if !slices.Contains(lastItemsChecked, valLast){
			lastItemsChecked = append(lastItemsChecked, valLast)
		}

		indexLastItem = len(layout)-1
		lastItem := layout[indexLastItem]
		for len(lastItem) == 0 || lastItem[0] == "." || slices.Contains(lastItemsChecked, lastItem[0]){
			indexLastItem -= 1
			if indexLastItem > 0 {
				lastItem = layout[indexLastItem]
			} else {
				indexLastItem = -1
				lastItem = layout[0]
				break
			}
		}

		allOpenIndices = []int{}
		for i := 0; i < indexLastItem; i++ {
			if len(layout[i]) > 0 && layout[i][0] == "."{
				allOpenIndices = append(allOpenIndices, i)
			}
		}

	}
	for i := 0; i < len(layout); i++ {
		s_out = append(s_out, layout[i]...)
	}
	
	return s_out
}


func calculateChecksum(digitsIn []string, part2 bool) int {
	var total int = 0;
	for i := 0; i < len(digitsIn); i++ {
		if part2 && digitsIn[i] == "."{
			continue
		} else {
			total += i*str2int(digitsIn[i])
		}
		
	}
	return total
}