matrix = [	
		[1,  4,  7,  11, 15],
		[2,  5,  8,  12, 19],
		[3,  6,  9,  16, 22],
		[10, 13, 14, 17, 24],
		[18, 21, 23, 26, 30]	
		 ]

k = 9
print("Key:", k)
my_set = set([item for row in matrix for item in row])
print(k, k in my_set)

# Time: O(1)
# Memory: O(n)