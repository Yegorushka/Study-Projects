arr = [-2, -1, 1, 2, 5, 9, 18, 25, 26, 30, 34, 52, 78, 103, 154, 204, 305, 555, 789]
k = 64

print("arr = ", arr)
print("k = ", k)

n, p = -1, 0

count = 0

for i in range(len(arr)):
	print(arr[p], "+", arr[n],"=",arr[p] + arr[n]) 
	if arr[p] + arr[n] == k:
		print("!!! -->", arr[p], arr[n]) 
		break
	elif arr[p] + arr[n] > k: 
		n -=1
	else:
		p +=1
	
	count +=1

print(f"Количество операций {count}")

# Time: O(n)
# Memory: O(1)