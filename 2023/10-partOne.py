class Node:
	def __init__(self, sym, x, y):
		self.x = x
		self.y = y
		self.symbol = sym
		self.value = 0
		self.visited = False
		self.dir = self.getDir()
		self.neighbors = None
	
	def getDir(self):
		match self.symbol:
			case "S":
				return [1,1,1,1]
			case "F":
				return [0,1,0,1]
			case "-":
				return [0,0,1,1]
			case "J":
				return [1,0,1,0]
			case "|":
				return [1,1,0,0]
			case "L":
				return [1,0,0,1]
			case "7":
				return [0,1,1,0]
			case default:
				return [0,0,0,0]
	
	def getNeighbors(self, grid_rep):
		n = []
		if self.dir[0] and self.x > 0 and grid_rep[self.x-1][self.y].dir[1]:
			n.append(grid_rep[self.x-1][self.y])
		#else:
		#	n.append(None)
		
		if self.dir[1] and self.x < len(grid_rep)-1 and grid_rep[self.x+1][self.y].dir[0]:
			n.append(grid_rep[self.x+1][self.y])
		#else:
		#	n.append(None)
		
		if self.dir[2] and self.y > 0 and grid_rep[self.x][self.y-1].dir[3]:
			n.append(grid_rep[self.x][self.y-1])
		#else:
		#	n.append(None)
		
		if self.dir[3] and self.y < len(grid_rep[self.x])-1 and grid_rep[self.x][self.y+1].dir[2]:
			n.append(grid_rep[self.x][self.y+1])
		#else:
		#	n.append(None)

		self.neighbors = n

grid = []
startNodeValues = (None, None)
with open("AdventOfCode/2023/10-input.txt", 'r') as textIn:
	i = 0
	for line in textIn:
		temp = []
		j = 0
		for symbol in line:
			if symbol != '\n':
				temp.append(Node(symbol, i, j))
				if symbol == "S":
					startNodeValues = (i, j)
			j += 1
		grid.append(temp)
		i += 1

for row in grid:
	for col in row:
		col.getNeighbors(grid)

print(startNodeValues)

def traverseGrid(startingTuple):
	check_nodes = []
	startNode = grid[startingTuple[0]][startingTuple[1]]
	for n in startNode.neighbors:
		check_nodes.append((startNode, n))
	while len(check_nodes) > 0:
		next_nodes = []
		for nod in check_nodes:
			if nod[0].value + 1 < nod[1].value or nod[1].value == 0:
				nod[1].value = nod[0].value +1
				for n in nod[1].neighbors:
					if n != nod[0]:
						next_nodes.append((nod[1], n))
		check_nodes = next_nodes

traverseGrid(startNodeValues)
highest = 0
for row in grid:
	r = ""
	for col in row:
		if col.value > highest:
			highest = col.value
		r += str(col.value) + " "
	#print(r)

print("Highest value is:", highest)