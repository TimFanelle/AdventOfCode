coords_in = []
with open("AdventOfCode/2023/22-input.txt", 'r') as textIn:
	for line in textIn:
		bottom, top = line.split("~")
		bottom = tuple([int(a) for a in bottom.split(",")])
		top = tuple([int(a) for a in top.split(",")])
		coords_in.append((bottom, top))
coords_in.sort(key=lambda k:k[0][2])

def determineIssues(block, finalizedBlocks):
	if block[0][2] < 1:
		return True
	for item in finalizedBlocks:
		if item[1][2] >= block[0][2] and item[0][2] <= block[1][2]:
			a = []
			for x in range(item[0][0], item[1][0]+1):
				for y in range(item[0][1], item[1][1]+1):
					a.append((x,y))
			a = set(a)
			b = []
			for x in range(block[0][0], block[1][0]+1):
				for y in range(block[0][1], block[1][1]+1):
					b.append((x,y))
			b = set(b)
			if len(b.intersection(a)) > 0:
				return True
	return False

def dropBlocks(coords_in):
	finalizedBlocks = []
	for block in range(len(coords_in)):
		if block == 0:
			if coords_in[block][0][2] != 1:
				dif = (1-coords_in[block][0][2])
				newBottom = (coords_in[block][0][0], coords_in[block][0][1], coords_in[block][0][2]+dif)
				newTop = (coords_in[block][1][0], coords_in[block][1][1], coords_in[block][1][2]+dif)
				coords_in[block] = (newBottom, newTop)
			finalizedBlocks.append(coords_in[block])
		else:
			issues = False
			curBlock = coords_in[block]
			while not(issues):
				newBottom = (curBlock[0][0], curBlock[0][1], curBlock[0][2]-1)
				newTop = (curBlock[1][0], curBlock[1][1], curBlock[1][2]-1)
				curBlock = (newBottom, newTop)
				issues = determineIssues(curBlock, finalizedBlocks)
			newBottom = (curBlock[0][0], curBlock[0][1], curBlock[0][2]+1)
			newTop = (curBlock[1][0], curBlock[1][1], curBlock[1][2]+1)
			curBlock = (newBottom, newTop)
			finalizedBlocks.append(curBlock)
	return finalizedBlocks

blocks = dropBlocks(coords_in)
removable = 0
for b in blocks:
	height = b[1][2]
	fin = [it for it in blocks if it[1][2] <= height and it != b]
	temp = [it for it in blocks if it[0][2] == height+1]
	iss = False
	for bl in temp:
		test = ((bl[0][0], bl[0][1], bl[0][2]-1), (bl[1][0], bl[1][1], bl[1][2]-1))
		if not(determineIssues(test, fin)):
			iss = True
			break
	if iss:
		pass
	else:
		removable += 1

print(removable)	
