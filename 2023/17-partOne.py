'''
inspiration from:
https://github.com/michaeljgallagher/Advent-of-Code/blob/master/2023/17.py
'''
from heapq import heappop, heappush
from collections import defaultdict
from math import inf
grid = []

with open("AdventOfCode/2023/17-input.txt", 'r') as textIn:
	for line in textIn:
		temp = []
		for c in line:
			if c != "\n":
				temp.append(int(c))
		grid.append(temp)

class updatedComplex(complex):
	def __lt__(self, other):
		if self.real < other.real:
			return True
		elif self.real > other.real:
			return False
		else:
			if self.imag < other.imag:
				return True
			elif self.imag > other.imag:
				return False
			else:
				return False
	def __le__(self, other):
		if self.real < other.real:
			return True
		elif self.real > other.real:
			return False
		else:
			if self.imag < other.imag:
				return True
			elif self.imag > other.imag:
				return False
			else:
				return True

def dijkstra(grid, endpoint, minMove, maxMove):
	dists = defaultdict(lambda: inf)
	heap = [(0, (updatedComplex(0, 0), 1j)), (0, (updatedComplex(0, 0), 1))]
	while heap:
		cost, (loc, dir) = heappop(heap)
		if (int(loc.real), int(loc.imag)) == endpoint:
			return cost
		if cost > dists[loc, dir]:
			continue
		for nd in (updatedComplex(-int(dir.imag), int(dir.real)), updatedComplex(int(dir.imag), -int(dir.real))):
			ncost = cost
			for dist in range(1, maxMove+1):
				n = updatedComplex(loc.real + nd.real * dist, loc.imag + nd.imag * dist)
				if 0 <= int(n.real) < endpoint[0]+1 and 0 <= int(n.imag) < endpoint[1]+1:
					ncost += grid[int(n.real)][int(n.imag)]
					if dist < minMove:
						continue
					k = (n, nd)
					if ncost < dists[k]:
						dists[k] = ncost
						heappush(heap, (ncost, k))
	return -1

print(dijkstra(grid, (len(grid)-1, len(grid[0])-1), 1, 3))
