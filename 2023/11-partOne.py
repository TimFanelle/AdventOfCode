def manhattanDist(x0, y0, x1, y1):
	return abs(x0-x1)+abs(y0-y1)

universe = []
with open("AdventOfCode/2023/11-input.txt", 'r') as textIn:
	for line in textIn:
		temp = []
		for char in line:
			if char != "\n":
				temp.append(char)
		universe.append(temp)
		if not('#' in line):
			universe.append(temp)

import numpy as np

universe = np.array(universe).transpose().tolist()
i = 0
while i < len(universe):
	found = False
	for j in range(len(universe[i])):
		if universe[i][j] == "#":
			found = True
			break
	if not(found):
		universe.insert(i, universe[i])
		i += 1
	i+=1

universe = np.array(universe).transpose().tolist()

galaxies = []
galNum = 0
for i in range(len(universe)):
	for j in range(len(universe[i])):
		if universe[i][j] == "#":
			galaxies.append((i,j, galNum))
			galNum += 1
galPairs = []
galDist = []
for i in range(len(galaxies)):
	for j in range(len(galaxies)):
		if galaxies[j][2] > i:
			galPairs.append((galaxies[i], galaxies[j]))
			galDist.append(manhattanDist(galaxies[i][0], galaxies[i][1], galaxies[j][0], galaxies[j][1]))
#print(len(galPairs))

print(sum(galDist))