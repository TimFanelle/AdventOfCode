'''
Code from:
https://topaz.github.io/paste/#XQAAAQDBAgAAAAAAAAA2mEu65RRUNtdenXdrhhhWtpWChxZ5NlRUV0cHAV4QddbbMOGhwGUfTEF213+cfOv3if1vqC12H3mwM+pCpva0Vs615ml6/f3ax2WWiewKIK2QAo6SdVxywBlzF7dYPRWBFXaGag0SzmAor5eVq/rlwNOk0DZeGnWATv5I7GybGQUkBCcy+WODONAuj4NMuO3yWZDlYkWUOh5HAYVVQQRwshPo4UOCi+mY84fPlmQHYyajJWpF9IRweU5v8STiTi+e9L2Rbo5T96HTgT/gVjKTdI5XfHdtoepngO1BCm5/4UBmL9sOjcfFfzA2Y8WXBlhu/ZejbrUboebB2O5LEDjrYuFyguZNaZgX5jBS0xkf33X96ASnYeD7t9FwjHuZp+RfcEHcDosdQ3ukhQB0FQv3to4vwUQuXr2cj0Jzp0OTHmHwIfefG2TxdjoxinEF4+oo/pby58dhRNuUqCkvGnb8dHffvuzha7yFJzDqHjY8Pw9ePMJiPfB9hsb6aehYHPqLNIaUgrOt7hnu9ZydOH0NgQsSyun12eP7df3Nxb0=
'''
maze = {complex(i,j): c for i,r in enumerate(open('AdventOfCode/2023/10-input.txt'))
                        for j,c in enumerate(r.strip())}

N, S, E, W = -1, +1, +1j, -1j
dirs = {'|': (N, S), '-': (E, W), 'L': (N, E),
        'J': (N, W), '7': (S, W), 'F': (S, E),
        'S': (N, E, S, W), '.':()}

graph = {p: {p+d for d in dirs[c]} for p,c in maze.items()}
start = [p for p,d in graph.items() if len(d) == 4][0]

seen = {start}
while todo := graph[start]:
    node = todo.pop()
    seen |= {node}
    todo |= graph[node]-seen

irange = lambda n: [complex(n.real, i) for i in range(int(n.imag))]

print(len(seen)//2,
      sum(sum(maze[m] in "|JLS" and m in seen for m in irange(p)) % 2
          for p in set(maze)-seen))