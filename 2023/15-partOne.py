def hashFunc(strIn):
	vals = [ord(c) for c in strIn]
	cur = 0
	for v in vals:
		cur += v
		cur *= 17
		cur %= 256

	return cur

with open("AdventOfCode/2023/15-input.txt", 'r') as textIn:
	strins = textIn.readline().split(',')

total = 0
for s in strins:
	total += hashFunc(s)

print(total)