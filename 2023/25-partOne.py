# solution from https://github.com/IgnacyBerent/AoC_2023_python/blob/master/day_25/task_1.py
from collections import defaultdict, deque

components = defaultdict(set)

# read in input
with open("25-input.txt", 'r') as inp:
	for line in inp:
		comp, connections = line.split(": ")
		if connections[-1] == "\n":
			connections = connections[:-1].split(" ")
		else:
			connections = connections.split(" ")
		for c in connections:
			components[comp].add(c)
			components[c].add(comp)
		
group_1 = 1
group_2 = 0

# select any node to start
starting_comp = list(components.keys())[0]

#cycle through every other node
for component in list(components.keys())[1:]:
	connects = 0 # each component will have connections back to the starting node (or any other node in the same group), if there are more than 3 then it is in the same group as the starting node and if there are 3 or less than it is in the opposite group
	used = {starting_comp}
	for s_component in components[starting_comp]:
		if s_component == component:
			connects += 1 # if the node connects back to itself
			continue
		qed = set()
		q = deque()
		q.append((s_component, [s_component]))
		found = False
		while q and not found and connects < 4:
			comp, path = q.popleft() # take the first component in the queue
			for c in components[comp]: # check all the components it connects to
				if component == c: # if it finds a path back to the initial component you are searching
					connects += 1
					used.update(path) # add to make sure they are not doubling back
					found = True
					break
				elif c not in qed and c not in path and c not in used: 
					q.append([c, path + [c]]) # extend the path
					qed.add(c)
	if connects >= 4:
		group_1 += 1 # same group as starting
	else:
		group_2 += 1 # opposite group

print(group_1 * group_2)