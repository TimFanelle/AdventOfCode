'''
See:
Shoelace method
Pick's theorem
'''
curLoc = [0, 0]
vertices = []
trenchLen = 0
convert = {'0': 'R', '1':'D', '2':'L', '3':'U'}
with open("AdventOfCode/2023/18-input.txt", 'r') as textIn:
	for line in textIn:
		dir, l, hex = line.split()
		l = hex[2:-2]
		dir = convert[hex[-2]]
		match dir:
			case "R":
				curLoc[0] += int(l, 16)
			case "L":
				curLoc[0] -= int(l, 16)
			case "U":
				curLoc[1] += int(l, 16)
			case "D":
				curLoc[1] -= int(l, 16)
			
		vertices.append((curLoc[0], curLoc[1]))
		trenchLen += int(l, 16)

vertices.append(vertices[0])
s1 = 0
s2 = 0
for i in range(len(vertices)-1):
	s1 += vertices[i][0] * vertices[i+1][1]
	s2 += vertices[i][1] * vertices[i+1][0]

area = (0.5 * (abs(s1-s2) + trenchLen)) + 1
print(area)