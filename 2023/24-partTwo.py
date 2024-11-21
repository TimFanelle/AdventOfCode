# explained by https://www.reddit.com/r/adventofcode/comments/18pnycy/comment/khlrstp/?utm_source=share&utm_medium=web3x&utm_name=web3xcss&utm_term=1&utm_content=share_button

import numpy as np

starting = []
lines = []
with open("C:\\Users\\quest\\Stuff_I_guess\\AdventOfCode\\2023\\24-input.txt", 'r') as textIn:
	for line in textIn:
		line = line.replace("\n", "")
		pos, vel = line.split(" @ ")
		pos = [int(i) for i in pos.split(", ")]
		vel = [int(i) for i in vel.split(", ")]
		starting.append((tuple(pos), tuple(vel)))

def add(hails, n):
	A = []
	B = []
	item_0 = [hails[0][0][0], hails[0][0][1],hails[0][0][2],hails[0][1][0],hails[0][1][1],hails[0][1][2]]
	item_n = [hails[n][0][0], hails[n][0][1],hails[n][0][2],hails[n][1][0],hails[n][1][1],hails[n][1][2]]
	
	A.append([item_0[4]-item_n[4], item_n[3]-item_0[3], 0*n, item_n[1]-item_0[1], item_0[0]-item_n[0], 0*n])
	B.append(item_0[0]*item_0[4]-item_0[1]*item_0[3]-item_n[0]*item_n[4]+item_n[1]*item_n[3])
	A.append([item_0[5]-item_n[5], 0*n, item_n[3]-item_0[3], item_n[2]-item_0[2], 0*n, item_0[0]-item_n[0]])
	B.append(item_0[0]*item_0[5]-item_0[2]*item_0[3]-item_n[0]*item_n[5]+item_n[2]*item_n[3])
	return A, B
	
def determinant(matrix): 
	""" Calculate the determinant of a matrix recursively using the Decimal class. Parameters: matrix (list of list of Decimal): Matrix to find the determinant of. Returns: Decimal: Determinant of the matrix. """ 
	if len(matrix) == 1: 
		return matrix[0][0] 
	det = Decimal(0) 
	for c in range(len(matrix)): 
		sub_matrix = [row[:c] + row[c+1:] for row in matrix[1:]] 
		det += ((-1) ** c) * matrix[0][c] * determinant(sub_matrix) 
	return det

def cramer(A, B):
	A = np.array(A, dtype=np.float16)
	B = np.array(B, dtype=np.float16)
	detA = np.linalg.det(A)
	n = len(B) 
	X = np.zeros(n) 
	for i in range(n): 
		A_i = np.copy(A) 
		A_i[:, i] = B 
		X[i] = np.linalg.det(A_i) / detA 
	return X

from decimal import Decimal, getcontext 
def cramers_rule(A, B, precision=50):
	""" Solve the system of linear equations AX = B using Cramer's rule with high precision. Parameters: A (list of list of Decimal): Coefficient matrix. B (list of Decimal): Constant terms vector. precision (int): Number of decimal places for precision. Returns: list of Decimal: Solution vector. """ 
	# Set the precision 
	getcontext().prec = precision 
	det_A = determinant(A) 
	if det_A == 0: 
		raise ValueError("The coefficient matrix A is singular, and the system does not have a unique solution.") 
	n = len(B) 
	X = [Decimal(0) for _ in range(n)] 
	for i in range(n): 
		A_i = [row[:] for row in A] # Make a deep copy of A 
		for row in A_i: 
			row[i] = B[A_i.index(row)] 
		det_A_i = determinant(A_i) 
		X[i] = det_A_i / det_A 
	return X

A, B = [], []
for i in range(1, 4):
	c, d = add(starting, i)
	A += c
	B += d

C = []
for line in A:
	D = []
	for item in line:
		D.append(Decimal(item))
	C.append(D)
A = C

#[pxr, pyr, pzr, _, _, _] = cramer(A, B)
[pxr, pyr, pzr, _, _, _] = cramers_rule(A, B)
print(np.sum([pxr, pyr, pzr]))

print(pxr + pyr + pzr)