arr = [-2, -1, 1, 2, 5, 9, 18, 25, 26, 30, 34, 52, 78, 103, 154, 204, 305, 555, 789]
k = 64

print("arr = ", arr)
print("k = ", k)

count = 0

break_out_flag = False

for i in range(len(arr)):

	for j in range(i+1 , len(arr)):
		count +=1

		if arr[i] + arr[j] == k:
			print("[!]", arr[i], "+", arr[j], "=", arr[i]+arr[j])
			break_out_flag = True
			break
		else:
			print(arr[i], "+", arr[j], "=", arr[i]+arr[j])
	
	if break_out_flag:
		break


print(f"Количество операций {count}")

# Time: O(n^2)
# Memory: O(1)