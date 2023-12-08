steps = 0
locs = dict()
with open("AdventOfCode/2023/8-input.txt", 'r') as textIn:
	directions = textIn.readline()[:-1]
	_ = textIn.readline()
	for line in textIn:
		k, other = line.split("=")
		left, right = other.split(", ")
		left = left[2:]
		right = right[:-2]
		locs[k[:-1]] = (left, right)

currLoc = 'AAA'
lr = {"L":0, "R":1}
dir = 0
while currLoc != 'ZZZ':
	currLoc = locs[currLoc][lr[directions[dir]]]
	steps += 1
	dir += 1
	if dir > len(directions)-1:
		dir = 0
print(steps)