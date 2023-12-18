'''
See:
Shoelace method
Pick's theorem
'''
curLoc = [0, 0]
vertices = []
trenchLen = 0
with open("AdventOfCode/2023/18-input.txt", 'r') as textIn:
	for line in textIn:
		dir, l, hex = line.split()
		hex = hex[2:-1]
		match dir:
			case "R":
				curLoc[0] += int(l)
			case "L":
				curLoc[0] -= int(l)
			case "U":
				curLoc[1] += int(l)
			case "D":
				curLoc[1] -= int(l)
			
		vertices.append((curLoc[0], curLoc[1], int(hex, 16)))
		trenchLen += int(l)

vertices.append(vertices[0])
s1 = 0 # vertices[-1][0] * vertices[0][1]
s2 = 0 # vertices[-1][1] * vertices[0][0]
for i in range(len(vertices)-1):
	s1 += vertices[i][0] * vertices[i+1][1]
	s2 += vertices[i][1] * vertices[i+1][0]

area = (0.5 * (abs(s1-s2) + trenchLen)) + 1
print(area)