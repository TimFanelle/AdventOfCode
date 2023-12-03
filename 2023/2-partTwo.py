total = 0
with open('AdventOfCode/2023/2-input.txt', 'r') as text:
	for game in text:
		(id, games) = game.split(':')
		(_, id) = id.split()
		id = int(id)
		sets = games.split(';')
		minimums = {'red': 0, 'green': 0, 'blue': 0}
		for g in sets:
			g_0 = g.split(',')
			checker = dict()
			for color in g_0:
				c = color.split()
				checker[c[1]] = int(c[0])
			for k in checker.keys():
				if checker[k] > minimums[k]:
					minimums[k] = checker[k]
		print(minimums)
		power = 1
		for k in minimums.keys():
			power *= minimums[k]
		print(power)
		total += power
print(total)