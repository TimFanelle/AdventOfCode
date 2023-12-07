'''
Code helped by:
https://github.com/xHyroM/aoc/blob/main/2023/7/first.py
Same solution, but different implementation because I was having some issues with the hand comparisons
'''
from collections import Counter
from enum import Enum

temp = ['A', 'K', 'Q', 'J', 'T', '9', '8', '7', '6', '5', '4', '3', '2']
temp.reverse()
strengths = dict()
for i in range(len(temp)):
	strengths[temp[i]] = i

class CardCombos(Enum):
	FIVE = 1
	FOUR = 2
	FULL = 3
	THREE = 4
	TWO = 5
	ONE = 6
	HIGH = 7

	def __str__(self) -> str:
		return self.name

class Hand():
	def __init__(self, cards, bid):
		self.cards = cards
		self.bid = bid
		self.combo = self.detectCombo(cards)
		self.list_card = list(cards)

	def detectCombo(self, cards_in):
		count = Counter(cards_in)
		card_occurances = [count for _, count in count.most_common()]

		match card_occurances:
			case [5]:
				return CardCombos.FIVE
			case [4, 1]:
				return CardCombos.FOUR
			case [3, 2]:
				return CardCombos.FULL
			case [3, 1, 1]:
				return CardCombos.THREE
			case [2, 2, 1]:
				return CardCombos.TWO
			case [2, 1, 1, 1]:
				return CardCombos.ONE
			case _:
				return CardCombos.HIGH
	
	def __lt__(self, other):
		if self.combo != other.combo:
			return self.combo.value < other.combo.value

		for i in range(5):
			if self.list_card[i] != other.list_card[i]:
				return strengths[self.list_card[i]] > strengths[other.list_card[i]]
hands = []
with open("AdventOfCode/2023/7-input.txt", 'r') as textIn:
	for line in textIn:
		cards, bid = line.split()
		bid = int(bid)
		hands.append(Hand(cards, bid))

hands.sort(reverse=True)
total = 0
for item in range(len(hands)):
	total += hands[item].bid * (item+1)
print(total)