starting = []
lines = []
with open("2023/24-input.txt", 'r') as textIn:
	for line in textIn:
		line = line.replace("\n", "")
		pos, vel = line.split(" @ ")
		pos = [int(i) for i in pos.split(", ")]
		vel = [int(i) for i in vel.split(", ")]
		starting.append((tuple(pos), tuple(vel)))

low = 200000000000000
high = 400000000000000

intersections = 0
for i in range(len(starting)):
    for j in range(i+1, len(starting)):
        p1x, p1y, _ = starting[i][0]
        p1mx, p1my, _ = starting[i][1]
        p2x, p2y, _ = starting[j][0]
        p2mx, p2my, _ = starting[j][1]
        p1m = p1my/p1mx
        p2m = p2my/p2mx
        if p1m == p2m:
            continue
        b1 = p1y - p1m*p1x
        b2 = p2y - p2m*p2x
        x = (b2-b1)/(p1m-p2m)
        y = p1m*x + b1
        if all((low <= x <= high,
                low <= y <= high,
                (x > p1x and p1mx > 0) or (x < p1x and p1mx < 0),
                (x > p2x and p2mx > 0) or (x < p2x and p2mx < 0)
		)):
            intersections += 1
print(intersections)