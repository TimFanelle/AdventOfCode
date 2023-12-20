class moduleNode:
	def __init__(self, inStr, children):
		type_, name = self.getTypeName(inStr)
		self.type = type_
		self.name = name
		self.OnOff = False
		self.children = children
		self.lastPulses = {}
	
	def getTypeName(self, s):
		match s[0]:
			case "%":
				return "Flip", s[1:]
			case "&":
				return "Conj", s[1:]
			case default:
				return None, s
	
	def addLastPulseNames(self, names):
		for n in names:
			self.lastPulses[n] = False
	
	def receivePulse(self, pulseType, inName):
		pulsesToAdd = []
		if self.type == "Flip":
			if not(pulseType):
				self.OnOff = not self.OnOff
				for c in self.children:
					pulsesToAdd.append((c, self.OnOff, self.name))
		elif self.type == "Conj":
			self.lastPulses[inName] = pulseType
			allHigh = True
			for k in self.lastPulses.keys():
				if not(self.lastPulses[k]):
					allHigh = False
					break
			for c in self.children:
				pulsesToAdd.append((c, not(allHigh), self.name))
		else:
			for c in self.children:
				pulsesToAdd.append((c, pulseType, self.name))
		return pulsesToAdd


nodes = dict()
conjNames = []
with open("AdventOfCode/2023/20-input.txt", 'r') as textIn:
	for line in textIn:
		info, child = line.split(" -> ")
		child = child.replace("\n", "").split(", ")
		temp = moduleNode(info, child)
		nodes[temp.name] = temp
		if temp.type == "Conj":
			conjNames.append(temp.name)

for n in conjNames:
	temp = []
	for k in nodes.keys():
		if k != n:
			if n in nodes[k].children:
				temp.append(k)
	nodes[n].addLastPulseNames(temp)

rxFeedInNodes = []
for k in nodes.keys():
	if k != "rx":
		if "rx" in nodes[k].children:
			rxFeedInNodes.append(k)
nodesToCheck = dict()
for n in rxFeedInNodes:
	for k in nodes.keys():
		if k != n:
			if n in nodes[k].children:
				nodesToCheck[k] = -1

startPulse = [("broadcaster", False, None)]
low = 0
high = 0
buttonPresses = 0
buttonPressed = False
while not(buttonPressed):
	pulses = startPulse.copy()
	i = 0
	while i < len(pulses):
		n, p, f = pulses[i]
		if p:
			high += 1
		else:
			low += 1
		try:
			pulses += nodes[n].receivePulse(p, f)
			for j in rxFeedInNodes:
				for l in nodes[j].lastPulses.keys():
					if nodes[j].lastPulses[l]:
						nodesToCheck[l] = buttonPresses+1
		except:
			pass
		i += 1
	buttonPresses += 1
	allNodesFound = True
	for j in nodesToCheck.keys():
		if nodesToCheck[j] == -1:
			allNodesFound = False
	if allNodesFound:
		buttonPressed = True

print(nodesToCheck)
from math import lcm
values = []
for k in nodesToCheck.keys():
	values.append(nodesToCheck[k])
print(lcm(*values))