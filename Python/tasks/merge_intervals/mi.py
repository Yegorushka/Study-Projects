arr = [[1,3], [2,6], [8,10], [15,18]]

extra_arr = [] # Создаем копию массив, но уже с заполненным диапозоном

for i in range(len(arr)):
	extra_arr.append(list(range(arr[i][0], arr[i][1]+1)))	# Заполняем диапозон

intersection = list(set(extra_arr[0]) | set(extra_arr[1]))

print(intersection)