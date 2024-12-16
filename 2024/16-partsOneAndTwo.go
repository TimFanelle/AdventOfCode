package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"container/heap"
)

func main(){
	//p1()
	p2()
}

func p1(){
	mappy := readInput_16("2024/16-input.txt")
	cL := findStart(mappy)
	mappyW := weightMap(mappy)

	shortestPathLength, _ := dijkstra(mappyW, cL, mappy)
	fmt.Println(shortestPathLength)
}

func p2(){
	mappy := readInput_16("2024/16-input.txt")
	cL := findStart(mappy)
	mappyW := weightMap(mappy)

	shortestPathLength, pp := dijkstra(mappyW, cL, mappy)
	fmt.Println(shortestPathLength)
	uniqueP := make(map[point]bool)
	for _, v := range pp {
		for _, k := range v{
			uniqueP[point{x: k.x, y: k.y, dir:0}] = true
		}
	}
	fmt.Println(len(uniqueP))
}

func readInput_16(fileName string) [][]string {
	mappy := [][]string{}
	file, _ := os.Open(fileName)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		mappy = append(mappy, line)
	}
	return mappy
}

func findStart(m [][]string) point {
	for i := 0; i<len(m); i++{
		for j := 0; j < len(m[i]); j++{
			if m[i][j] == "S" {
				return point{x: j, y: i, dir: 1}
			}
		}
	}
	return point{x: 0, y: 0, dir: 0}
}

type point struct {
	x int
	y int
	dir int
}

type edge struct {
	p point
	score int
}

func weightMap(m [][]string) map[point][]edge {
	weighted := make(map[point][]edge)
	
	for i := 0; i < len(m); i++{
		for j := 0; j < len(m[i]); j++{
			if m[i][j] == "." || m[i][j] == "S" || m[i][j] == "E"{
				for k := 0; k < 4; k++{
					locEdges := determineEdges(m, point{x:j, y: i, dir: k})
					weighted[point{x:j, y: i, dir: k}] = locEdges
				}
			} 
		}
	}
	return weighted
}

func determineEdges(m [][]string, p point) []edge {
	out := []edge{}

	valid, np := nextPointValid(m, p)
	if  valid{
		out = append(out, edge{p: np, score: 1})
	}

	// turn left
	ndir := p.dir - 1
	if ndir < 0{
		ndir = 3
	}
	valid, _ = nextPointValid(m, point{x: p.x, y: p.y, dir: ndir})
	if valid {
		out = append(out, edge{p:point{x: p.x, y: p.y, dir: ndir}, score: 1000})
	}

	// turn right
	ndir = p.dir + 1
	if ndir > 3{
		ndir = 0
	}
	valid, _ = nextPointValid(m, point{x: p.x, y: p.y, dir: ndir})
	if valid {
		out = append(out, edge{p:point{x: p.x, y: p.y, dir: ndir}, score: 1000})
	}

	return out
}

func nextPointValid(m [][]string, p point) (bool, point) {
	nextPoint := p
	switch p.dir {
	case 0:
		nextPoint.y -= 1
	case 1:
		nextPoint.x += 1
	case 2:
		nextPoint.y += 1
	case 3:
		nextPoint.x -= 1
	default:
		// do nothing
	}
	if (nextPoint.y > 0 && nextPoint.y < len(m) && nextPoint.x > 0 && nextPoint.x < len(m[0])) && (m[nextPoint.y][nextPoint.x] == "." || m[nextPoint.y][nextPoint.x] == "S" || m[nextPoint.y][nextPoint.x] == "E") { // next point directionally
		return true, nextPoint
	}
	return false, p
}

// Priority queue item 
type Item struct { 
	point point 
	priority int 
	index int 
	} 
// Priority queue implementation 
type PriorityQueue []*Item 
func (pq PriorityQueue) Len() int { 
	return len(pq) 
} 

func (pq PriorityQueue) Less(i, j int) bool { 
	return pq[i].priority < pq[j].priority 
} 

func (pq PriorityQueue) Swap(i, j int) { 
	pq[i], pq[j] = pq[j], pq[i] 
	pq[i].index = i 
	pq[j].index = j 
} 

func (pq *PriorityQueue) Push(x interface{}) { 
	n := len(*pq) 
	item := x.(*Item) 
	item.index = n 
	*pq = append(*pq, item) 
} 

func (pq *PriorityQueue) Pop() interface{} { 
	old := *pq 
	n := len(old) 
	item := old[n-1] 
	old[n-1] = nil // avoid memory leak 
	item.index = -1 // for safety 
	*pq = old[0 : n-1] 
	return item 
} 

func (pq *PriorityQueue) update(item *Item, point point, priority int) { 
	item.point = point 
	item.priority = priority 
	heap.Fix(pq, item.index) 
} 

// Dijkstra's algorithm implementation 
func dijkstra(weighted map[point][]edge, start point, grid [][]string) (int, [][]point) { 
	dist := make(map[point]int) 
	previous := make(map[point][]point) 
	for p := range weighted { 
		dist[p] = int(^uint(0) >> 1) // Set to maximum int value 
	} 
	dist[start] = 0 
	pq := make(PriorityQueue, 0) 
	heap.Init(&pq) 
	heap.Push(&pq, &Item{ point: start, priority: 0, }) 
	visited := make(map[point]bool) 
	var endPoint *point 
	for pq.Len() > 0 { 
		u := heap.Pop(&pq).(*Item).point 
		if grid[u.y][u.x] == "E" { 
			endPoint = &u 
			break 
		} 
		if visited[u] { 
			continue 
		} 
		visited[u] = true 
		for _, e := range weighted[u] { 
			if visited[e.p] { 
				continue 
			} 
			alt := dist[u] + e.score 
			if alt < dist[e.p] { 
				dist[e.p] = alt 
				previous[e.p] = []point{u} 
				heap.Push(&pq, &Item{ point: e.p, priority: alt, }) 
			} else if alt == dist[e.p] { 
					previous[e.p] = append(previous[e.p], u) 
			} 
		} 
	} 
	if endPoint == nil { 
		return -1, nil // If no path found 
	} 
	// Reconstruct all paths 
	paths := reconstructPaths(previous, *endPoint) 
	return dist[*endPoint], paths 
} 
func reconstructPaths(previous map[point][]point, end point) [][]point { 
	var paths [][]point 
	var stack [][]point 
	stack = append(stack, []point{end}) 
	for len(stack) > 0 { 
		path := stack[len(stack)-1] 
		stack = stack[:len(stack)-1] 
		cur := path[0] 
		if len(previous[cur]) == 0 { 
			paths = append(paths, path) 
		} else { 
			for _, prev := range previous[cur] { 
				newPath := make([]point, len(path)+1) 
				copy(newPath[1:], path) 
				newPath[0] = prev 
				stack = append(stack, newPath) 
			} 
		} 
	} 
	// Reverse paths to start from the beginning 
	for i, path := range paths { 
		for j := 0; j < len(path)/2; j++ { 
			path[j], path[len(path)-1-j] = path[len(path)-1-j], path[j] 
		} 
	paths[i] = path 
	} 
	return paths
}