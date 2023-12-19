'''
from copy import deepcopy
workflows = dict()
items = []
shifted = False
with open("AdventOfCode/2023/19-input.txt") as textIn:
	for line in textIn:
		if line == '\n':
			break
		else:
			temp_line = line.split("{")
			rules = temp_line[1].split("}")
			rules = rules[0].split(",")
			workflows[temp_line[0]] = rules

v = {'x':(1,4000), 'm':(1,4000), 'a':(1,4000), 's':(1,4000)}

def detRejAcc(work_name, values):
	curRules = workflows[work_name]
	accept = 0
	passVals = deepcopy(values)
	for rule in curRules:
		if ":" in rule:
			r, out = rule.split(":")
			newVals = deepcopy(passVals)
			if ">" in r:
				r_updated = r.split(">")
				if newVals[r_updated[0]][1] > int(r_updated[1]):
					newVals[r_updated[0]] = (int(r_updated[1])+1, newVals[r_updated[0]][1])
					if not(out in ["R", "A"]):
						acc = detRejAcc(out, newVals)
						accept += acc
					else:
						if out == "R":
							return 0
							#passVals[r_updated[0]] = (newVals[r_updated[0]][0], int(r_updated[1])-1)
						else:
							accept_add = 1
							no_add = False
							for k in values.keys():
								if passVals[k][1] >= passVals[k][0]:
									accept_add *= passVals[k][1]-passVals[k][0] + 1
								else:
									no_add = True
							accept += accept_add*(not(no_add))
							return accept
							#passVals[r_updated[0]] = (newVals[r_updated[0]][0], int(r_updated[1])-1)
				else:
					passVals[r_updated[0]] = (newVals[r_updated[0]][0], min(int(r_updated[1])-1, newVals[r_updated[0]][1]))
			else:
				r_updated = r.split("<")
				if newVals[r_updated[0]][0] < int(r_updated[1]):
					newVals[r_updated[0]] = (newVals[r_updated[0]][0], int(r_updated[1])-1)
					if not(out in ["R", "A"]):
						acc = detRejAcc(out, newVals)
						accept += acc
					else:
						if out == "R":
							#passVals[r_updated[0]] = (newVals[r_updated[0]][0], int(r_updated[1])-1)
							return 0
						else:
							accept_add = 1
							no_add = False
							for k in values.keys():
								if passVals[k][1] >= passVals[k][0]:
									accept_add *= passVals[k][1]-passVals[k][0] + 1
								else:
									no_add = True
							accept += accept_add*(not(no_add))
							return accept
							#passVals[r_updated[0]] = (newVals[r_updated[0]][0], int(r_updated[1])-1)
				else:
					passVals[r_updated[0]] = (max(int(r_updated[1])+1, newVals[r_updated[0]][0]), newVals[r_updated[0]][1])

		else:
			if rule == "A":
				accept_add = 1
				no_add = False
				for k in values.keys():
					if passVals[k][1] >= passVals[k][0]:
						accept_add *= passVals[k][1]-passVals[k][0] + 1
					else:
						no_add = True
				accept += accept_add*(not(no_add))
				return accept
			elif rule == "R":
				return 0
			else:
				acc = detRejAcc(rule, passVals)
				accept += acc
	return accept

j = detRejAcc("in", v)
q = 167409079868000
print(j)
print("Off by:", q-j)
'''

from time import time
import re
from math import prod


def timer_func(func):
	# This function shows the execution time of
	# the function object passed
	def wrap_func(*args, **kwargs):
		t1 = time()
		result = func(*args, **kwargs)
		t2 = time()
		print(f'Function {func.__name__!r} executed in {(t2 - t1):.4f}s')
		return result

	return wrap_func


def rate_part(part, workflows, wfn):
	if wfn in 'AR':
		return wfn
	cwf = workflows[wfn]
	for step in cwf:
		# if the step is a string, it is either accepted, rejected, or sent to another workflow
		if isinstance(step, str):
			# if the step is either A or R
			if step in 'AR':
				# return the step
				return step
			else:
				# else, run it through the workflow it says to
				return rate_part(part, workflows, step)
		# otherwise we have to do the evaluation
		else:
			if step['op'] == '<':
				if part[step['cat']] < step['val']:
					return rate_part(part, workflows, step['dst'])
			else:
				if part[step['cat']] > step['val']:
					return rate_part(part, workflows, step['dst'])


def less_than_range(r, v):
	l, u = r
	if l < v <= u:
		return (l, v-1), (v, u)
	elif v <= l:
		return None, (l, u)
	elif v > u:
		return (l, u), None


def greater_than_range(r, v):
	l, u = r
	if l <= v < u:
		return (l, v), (v + 1, u)
	elif v < l:
		return None, (l, u)
	elif v >= u:
		return (l, u), None


def rate_part_range(part, workflows, wfn):
	if wfn == 'A':
		return prod([v[1] - v[0] + 1 for c, v in part.items() if c in 'xmas'])
	elif wfn == 'R':
		return 0
	combos = 0
	cwf = workflows[wfn]
	for step in cwf:
		# if the step is a string, it is either accepted, rejected, or sent to another workflow
		if isinstance(step, str):
			# if the step is either A or R
			if step in 'A':
				# add the product of the ranges to the combos
				combos += prod([v[1] - v[0] + 1 for c, v in part.items() if c in 'xmas'])
			else:
				# else, run it through the workflow it says to
				combos += rate_part_range(part, workflows, step)
		# otherwise we have to do the evaluation
		else:
			if step['op'] == '<':
				# split the range into the less than and greater than
				l, u = less_than_range(part[step['cat']], step['val'])
				# lower than range gets sent to the next workflow
				if l:
					new_part = part.copy()
					new_part[step['cat']] = l
					combos += rate_part_range(new_part, workflows, step['dst'])
				# upper range remains to go through the rest of this workflow
				if u:
					part[step['cat']] = u
			else:
				# split the range
				l, u = greater_than_range(part[step['cat']], step['val'])
				# lower range remains to go through the rest of this workflow
				if l:
					part[step['cat']] = l
				# upper range goes to the next workflow
				if u:
					new_part = part.copy()
					new_part[step['cat']] = u
					combos += rate_part_range(new_part, workflows, step['dst'])

	return combos

wfl = dict()
items = []
shifted = False
with open("AdventOfCode/2023/19-input.txt") as textIn:
	for line in textIn:
		if line == '\n':
			break
		else:
			temp_line = line.split("{")
			rules = temp_line[1].split("}")
			rules = rules[0].split(",")
			wfl[temp_line[0]] = []
			for r in rules:
				if ":" in r:
					cat = r[0]
					op = r[1]
					val = r[2:r.index(":")]
					dst = r[r.index(":")+1:]
					wfl[temp_line[0]].append({'cat': cat, 'op': op, 'val': int(val), 'dst': dst})
				else:
					wfl[temp_line[0]].append(r)

part = {i: (1, 4000) for i in 'xmas'}
print(rate_part_range(part, wfl, 'in'))

'''
Solution pulled from:
https://pastebin.com/ALpfax5Q
created by: https://www.reddit.com/user/illuminati229/
'''