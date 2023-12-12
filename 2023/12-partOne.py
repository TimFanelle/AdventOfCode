def nextInterest(strIn):
	for i in range(len(strIn)):
		if strIn[i] == "#" or strIn[i] == "?":
			return i
	return len(strIn)

def checkHowMany(strIn, vals):
	minLen = 0
	for item in vals:
		minLen += item + 1
	minLen -= 1
	if len(strIn) < minLen:
		return 0
	elif (strIn == "" or (not("#" in strIn) and not("?" in strIn))):
		if len(vals) == 0:
			return 1
		else:
			return 0
	else:
		if strIn[0] == ".":
			return checkHowMany(strIn[nextInterest(strIn):], vals)
		elif strIn[0] == "#":
			if len(vals) == 0:
				return 0
			elif (len(strIn) >= vals[0]) and ("#"*vals[0] in strIn[:vals[0]+1].replace("?", "#") and (strIn[vals[0]:] == "" or strIn[vals[0]] == "." or strIn[vals[0]] == "?")):
				return checkHowMany(strIn[vals[0]+1:], vals[1:])
			else:
				return 0
		else:
			return checkHowMany("."+strIn[1:], vals) + checkHowMany("#"+strIn[1:], vals)

total = 0
with open("AdventOfCode/2023/12-input.txt", 'r') as textIn:
	for line in textIn:
		st, vals = line.split()
		vals = [int(v) for v in vals.split(",")]
		total += checkHowMany(st, vals)

print(total)