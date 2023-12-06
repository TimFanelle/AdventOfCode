lens = []
with open("AdventOfCode/2023/6-input.txt", 'r') as textIn:
	vals = []
	for line in textIn:
		nums = [int(k) for k in line.split(":")[1].split()]
		vals.append(nums)
	vals = [list(l) for l in zip(*vals)]
	for item in vals:
		low_found = False
		high_found = False
		start = item[0]
		dist = item[1]
		low = 0
		high = item[0]
		while not high_found:
			high -= 1
			if high * (start - high) > dist:
				high_found = True
		while not low_found:
			low += 1
			if low * (start - low) > dist:
				low_found = True
		lens.append(high - low + 1)

from functools import reduce
print(reduce((lambda x,y: x*y), lens))
