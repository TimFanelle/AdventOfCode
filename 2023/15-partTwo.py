def hashFunc(strIn):
	vals = [ord(c) for c in strIn]
	cur = 0
	for v in vals:
		cur += v
		cur *= 17
		cur %= 256
	return cur

def checkBox(box, label, focal):
	for item in range(len(box)):
		if label in box[item]:
			l, f = box[item].split("=")
			return (item, focal)
	return (-1, focal)

def determineIndex(box, label):
	for item in range(len(box)):
		if label in box[item]:
			return item
	return -1

with open("AdventOfCode/2023/15-input.txt", 'r') as textIn:
	strins = textIn.readline().split(',')

total = 0
boxes = {}
for s in strins:
	if "=" in s:
		label, focal = s.split("=")
		focal = int(focal)
		boxNum = hashFunc(label)
		if str(boxNum) in boxes.keys():
			ind, foc = checkBox(boxes[str(boxNum)], label, focal)
			if ind != -1:
				boxes[str(boxNum)][ind] = label+"="+str(foc)
			else:
				boxes[str(boxNum)].append(s)
		else:
			boxes[str(boxNum)] = [s]
	else:
		label, _ = s.split('-')
		boxNum = hashFunc(label)
		if str(boxNum) in boxes.keys():
			revBoxInd = determineIndex(boxes[str(boxNum)], label)
			if revBoxInd != -1:
				boxes[str(boxNum)] = boxes[str(boxNum)][:revBoxInd] + boxes[str(boxNum)][revBoxInd+1:]
		else:
			pass

for k in boxes.keys():
	for j in range(len(boxes[k])):
		_, val = boxes[k][j].split("=")
		val = int(val)
		total += (int(k)+1)*(j+1)*val
print(total)