class Node:
	def __init__(self, loc, val):
		self.loc = loc
		self.val = val
		self.options = []
	
	def updateOptions(self, grid):
		for item in [1,-1, 1j, -1j]:
			updatedLoc = self.loc + item
			try:
				if grid[int(updatedLoc.real)][int(updatedLoc.imag)].val != "#":
					self.options.append(item)
			except(IndexError):
				if updatedLoc.real > len(grid)-1:
					updatedLoc = complex(updatedLoc.real%len(grid), updatedLoc.imag)
				if updatedLoc.imag > len(grid[0]) -1:
					updatedLoc = complex(updatedLoc.real, updatedLoc.imag%len(grid[0]))
				if grid[int(updatedLoc.real)][int(updatedLoc.imag)].val != "#":
					self.options.append(item)

def bfs(grid, start, stepCount):
	visited = set()
	q = deque([(start, 0)])
	while q:
		val, steps = q.popleft()
		if steps > stepCount:
			continue
		updatedVal = complex(val.real%len(grid), val.imag%len(grid[0]))
		for op in grid[int(updatedVal.real)][int(updatedVal.imag)].options:
			newVal = val + op
			if newVal not in visited:
				visited.add(newVal)
				q.append((newVal, steps + 1))
	return len([v for v in visited if (v.real + v.imag)%2 == stepCount % 2])

from collections import deque
import numpy as np
grid = []
with open("AdventOfCode/2023/21-input.txt", 'r') as textIn:
	a = 0
	for line in textIn:
		b = 0
		temp = []
		for c in line:
			if c != "\n":
				temp.append(Node(complex(a,b), c))
			b += 1
		grid.append(temp)
		a += 1

startingLoc = None
for row in grid:
	for item in row:
		item.updateOptions(grid)
		if item.val == "S":
			startingLoc = item.loc
			item.val = "."

'''
curLocs = [startingLoc]
numSteps = 50
for _ in range(numSteps):
	nexLocs = []
	for item in curLocs:
		try:
			for op in grid[int(item.real)][int(item.imag)].options:
				nexLocs.append(item+op)
		except(IndexError):
			updatedLoc = item
			if updatedLoc.real > len(grid)-1 or updatedLoc.real < -len(grid):
				updatedLoc = complex(updatedLoc.real%len(grid), updatedLoc.imag)
			if updatedLoc.imag > len(grid[0]) -1 or updatedLoc.imag < len(grid[0]):
				updatedLoc = complex(updatedLoc.real, updatedLoc.imag%len(grid[0]))
			for op in grid[int(updatedLoc.real)][int(updatedLoc.imag)].options:
				nexLocs.append(item+op)
	curLocs = list(set(nexLocs))
print(len(curLocs))
'''

'''
polynomial fitting suggested by:
https://www.reddit.com/r/adventofcode/comments/18nevo3/comment/kec8fxc/?utm_source=share&utm_medium=web2x&context=3
'''
visited = bfs(grid, startingLoc, 64)
print(visited)

y_vals = [bfs(grid, startingLoc, stepCount) for stepCount in [65, 196, 327, 65+(131*3), 65+(131*4)]]
# 65, 65 + 131, 65 + 131 * 2
x_values = np.array([0, 1, 2, 3, 4]) # need to ensure there is enough data or it will bounce around the answer

target = (26501365 - 65) // 131
coefficients = np.polyfit(x_values, y_vals, 2)
result = np.polyval(coefficients, target)
print(np.round(result, 0))