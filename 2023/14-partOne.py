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

def moveBoulders():
	global rocks
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
	rocks = transposeArray(rocks)

def calculateLoad():
	global rocks
	startVal = len(rocks)
	totalLoad = 0
	for i in range(startVal):
		for item in rocks[i]:
			if item == "O":
				totalLoad += (startVal - i)
	return totalLoad

moveBoulders()
print(calculateLoad())