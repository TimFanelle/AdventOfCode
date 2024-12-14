package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	//p1()
	p2()
}

func p1(){
	robots := readInput_14("2024/14-input.txt")
	maxVals := []int{101, 103}

	var updatedRobots []robot
	quadrantCounts := []int{0, 0, 0, 0, 0}
	for i, r := range robots {
		updatedRobots = append(updatedRobots, simulateMovement(r, 100, maxVals[0], maxVals[1]))
		quadrantCounts[findQuadrant(updatedRobots[i], maxVals)] += 1
	}

	count := 1
	for i := 0; i<4; i++{
		count *= quadrantCounts[i]
	}
	fmt.Println(count)
}

func p2(){
	robots := readInput_14("2024/14-input.txt")
	maxVals := []int{101, 103}
	minSafetyAndTimestep := []int{215987200, 100}

	for i := 1; i < 10000; i++{
		var updatedRobots []robot
		quadrantCounts := []int{0, 0, 0, 0, 0}
		for j, r := range robots {
			updatedRobots = append(updatedRobots, simulateMovement(r, i, maxVals[0], maxVals[1]))
			quadrantCounts[findQuadrant(updatedRobots[j], maxVals)] += 1
		}
		count := 1
		for j := 0; j<4; j++{
			count *= quadrantCounts[j]
		}
		if count < minSafetyAndTimestep[0]{ // christmas tree will be when there is the highest density of robots in a single area and thus the lowest safety score
			minSafetyAndTimestep[0] = count
			minSafetyAndTimestep[1] = i
		}
	}
	
	fmt.Println(minSafetyAndTimestep)
}

type robot struct {
	x int
	y int
	dx int
	dy int
}

func readInput_14(fileName string) []robot {
	var robots []robot;

	file, _ := os.Open(fileName)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan(){
		line := strings.Split(scanner.Text(), " ")
		valPos := strings.Split(line[0][2:], ",")
		valDelt := strings.Split(line[1][2:], ",")

		nRobot := robot{x: convString2Int(valPos[0]), y:convString2Int(valPos[1]), dx: convString2Int(valDelt[0]), dy: convString2Int(valDelt[1])}
		robots = append(robots, nRobot)
	}

	return robots
}

func simulateMovement(r robot, t int, mx int, my int) robot {
	newX := (((r.x+(r.dx*t)) % mx) + mx) % mx
	newY := (((r.y+(r.dy*t)) % my) + my) % my
	nRobot := robot{x: newX, y: newY, dx:0, dy:0}
	return nRobot
}

func convString2Int(s string) int {
	val, _ := strconv.Atoi(s)
	return val
}

func findQuadrant(r robot, maxVals []int) int {
	if r.x == maxVals[0]/2 || r.y == maxVals[1]/2{
		return 4
	}
	if r.x > maxVals[0]/2{
		if r.y > maxVals[1]/2 {
			return 3
		} else {
			return 1
		}
	} else {
		if r.y > maxVals[1]/2 {
			return 2
		} else {
			return 0
		}
	}
}