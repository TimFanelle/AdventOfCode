coords = dict()
grid = []
with open("AdventOfCode/2023/16-input.txt", 'r') as textIn:
	i = 0
	for line in textIn:
		temp = []
		j = 0
		for c in line:
			if c != '\n':
				temp.append((c, False))
				if c in ['|', '-', '/', '\\']:
					coords[complex(i,j)] = []
				j += 1
		grid.append(temp)
		i += 1

cur = [(0-1j, 1j)]
coord_keys = coords.keys()
while len(cur) > 0:
	next_cur = []
	for ind in cur:
		dir = ind[1]
		val = ind[0]
		if val.imag >= 0:
			grid[int(val.real)][int(val.imag)] = (grid[int(val.real)][int(val.imag)][0], True)
		val += dir		
		if val in coord_keys:
			working = []
			grid_val = grid[int(val.real)][int(val.imag)][0]
			if grid_val == '-':
				if dir == 1j or dir == -1j:
					working.append((val, dir))
				else:
					if not(dir in coords[val]):
						working.append((val, -1j))
						working.append((val, 1j))
						coords[val].append(dir)
			elif grid_val == '|':
				if dir == 1 or dir == -1:
					working.append((val, dir))
				else:
					if not(dir in coords[val]):
						working.append((val, -1))
						working.append((val, 1))
						coords[val].append(dir)
			elif grid_val == "/":
				if not(dir in coords[val]):
					if dir == 1:
						working.append((val, -1j))
					elif dir == -1:
						working.append((val, 1j))
					elif dir == 1j:
						working.append((val, -1))
					else:
						working.append((val, 1))
					coords[val].append(dir)
			elif grid_val == "\\":
				if not(dir in coords[val]):
					if dir == 1:
						working.append((val, 1j))
					elif dir == -1:
						working.append((val, -1j))
					elif dir == 1j:
						working.append((val, 1))
					else:
						working.append((val, -1))
					coords[val].append(dir)
			for item in working:
				if -1 < item[0].real < len(grid) and -1 < item[0].imag < len(grid[int(item[0].real)]):
					next_cur.append(item)
		else:
			if 0 <= val.real < len(grid) and 0 <= val.imag < len(grid[int(ind[0].real)]):
				next_cur.append((val, dir))
	cur = next_cur

total = 0
for line in grid:
	for item in line:
		if item[1]:
			total += 1
		else:
			pass

print(total)