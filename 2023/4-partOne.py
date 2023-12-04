total = 0
with open("AdventOfCode/2023/4-input.txt", 'r') as textIn:
	for line in textIn:
		(id, nums) = line.split(":")
		(winning, selected) = nums.split("|")
		winning = [int(k) for k in winning.split()]
		selected = [int(k) for k in selected.split()]
		card = 0
		for item in winning:
			if item in selected:
				if card == 0:
					card = 1
				else:
					card *= 2
		total += card
print(total)