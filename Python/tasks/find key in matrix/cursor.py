matrix = [	
		[1,  4,  7,  11, 15],
		[2,  5,  8,  12, 19],
		[3,  6,  9,  16, 22],
		[10, 13, 14, 17, 24],
		[18, 21, 23, 26, 30]	
		 ]

k = 9
print("Key:", k)

flag = False

count = 0

for i in range(len(matrix[0])-1, -1, -1):
	count +=1
	if matrix[0][i] > k:
		print("col:", matrix[0][i])
		continue
	for j in range(len(matrix)):
		count +=1
		if matrix[j][i] > k:
			break
		print("row:", matrix[j][i])	
		if matrix[j][i] == k:
			print(matrix[j][i], True)
			flag = True
			break
	if flag:
		break
	print()
	

print("Количество операций:", count)

# Time: O(log2N)
# Memory: O(1)