total = 0
with open('AdventOfCode/2023/3-input.txt', 'r') as textIn:
	two_drep = []
	row = 0
	two_drep.append([])
	symbol_spaces = []
	for line in textIn:
		col = 0
		digit_str = ""
		for char in line:
			try:
				int(char)
				digit_str += char
			except:
				if len(digit_str) > 0:
					for _ in range(len(digit_str)):
						two_drep[row].append(int(digit_str))
					digit_str = ""
				if char != '.' and char != "\n":
					symbol_spaces.append((row, col))
					two_drep[row].append("X")
				else:
					two_drep[row].append(0)
			col += 1
		two_drep.append([])
		row += 1

if len(two_drep[row-1]) < len(two_drep[row-2]):
	for _ in range(len(two_drep[row-2])-len(two_drep[row-1])):
		two_drep[row-1].append(0)

#for l in two_drep:
#	print(l)
for pair in symbol_spaces:
	spaces_to_check = []
	scene_nums = []
	spaces_to_check.append((pair[0]-1, pair[1]-1))
	spaces_to_check.append((pair[0]-1, pair[1]))
	spaces_to_check.append((pair[0]-1, pair[1]+1))
	spaces_to_check.append((pair[0], pair[1]-1))
	spaces_to_check.append((pair[0]+1, pair[1]-1))
	spaces_to_check.append((pair[0]+1, pair[1]))
	spaces_to_check.append((pair[0], pair[1]+1))
	spaces_to_check.append((pair[0]+1, pair[1]+1))
	for space in spaces_to_check:
		if space[0] >= 0 and space[0] < row:
			if space[1] >= 0 and space[1] < len(two_drep[0]):
				scene_nums.append(two_drep[space[0]][space[1]])
	scene_nums = set(scene_nums)
	#print(scene_nums)
	for n in scene_nums:
		total += n
print(total)