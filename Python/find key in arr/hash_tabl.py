arr = [-2, -1, 1, 2, 5, 9, 18, 25, 26, 30, 34, 52, 78, 103, 154, 204, 305, 555, 789]
k = 64

print("arr = ", arr)
print("k = ", k)

count = 0

my_set = set(arr)

for i in range(len(arr)):
	count +=1

	search_num = k - arr[i]
	print(f"{k}-{arr[i]}={search_num}")

	if search_num in my_set:
		print("!!! -->", arr[i], search_num)
		break

print(f"Количество операций {count}")

# Time: O(n)
# Memory: O(n)