workflows = dict()
items = []
shifted = False
with open("AdventOfCode/2023/19-input.txt") as textIn:
	for line in textIn:
		if shifted:
			temp = dict()
			temp_line = line.replace('\n', '')[1:-1].split(",")
			for item in temp_line:
				t = item.split("=")
				temp[t[0]] = int(t[1])
			items.append(temp)
		elif line == '\n':
			shifted = True
		else:
			temp_line = line.split("{")
			rules = temp_line[1].split("}")
			rules = rules[0].split(",")
			workflows[temp_line[0]] = rules

workflow_names = workflows.keys()
def ruleOutput(item, wfn):
	determination = None
	curWorkflow = "in"
	while not(determination in ["R", "A"]):
		curLen = len(workflows[curWorkflow])
		curRules = workflows[curWorkflow]
		for rule_no in range(curLen):
			if rule_no == curLen-1:
				determination = curRules[rule_no]
			else:
				rule, out = curRules[rule_no].split(":")
				if "<" in rule:
					rule = rule.split("<")
					if item[rule[0]] < int(rule[1]):
						determination = out
						break
				else:
					rule = rule.split(">")
					if item[rule[0]] > int(rule[1]):
						determination = out
						break
		if determination in wfn:
			curWorkflow = determination
	return determination

total = 0
for i in items:
	det = ruleOutput(i, workflow_names)
	if det == "A":
		for k in i.keys():
			total += i[k]

print(total)
