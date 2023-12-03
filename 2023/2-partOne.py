total = 0
with open('AdventOfCode/2023/2-input.txt', 'r') as text:
	for game in text:
		(id, games) = game.split(':')
		(_, id) = id.split()
		id = int(id)
		sets = games.split(';')
		valid = True
		for g in sets:
			g_0 = g.split(',')
			checker = dict()
			for color in g_0:
				c = color.split()
				checker[c[1]] = int(c[0])
			if 'red' in checker.keys() and checker['red'] > 12:
				valid = False
			if 'green' in checker.keys() and checker['green'] > 13:
				valid = False
			if 'blue' in checker.keys() and checker['blue'] > 14:
				valid = False
		if valid:
			total += id
print(total)