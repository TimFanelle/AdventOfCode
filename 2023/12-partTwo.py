class memoize():
	def __init__(self, func):
		self.func = func
		self.cache = {}
	def __call__(self, *x):
		if x not in self.cache:
			self.cache[x] = self.func(*x)
		return self.cache[x]

@memoize	
def nextInterest(strIn):
	for i in range(len(strIn)):
		if strIn[i] == "#" or strIn[i] == "?":
			return i
	return len(strIn)

@memoize
def checkHowMany(strIn, vals):
	minLen = 0
	for item in vals:
		minLen += item + 1
	minLen -= 1
	if len(strIn) < minLen:
		return 0
	if (strIn == "" or (not("#" in strIn) and not("?" in strIn))):
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
		st_temp = ""
		for _ in range(5):
			st_temp += st + "?"
		st = st_temp[:-1]
		vals = vals*5
		total += checkHowMany(st, tuple(vals))

print(total)