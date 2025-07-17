arr = [7, 8, 9, 4, 6, 1, 0]


# Функция для поиска наименьшего числа в массиве
def findSmallest(arr):
    smallest = arr[0]
    smallest_index = 0

    for i in range(1, len(arr)):
        if smallest > arr[i]:
            smallest = arr[i]
            smallest_index = i
    return smallest_index


# Сортировка выборкой
def selectionSort(arr):
    new_arr = []
    for i in range(len(arr)):
        smallest = findSmallest(arr)
        new_arr.append(arr.pop(smallest))
        print(arr, " - ", new_arr)
    return new_arr


print("Сортировка выбором:")
print("Исходный массив:", arr)
print("Отсортированный массив:", selectionSort(arr))
