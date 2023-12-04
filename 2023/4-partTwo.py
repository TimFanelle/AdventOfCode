with open("AdventOfCode/2023/4-input.txt", 'r') as textIn:
	all_instances = []
	i = 1
	for line in textIn:
		(id, nums) = line.split(":")
		(winning, selected) = nums.split("|")
		winning = [int(k) for k in winning.split()]
		selected = [int(k) for k in selected.split()]
		all_instances.append([(winning, selected)])
		card = 0
		for item in winning:
			if item in selected:
				card += 1
		cards_to_add = []
		for j in range(1, card+1):
			cards_to_add.append(i+j)
		all_instances[i-1].append(cards_to_add)
		i += 1
for u in all_instances:
	print(u)
cards_to_add = []
for u in all_instances:
	cards_to_add += u[1]
l = 0
while l < len(cards_to_add):
	if cards_to_add[l]-1 < len(all_instances):
		cards_to_add += all_instances[cards_to_add[l]-1][1]
	l += 1
cards_to_add += [j+1 for j in range(len(all_instances))]
print(len(cards_to_add))