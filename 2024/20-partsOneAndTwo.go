package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
)

func main(){
	//p1()
	p2()
}

var aStarMemo = make(map[PointPair][]Point)

func p1() {
	grid := readInput_20("2024/20-input.txt")
	start, end := getStartEnd(grid)
	path, _ := AStar(start, end, grid)
	pthCopy := make([]Point, len(path))
	copy(pthCopy, path)
	slices.Reverse(pthCopy)
	startingShortestLength := len(path)-1

	var distMemo = make(map[PointPair]int)
	for i, p := range pthCopy{
		distMemo[PointPair{p1: p, p2: end}] = i+1
	}

	cheatCounts := make(map[int]int)

	for i, p := range path{
		viableCheats := getPointsFromManhattan(p, grid, 2)
		for _, c := range viableCheats{
			if grid[c.Y][c.X] != "#"{
				lenTo := i + distMemo[PointPair{p1: c, p2: end}]+1
				if startingShortestLength-lenTo > 0 {
					cheatCounts[startingShortestLength-lenTo] += 1
				}
			}
		}
	}
	//fmt.Println(cheatCounts)
	count := 0
	for k, v := range cheatCounts {
		if k >= 100 {
			count += v
		}
	}
	fmt.Println(count)
}

func p2(){
	grid := readInput_20("2024/20-input.txt")
	start, end := getStartEnd(grid)
	path, _ := AStar(start, end, grid)
	pthCopy := make([]Point, len(path))
	copy(pthCopy, path)
	slices.Reverse(pthCopy)
	startingShortestLength := len(path)-1

	var distMemo = make(map[PointPair]int)
	for i, p := range pthCopy{
		distMemo[PointPair{p1: p, p2: end}] = i+1
	}

	cheatCounts := make(map[int]int)

	for i, p := range path{
		viableCheats := getPointsFromManhattan(p, grid, 20)
		for _, c := range viableCheats{
			if grid[c.Y][c.X] == "."{
				lenTo := i + distMemo[PointPair{p1: c, p2: end}]+1
				if startingShortestLength-lenTo > 0 {
					cheatCounts[startingShortestLength-lenTo] += 1
				}
			}
		}
	}
	fmt.Println(cheatCounts)
	count := 0
	for k, v := range cheatCounts {
		if k >= 100 {
			count += v
		}
	}
	fmt.Println(count)
}

func readInput_20(fileName string) [][]string {
	grid := [][]string{}
	file, _ := os.Open(fileName)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		grid = append(grid, line)
	}
	return grid
}

func getStartEnd(g [][]string) (Point, Point){
	start, end := Point{X: 0, Y:0}, Point{X: 0, Y:0}
	for i := 0; i < len(g); i++{
		for j := 0; j < len(g[i]); j++ {
			if g[i][j] == "S"{
				start = Point{X:j, Y:i}
			} else if g[i][j] == "E" {
				end = Point{X:j, Y:i}
			}
		}
	}
	return start, end
}

type Point struct { 
	X, Y int 
}
type PointPair struct {
	p1 Point
	p2 Point
} 
// Node represents a point in the grid along with cost and priority 
type Node struct { 
	Point Point 
	Cost int 
	Priority int 
	Index int 
} 
// PriorityQueue implements heap.Interface and holds Nodes 
type PriorityQueue []*Node 
func (pq PriorityQueue) Len() int { 
	return len(pq) 
} 

func (pq PriorityQueue) Less(i, j int) bool { 
	return pq[i].Priority < pq[j].Priority 
} 

func (pq PriorityQueue) Swap(i, j int) { 
	pq[i], pq[j] = pq[j], pq[i] 
	pq[i].Index = i 
	pq[j].Index = j 
} 

func (pq *PriorityQueue) Push(x interface{}) { 
	n := len(*pq) 
	node := x.(*Node) 
	node.Index = n 
	*pq = append(*pq, node) 
} 

func (pq *PriorityQueue) Pop() interface{} { 
	old := *pq 
	n := len(old) 
	node := old[n-1] 
	old[n-1] = nil // avoid memory leak 
	node.Index = -1 // for safety 
	*pq = old[0 : n-1] 
	return node 
} 

func (pq *PriorityQueue) update(node *Node, priority int) { 
	node.Priority = priority 
	heap.Fix(pq, node.Index) 
} 
// ManhattanDistance calculates the Manhattan distance between two points 
func ManhattanDistance(a, b Point) int { 
	return int(math.Abs(float64(a.X-b.X)) + math.Abs(float64(a.Y-b.Y))) 
} 
// AStar performs the A* algorithm to find the shortest path 
func AStar(start, goal Point, grid [][]string) ([]Point, bool) {
	pr, ok := aStarMemo[PointPair{p1:start, p2:goal}]
	if ok {
		return pr, true
	} 
	rows, cols := len(grid), len(grid[0]) 
	openSet := &PriorityQueue{} 
	heap.Init(openSet) 
	heap.Push(openSet, &Node{Point: start, Cost: 0, Priority: 0}) 
	cameFrom := make(map[Point]*Point) 
	gScore := make(map[Point]int) 
	gScore[start] = 0 
	fScore := make(map[Point]int) 
	fScore[start] = ManhattanDistance(start, goal) 
	for openSet.Len() > 0 { 
		currentNode := heap.Pop(openSet).(*Node) 
		current := currentNode.Point 
		if current == goal { 
			return reconstructPath(cameFrom, current), true 
		} 
		for _, neighbor := range getNeighbors(current, rows, cols) { 
			if grid[neighbor.Y][neighbor.X] == "#" { // 1 indicates obstacle 
				continue 
			} 
			tentativeGScore := gScore[current] + 1 
			if score, exists := gScore[neighbor]; !exists || tentativeGScore < score { 
				cameFrom[neighbor] = &current 
				gScore[neighbor] = tentativeGScore 
				fScore[neighbor] = tentativeGScore + ManhattanDistance(neighbor, goal) 
				heap.Push(openSet, &Node{Point: neighbor, Cost: tentativeGScore, Priority: fScore[neighbor]}) 
			} 
		} 
	} 
	return nil, false 
} 
// getNeighbors returns the neighbors of a point 
func getNeighbors(point Point, rows, cols int) []Point { 
	dirs := []Point{{X: -1, Y: 0}, {X: 1, Y: 0}, {X: 0, Y: -1}, {X: 0, Y: 1}} 
	neighbors := []Point{} 
	for _, dir := range dirs { 
		neighbor := Point{X: point.X + dir.X, Y: point.Y + dir.Y} 
		if neighbor.X >= 0 && neighbor.X < cols && neighbor.Y >= 0 && neighbor.Y < rows { 
			neighbors = append(neighbors, neighbor) 
		} 
	} 
	return neighbors 
} 
// reconstructPath reconstructs the path from start to goal 
func reconstructPath(cameFrom map[Point]*Point, current Point) []Point { 
	path := []Point{current} 
	for cameFrom[current] != nil { 
		current = *cameFrom[current] 
		path = append([]Point{current}, path...) 
	} 
	return path
}

func getPointsFromManhattan(p Point, g [][]string, cD int) []Point {
	viable := []Point{}
	for i := 0; i < len(g); i++{
		for j := 0; j < len(g[i]); j++{
			temp := Point{X:j, Y:i}
			if temp != p && ManhattanDistance(p, temp) <= cD{
				viable = append(viable, temp)
			}
		}
	}
	return viable
}