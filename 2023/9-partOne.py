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
		lastVals.append(vals[-1])
		temp = np.diff(vals)
		diffVals.append(temp[-1])
		while not(allZero(temp)):
			lastVals.append(temp[-1])
			temp = np.diff(temp)
			diffVals.append(temp[-1])
		total += sum(lastVals)
print(total)