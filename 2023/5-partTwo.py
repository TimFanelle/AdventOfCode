'''
Based on solution by:
https://github.com/CalSimmon/advent-of-code/blob/main/2023/day_05/solution.py
'''

def convert_range_to_map(conversionMap):
	conversionList = []
	for m in conversionMap:
		conversionList.append([m[1], m[1]+m[2]-1, m[0]-m[1]])
	return sorted(conversionList, key=lambda x:x[0])

def convert_seed_range_list(seed_range, conversion_range):
	seed_range_list = []
	curr_value = seed_range[0]
	for i in range(len(conversion_range)):
		if curr_value < conversion_range[i][0]:  # If the first entry in conversion_range is higher than lowest seed, add range until that point
			seed_range_list.append([int(curr_value), (int(conversion_range[i][0] - 1))])
			curr_value = int(conversion_range[i][0])
		if curr_value >= int(conversion_range[i][0]) and curr_value <= int(conversion_range[i][1]):  # If between the current conversion range, add the range, unless we are at the end of the seed range
			max_value = int(conversion_range[i][1]) if not int(seed_range[1]) <= int(conversion_range[i][1]) else int(seed_range[1])  # Set max value to seed range end if conversion range is higher than seed range
			seed_range_list.append([(curr_value + conversion_range[i][2]), (max_value + conversion_range[i][2])])
			curr_value = max_value + 1
		if curr_value - 1 == int(seed_range[1]):  # If we've hit the end of the seed range, break
			break
	if curr_value < seed_range[1]:  # If we did not hit the end of the seed range, add the rest of the range
		seed_range_list.append([curr_value, seed_range[1]])
	
	return sorted(seed_range_list, key=lambda x: x[0])

lowest = None
with open("AdventOfCode/2023/5-input.txt", 'r') as textIn:
	firstLine = textIn.readline()
	seed_initial = [int(k) for k in firstLine.split(":")[1].split()]
	seed_ranges = [[seed_initial[i], seed_initial[i] + seed_initial[i+1]-1] for i in range(0, len(seed_initial), 2)]
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
		sect_dict[temp[0]] = convert_range_to_map(temp_0)	
	
	for seed_range in seed_ranges:
		curr_seed_range = [seed_range.copy()]  # Convert the first seed pair into a nested list so the future loops will function properly
		for curr_conv_map in sect_dict.keys():
			curr_output_list = []
			for entry in curr_seed_range:
				output_list = convert_seed_range_list(entry, sect_dict[curr_conv_map])  # Output the new range for each seed range supplied
				curr_output_list.append(output_list)  
			curr_output_list = [item for sublist in curr_output_list for item in sublist]  # Make sure output is only a single nested list sorted
			curr_output_list = sorted(curr_output_list, key=lambda x: x[0])
			curr_seed_range = curr_output_list.copy()  # Set the curr seed range to the output list and run through the next conversion map
		print(f"The lowest of seed range {seed_range} is {curr_seed_range[0][0]}")

		if lowest is None or curr_seed_range[0][0] < lowest:
			lowest = curr_seed_range[0][0]
print(lowest)