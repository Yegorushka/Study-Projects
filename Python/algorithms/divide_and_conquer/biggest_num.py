def biggest_num(arr):
    if len(arr) == 2:
        return arr[0] if arr[0] > arr[1] else arr[1]
    sub_max = biggest_num(arr[1:])
    return arr[0] if arr[0] > sub_max else sub_max


print(biggest_num([1, 3, 5, 8]))
