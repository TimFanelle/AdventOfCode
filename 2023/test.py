class updatedComplex(complex):
	def __lt__(self, other):
		if self.real < other.real:
			return True
		elif self.real > other.real:
			return False
		else:
			if self.imag < other.imag:
				return True
			elif self.imag > other.imag:
				return False
			else:
				return False
	def __le__(self, other):
		if self.real < other.real:
			return True
		elif self.real > other.real:
			return False
		else:
			if self.imag < other.imag:
				return True
			elif self.imag > other.imag:
				return False
			else:
				return True
