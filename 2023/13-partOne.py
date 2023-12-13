total = 0
def isSymmetrical(arr):
	if len(arr) <= 1:
		return False
	for i in range(len(arr)//2):
		if arr[i] != arr[-i-1]:
			return False
	return True

def transposeArray(arr):
	return list(map(list, zip(*arr)))

def determineMirror(arr):
	for i in range(2):
		for rowNum in range(len(arr)):
			if (rowNum+1)*2 < len(arr):
				temp = isSymmetrical(arr[:(rowNum+1)*2])
			else:
				temp = isSymmetrical(arr[-((len(arr)-1)-rowNum)*2:])
			if temp:
				if i == 1:
					return rowNum + 1
				if i == 0:
					return (rowNum + 1) * 100
		arr = transposeArray(arr)

images = []
with open('AdventOfCode/2023/13-input.txt', 'r') as textIn:
	im = []
	for line in textIn:
		if line != '\n':
			temp = []
			for char in line:
				if char != '\n':
					temp.append(char)
			im.append(temp)
		else:
			images.append(im)
			im = []
	images.append(im)

for image in images:
	total += determineMirror(image)

print(total)