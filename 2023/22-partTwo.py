coords_inn = []
with open("AdventOfCode/2023/22-input.txt", 'r') as textIn:
	for line in textIn:
		bottom, top = line.split("~")
		bottom = tuple([int(a) for a in bottom.split(",")])
		top = tuple([int(a) for a in top.split(",")])
		coords_inn.append((bottom, top))
coords_inn.sort(key=lambda k:k[0][2])

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

def differencesBetween(la, lb):
	d = 0
	for i in range(len(la)):
		if la[i] != lb[i]:
			d += 1
	return d

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

blocks = dropBlocks(coords_inn)
fallen = 0
for b in range(len(blocks)):
	#height = blocks[b][1][2]
	toDrop = [lt for lt in blocks if lt != blocks[b]]
	#fin = [it for it in blocks if it[1][2] <= height and it != b]
	separate = toDrop.copy()
	nextUp = dropBlocks(separate)
	fallen += differencesBetween(toDrop, nextUp)

print(fallen)	
