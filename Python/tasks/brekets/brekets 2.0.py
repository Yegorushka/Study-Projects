brekets = "[[({(}))]]"
hash_t = {}

hash_t["{"] = "00"
hash_t["["] = "01"
hash_t["("] = "02"
hash_t["}"] = "10"
hash_t["]"] = "11"
hash_t[")"] = "12"

# Моя логика заключатеся в том, что я использую
# скобки в хт для получения АйДи скобки, т.е.
# первая цифра в значение означает, что скобка открываеся,
# вторая цифра, ее номер

flag = True
print(brekets)

stack = []
for i in range(len(brekets)):
	if hash_t[brekets[i]][0] == "0":
		print("append", brekets[i], hash_t[brekets[i]]) # "Добавляем в стэк левую часть скобки"
		stack.append(hash_t[brekets[i]])
	else:
		pop = stack.pop()
		print("pop", pop, hash_t[brekets[i]], brekets[i]) # "Делаем поп из стека"
		if pop[1] != hash_t[brekets[i]][1]:
			flag = False
	
if flag == True:
	print("OK")
else:
	print("No OK")