def sum_rec(arr: list):
    if arr == []:
        return 0
    return arr[0] + sum(arr[1:])


def sum_loop(arr):
    sum = 0
    for i in range(len(arr)):
        sum += arr[i]
    return sum


print(sum_loop([1, 3, 5, 8]))
print(sum_rec([1, 3, 5, 8]))
