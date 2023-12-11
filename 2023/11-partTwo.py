def manhattanDist(x0, y0, x1, y1):
	return abs(x0-x1)+abs(y0-y1)

def transposeArray(arr):
	return list(map(list, zip(*arr)))

universe = []
universeGrowth = 1000000
curRow = 0
with open("AdventOfCode/2023/11-input.txt", 'r') as textIn:
	for line in textIn:
		temp = []
		for char in line:
			if char != "\n":
				temp.append((char, curRow, 0))
		universe.append(temp)
		if not('#' in line):
			curRow += (universeGrowth - 1)
		curRow += 1

universe = transposeArray(universe)
curCol = 0
for i in range(len(universe)):
	found = False
	for j in range(len(universe[i])):
		if universe[i][j][0] == "#":
			found = True
		universe[i][j] = (universe[i][j][0], universe[i][j][1], curCol) 
			#break
	if not(found):
		curCol += (universeGrowth - 1)
	curCol += 1

universe = transposeArray(universe)

galaxies = []
galNum = 0
for i in range(len(universe)):
	for j in range(len(universe[i])):
		if universe[i][j][0] == "#":
			galaxies.append((universe[i][j][1],universe[i][j][2], galNum))
			galNum += 1
galPairs = []
galDist = []
for i in range(len(galaxies)):
	for j in range(len(galaxies)):
		if galaxies[j][2] > i:
			galPairs.append((galaxies[i], galaxies[j]))
			galDist.append(manhattanDist(galaxies[i][0], galaxies[i][1], galaxies[j][0], galaxies[j][1]))

print(sum(galDist))