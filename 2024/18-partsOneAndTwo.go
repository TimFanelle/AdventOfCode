package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
	"container/heap"
	"math"
)

func main(){
	//p1()
	p2()
}

func p1(){
	falling := readInput_18("2024/18-input.txt")
	corruptedMap := genCorrupted(70, falling, 1024)
	start := Point{X: 0, Y:0}
	goal := Point{X:70, Y:70}
	path, found := AStar(start, goal, corruptedMap)
	if found {
		fmt.Println(len(path)-1)
	} else {
		fmt.Println("nothing found")
	}
}

func p2(){
	falling := readInput_18("2024/18-input.txt")
	start := Point{X: 0, Y:0}
	goal := Point{X:70, Y:70}

	i := 1
	for true {
		corruptedMap := genCorrupted(70, falling, i)
	
		_, found := AStar(start, goal, corruptedMap)
		if found {
			i += 1
		} else {
			fmt.Println(falling[i-1])
			break
		}
	}
	
}

func readInput_18(fileName string) [][]int {
	out := [][]int{}
	file, _ := os.Open(fileName)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		out = append(out, []int{convString2Int(line[0]), convString2Int(line[1])})
	}
	return out
}

func convString2Int(s string) int {
	val, _ := strconv.Atoi(s)
	return val
}

func genCorrupted(size int, fallingBytes [][]int, numBytes int) [][]string {
	coreMap := [][]string{}

	for i := 0; i <= size; i++{
		mini := []string{}
		for j := 0; j <= size; j++{
			mini = append(mini, ".")
		}
		coreMap = append(coreMap, mini)
	}

	bytesToAdd := fallingBytes[:numBytes]
	for _, v := range bytesToAdd {
		coreMap[v[1]][v[0]] = "#"
	}
	return coreMap
}

type Point struct { 
	X, Y int 
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