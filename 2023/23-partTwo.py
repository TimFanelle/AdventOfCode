#solution from https://topaz.github.io/paste/#XQAAAQAaCAAAAAAAAAAzHIoib6pXbueH4X9F244lVRDcOZab5q1+VXY/ex42qR7D+RY7eLybVRP3dw6C3Q/38Mc2/ci4Da/IlFot/yvagQO/q3bMwccm4C5+AC63A022tLH8GA/FZX5R4iy/vyld7TDwd6QwIh44JeEVGX3j7oq4iqf4FO0YBWGU3g49zXrywfdDoKYmwREm+yA081RC44zg+66MMx6BbSXXwpBme/7lGUsaR8bgsCTh/uM47JoX+NdhaOZr15XTlcNoZZn3zt4x8by7siEbEk0DYE7/UA/lbxVd+fEidFGBO5zH2eCPZ1ep6gx80kj9Vb8YNGFeEi8AllUXTNmN1ed70EM0P6Ylyr3D03jXcKjfWDL8M6vK9j19hX4/On66RQD+Q6jXywHqaEqKHzhRVXgHrMUL9EMO8aBm/NeHp9wHAg2j0Xr1239vCglQk7vZjJQos33//NnrhckVj4lx8Os4zzArRx9bd+/3Og0M1/siP97h9qQMd7fJYL6CE0GxPFOCMMCrv8YgKBUP8KaA02UVSm2JNOmOSURb1CsbnmKSr6QznaCzucq69DGtQEdlBBWdP6v6bLACcFHWuGDDIdsSfa/++fS8+yjSKXLZFZa9jyc/Ay2GBD0GjpIwezONMYmZkwham1FH1P7IQ/vBK3Mget4Za4S4O5K9B9ymXr5mDF23g+nuC70AVsFLAjYlnMjkgBeRho6Nl0ZuOcTWlgNh6kT52zYzRWcTt+ODBv+oHr53/qS6y/TVZMhr1vGSrfUrWK40tndTd5MmSb0+YDLfB3+eot6wlV2zNd5b2pT5bWpJrFtQM28aq/WlVXUZz5Mk2w1hXjVNeexAq7Q75m3ot3SMyleHgzjQ5dRuHfbQjo3oyE8jNPdbwgnfs5ZXtORlhmr/FOQzoLzYV5besKji7EFQPdFInO/h+PUhz7h669TXlOK15e8R3K4Tl8n45tV0+w0roIXQbHn48LP/fOtzAA==
from collections import defaultdict

NDIRS = {'>':1, '^': -1j, '<': -1, 'v':1j}
n4 = lambda p: [p - 1j, p - 1, p + 1, p + 1j] # coordinates of 4 neighbors

grid = []
with open("23-input.txt", 'r') as textIn:
	for line in textIn:
		temp = []
		for c in line:
			if c != "\n":
				temp.append(c)
		grid.append(temp)

def find_adjacent(start, grid, terminals, ndirs):
    adj, q = [], [(start,0,{start})]
    while q:
        p, l, seen = q.pop(0)
        if p in terminals and p != start:
            adj.append((p,l))
            continue

        neighbors = [n for n in n4(p) if n in grid and n not in seen and grid[n] != '#']
        if len(neighbors) > 1 and p != start:
            adj.append((p, l))
            continue

        for n in neighbors:
            if ndirs and grid[n] in NDIRS and n + NDIRS[grid[n]] != p:
                q.append((n + NDIRS[grid[n]], l+2, seen | {n, n+NDIRS[grid[n]]}))
            elif grid[n] == '.' or not ndirs:
                q.append((n, l+1, seen | {n}))
    return adj

def build_graph(grid, start, end, ndirs):
    graph, seen, q = defaultdict(list), set(), [start]
    while q:
        p = q.pop()
        if p in seen: 
            continue
        seen.add(p)
        for n, l in find_adjacent(p, grid, [start, end], ndirs):
            graph[p].append((n, l))
            if n not in seen:
                q.append(n)
    return graph

def longest_path(graph, start, end):
    longest, q = 0, [(start, 0, {start})]
    while q:
        p, l, seen = q.pop()
        if p == end:
            longest = max(longest, l)
            continue
        for n,nl in graph[p]:
            if n not in seen:
                q.append((n, l+nl, seen | {n}))
    return longest

def part2(grid, start, end):
    graph = build_graph(grid, start, end, False)
    print(longest_path(graph, start, end))

def parse(fn):
    grid = {x+1j*y:c for y,l in enumerate(open(fn).readlines()) for x,c in enumerate(l.strip())}
    start, end = 1, max(p.real for p in grid)- 1+ 1j*max(p.imag for p in grid)
    return grid, start, end

grid, s, e = parse("23-input.txt")
part2(grid, s, e)