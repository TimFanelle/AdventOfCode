with open("AdventOfCode/2023/5-input.txt", 'r') as textIn:
	seedCheck = dict()
	firstLine = textIn.readline()
	seeds = [int(k) for k in firstLine.split(":")[1].split()]
	#print(seeds)
	fullString = ""
	for line in textIn:
		fullString += line
	sections = fullString.split('\n\n')
	sect_dict = dict()
	for s in sections:
		temp = s.split(":\n")
		temp_0 = [k for k in temp[1].split('\n')]
		for j in range(len(temp_0)):
			temp_1 = temp_0[j].split()
			for i in range(len(temp_1)):
				temp_1[i] = int(temp_1[i])
			temp_0[j] = temp_1
		sect_dict[temp[0]] = temp_0
	#print(sect_dict)
	for seed in seeds:
		seedCheck[str(seed)] = []
		current_val = seed
		for k in sect_dict.keys():
			for trials in sect_dict[k]:
				if current_val >= trials[1] and current_val < trials[1]+trials[2]:
					current_val = trials[0] + (current_val-trials[1])
					break
			seedCheck[str(seed)].append(current_val)
	locs = []
	for k in seedCheck:
		#print(seedCheck[k])
		locs.append(seedCheck[k][-1])
	print(min(locs))