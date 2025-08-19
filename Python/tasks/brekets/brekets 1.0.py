brekets = "{[}]"
print(brekets)

mid = int(len(brekets) / 2)

left = brekets[:mid]
right = brekets[mid:]

flag = True

for i in range(mid):
	if right[-1-i] == "}":
		r = "{"
	elif right[-1-i] == "]":
		r = "["
	elif right[-1-i] == ")":
		r = "("
	
	print(left[i], " ", right[-1-i])

	if left[i] != r:
		flag = False

if flag == True:
	print("OK")
else:
	print("No OK")


# Скобки (Valid Parentheses)

# Задача:
# Проверить строку, состоящую из скобок ()[]{}, на корректность вложенности.
# Например:

# {[()]} → OK

# {[(])} → Не OK

# Ограничение:
# Алгоритм за O(n) по времени и O(n) по памяти.