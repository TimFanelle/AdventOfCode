steps = 0
locs = dict()
currLoc = []
with open("AdventOfCode/2023/8-input.txt", 'r') as textIn:
	directions = textIn.readline()[:-1]
	_ = textIn.readline()
	for line in textIn:
		k, other = line.split("=")
		left, right = other.split(", ")
		left = left[2:]
		if right[-1] == "\n":
			right = right[:-2]
		else:
			right = right[:-1]
		locs[k[:-1]] = (left, right)
		if k[:-1][-1] == 'A':
			currLoc.append(k[:-1])

cycleLengths = [0]*len(currLoc)

def allEndInZ(locList):
	yes = True
	i = 0
	for item in locList:
		if item[-1] != "Z":
			yes = False
		if cycleLengths[i] == 0 and item[-1] == "Z":
			cycleLengths[i] = steps
		i +=1
	return yes

def allCycleLengthsFilled():
	yes = True
	for item in cycleLengths:
		if item == 0:
			yes = False
	return yes

lr = {"L":0, "R":1}
dir = 0
while not(allEndInZ(currLoc)) and not(allCycleLengthsFilled()):
	for item in range(len(currLoc)):
		currLoc[item] = locs[currLoc[item]][lr[directions[dir]]]
	steps += 1
	dir += 1
	if dir > len(directions)-1:
		dir = 0

from math import lcm
print(lcm(*cycleLengths))
