# Тут код уже не мой, идею подкинул гпт

def is_valid(s: str) -> bool:
	stack = []
	pairs = {')': '(', ']': '[', '}': '{'}

	for ch in s:
		if ch in pairs.values(): # открывающая
			stack.append(ch)
			# print(ch)
		elif ch in pairs: # закрывающая
			if not stack or stack[-1] != pairs[ch]:
				return False
			stack.pop()
		else:
			return False

	# print(stack, not not stack)
	return not stack

print(is_valid("[[({(}))]]"))  # False
print(is_valid("{[()]}"))      # True
print(is_valid("((("))         # False
print(is_valid("([])"))        # True