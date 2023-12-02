total = 0
digits = {"one":"o1e", "two":"t2o", "three":"t3e", "four":"f4r", "five":"f5e", "six":"s6x", "seven":"s7n", "eight":"e8t", "nine":"n9e"}
digit_names = list(digits.keys())
with open("AdventOfCode/2023/1-input.txt", 'r') as textIn:
	for line in textIn:
		str = ""
		for i in range(len(line)):
			for d in digit_names:
				if d in line[i:i+5]:
					line = line.replace(d, digits[d])
		for i in range(len(line)):
			try:
				int(line[i])
				str += line[i]
			except:
				pass
		combo_int = str[0] + str[-1]
		total += int(combo_int)

print(total)
		