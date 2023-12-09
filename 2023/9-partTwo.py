import numpy as np

def allZero(l):
	yes = True
	for item in l:
		if item != 0:
			yes = False
			break
	return yes

total = 0
with open("AdventOfCode/2023/9-input.txt", 'r') as textIn:
	for line in textIn:
		t_val = line.split()
		vals = [int(v) for v in t_val]
		lastVals = []
		diffVals = []
		lastVals.append(vals[0])
		temp = np.diff(vals)
		diffVals.append(temp[0])
		while not(allZero(temp)):
			lastVals.append(temp[0])
			temp = np.diff(temp)
			diffVals.append(temp[0])
		p = 0
		for i in range(len(lastVals)-1, -1, -1):
			p = lastVals[i] - p
		total += p
print(total)