import copy
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

def checker(i,j, dirIn):
	cur = [(complex(i,j), dirIn)]
	coords_cop = copy.deepcopy(coords)
	coord_keys = coords.keys()
	grid_cop = copy.deepcopy(grid)
	while len(cur) > 0:
		next_cur = []
		for ind in cur:
			dir = ind[1]
			val = ind[0]
			if val.imag >= 0 and val.real >= 0 and val.imag < len(grid_cop[0]) and val.real < len(grid_cop):
				grid_cop[int(val.real)][int(val.imag)] = (grid_cop[int(val.real)][int(val.imag)][0], True)
			val += dir		
			if val in coord_keys:
				working = []
				grid_val = grid_cop[int(val.real)][int(val.imag)][0]
				if grid_val == '-':
					if dir == 1j or dir == -1j:
						working.append((val, dir))
					else:
						if not(dir in coords_cop[val]):
							working.append((val, -1j))
							working.append((val, 1j))
							coords_cop[val].append(dir)
				elif grid_val == '|':
					if dir == 1 or dir == -1:
						working.append((val, dir))
					else:
						if not(dir in coords_cop[val]):
							working.append((val, -1))
							working.append((val, 1))
							coords_cop[val].append(dir)
				elif grid_val == "/":
					if not(dir in coords_cop[val]):
						if dir == 1:
							working.append((val, -1j))
						elif dir == -1:
							working.append((val, 1j))
						elif dir == 1j:
							working.append((val, -1))
						else:
							working.append((val, 1))
						coords_cop[val].append(dir)
				elif grid_val == "\\":
					if not(dir in coords_cop[val]):
						if dir == 1:
							working.append((val, 1j))
						elif dir == -1:
							working.append((val, -1j))
						elif dir == 1j:
							working.append((val, 1))
						else:
							working.append((val, -1))
						coords_cop[val].append(dir)
				for item in working:
					if -1 < item[0].real < len(grid_cop) and -1 < item[0].imag < len(grid_cop[int(item[0].real)]):
						next_cur.append(item)
			else:
				if 0 <= val.real < len(grid_cop) and 0 <= val.imag < len(grid_cop[int(val.real)]):
					next_cur.append((val, dir))
		cur = next_cur

	total = 0
	for line in grid_cop:
		for item in line:
			if item[1]:
				total += 1
			else:
				pass
	return total

m = 0
for i in range(len(grid)):
	print("i=", i)
	m = max(checker(i, -1, 1j), m)
	m = max(checker(i, len(grid[i]), -1j), m)
for j in range(len(grid[0])):
	print("j=", j)
	m = max(checker(-1, j, 1), m)
	m = max(checker(len(grid), j, -1), m)
print(m)