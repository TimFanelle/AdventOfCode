class Node:
	def __init__(self, loc, val):
		self.loc = loc
		self.val = val
		self.options = []
	
	def updateOptions(self, grid):
		for item in [1,-1, 1j, -1j]:
			updatedLoc = self.loc + item
			if updatedLoc.real >= 0 and updatedLoc.imag >= 0:
				try:
					if grid[int(updatedLoc.real)][int(updatedLoc.imag)].val != "#":
						self.options.append(item)
				except(IndexError):
					pass
			else:
				pass

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

curLocs = [startingLoc]
numSteps = 64
for _ in range(numSteps):
	nexLocs = []
	for item in curLocs:
		for op in grid[int(item.real)][int(item.imag)].options:
			nexLocs.append(item+op)
	curLocs = list(set(nexLocs))
print(len(curLocs))