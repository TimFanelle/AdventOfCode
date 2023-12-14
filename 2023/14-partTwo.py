rocks = []
with open("AdventOfCode/2023/14-input.txt", 'r') as textIn:
	for line in textIn:
		temp = []
		for char in line:
			if char != '\n':
				temp.append(char)
		rocks.append(temp)

def transposeArray(arr):
	return list(map(list, zip(*arr)))

def moveBoulders_Cycle(rocks):
	#North
	rocks = transposeArray(rocks)
	for col in rocks:
		nextAvail = len(col)
		for i in range(len(col)):
			if col[i] == "." and i < nextAvail:
				nextAvail = i
			elif col[i] == "#":
				nextAvail = i+1
			elif col[i] == "O" and nextAvail <= i:
				col[nextAvail] = "O"
				if nextAvail != i:
					col[i] = "."
				nextAvail += 1
	#West
	rocks = transposeArray(rocks)
	for col in rocks:
		nextAvail = len(col)
		for i in range(len(col)):
			if col[i] == "." and i < nextAvail:
				nextAvail = i
			elif col[i] == "#":
				nextAvail = i+1
			elif col[i] == "O" and nextAvail <= i:
				col[nextAvail] = "O"
				if nextAvail != i:
					col[i] = "."
				nextAvail += 1
	#South
	rocks = transposeArray(rocks)
	for col in rocks:
		nextAvail = 0
		for i in range(len(col)-1, -1, -1):
			if col[i] == "." and i > nextAvail:
				nextAvail = i
			elif col[i] == "#":
				nextAvail = i-1
			elif col[i] == "O" and nextAvail >= i:
				col[nextAvail] = "O"
				if nextAvail != i:
					col[i] = "."
				nextAvail -= 1
	#East
	rocks = transposeArray(rocks)
	for col in rocks:
		nextAvail = 0
		for i in range(len(col)-1, -1, -1):
			if col[i] == "." and i > nextAvail:
				nextAvail = i
			elif col[i] == "#":
				nextAvail = i-1
			elif col[i] == "O" and nextAvail >= i:
				col[nextAvail] = "O"
				if nextAvail != i:
					col[i] = "."
				nextAvail -= 1
	return rocks

def calculateLoad(rocks):
	startVal = len(rocks)
	totalLoad = 0
	for i in range(startVal):
		for item in rocks[i]:
			if item == "O":
				totalLoad += (startVal - i)
	return totalLoad

states = []
goUntil = 1000000000
foundState = False
while not(foundState):
	rocks = moveBoulders_Cycle(rocks)
	if not(rocks in states):
		states.append(rocks)
	else:
		foundState = True

curr = len(states)
stateIndex = states.index(rocks)
loopLen = curr-stateIndex
print(loopLen)
endState = stateIndex + (goUntil-curr)%loopLen
checkEndState = states[endState-1]
score = calculateLoad(checkEndState)
print(score)