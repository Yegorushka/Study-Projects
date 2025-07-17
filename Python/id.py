import sys
import array

num = 15
print(f"type: {type(num)}, {num}, id {id(num)}, size {sys.getsizeof(num)}\n")

num_d = 3.15
print(f"type: {type(num_d)}, {num_d}, id {id(num_d)}, size {sys.getsizeof(num_d)}\n")

string = "string"
print(f"type: {type(string)}, {string}, id {id(string)}, size {sys.getsizeof(string)}\n")

list_ = [0, 1, 2, 3, 4]
print(f"type: {type(list_)}, {list_}, id {id(list_)}, size {sys.getsizeof(list_)}")
for i in range(len(list_)):
	print(f"type: {type(list_[i])}, {list_[i]}, id {id(list_[i])}, size {sys.getsizeof(list_[i])}")

list_ = [0, 1, True, "string", 's', 3.15, [1, 2, 3]]
print(f"\ntype: {type(list_)}, {list_}, id {id(list_)}, size {sys.getsizeof(list_)}")
for i in range(len(list_)):
	print(f"type: {type(list_[i])}, {list_[i]}, id {id(list_[i])}, size {sys.getsizeof(list_[i])}")

my_array = array.array('i', [1, 2, 3, 4])
print(f"\ntype: {type(my_array)}, {my_array}, id {id(my_array)}, size {sys.getsizeof(my_array)}")
for i in range(len(my_array)):
	print(f"type: {type(my_array[i])}, {my_array[i]}, id {id(my_array[i])}, size {sys.getsizeof(my_array[i])}")