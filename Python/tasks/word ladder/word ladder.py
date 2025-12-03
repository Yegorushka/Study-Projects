from collections import deque

words = [
    "hit",
    "cog",
    "dot",
    "fig",
    "dog",
    "hot",
    "mag",
    "dag",
    "toh",
    "hat",
    "het"  #11
]

print("List:", end=" ")
for i in words:
    print(i, end=" ")

start = "hit"
goal = "cog"
print(f"\nStart word: {start}. Goal word: {goal}")

# Приобразуем список в словарь с связями
graph = {}
for i in range(len(words)):
    for j in range(len(words)): 
        matches = sum(c1 == c2 for c1, c2 in zip(words[i], words[j]))
        if matches == len(words[0])-1:
            # print(words[i], words[j])
            if words[i] not in graph:
                graph[words[i]] = []
            graph[words[i]].append(words[j])

# Вывод словаря
# for key, value in graph.items(): 
#     print(f"{key}: {value}")


# Используем алгоритм поиск в ширину
# Очередь хранит вершину и путь до неё
queue = deque()
queue.append((start, [start]))

visited = set()

while queue:
    word, path = queue.popleft()
    if word == goal:
        print(f"Кратчайший путь: {path}.\nДлина пути: {len(path)-1}")
        break
    visited.add(word)
    for neighbor in graph.get(word, []):
        if neighbor not in visited:
            queue.append((neighbor, path + [neighbor]))
            visited.add(neighbor)


# Задача:
# Даны:
#     Начальное слово.
#     Конечное слово.
#     Список слов.
# За одну операцию можно менять одну букву. 
# Нужно найти минимальное количество шагов, чтобы 
# превратить начальное слово в конечное, все промежуточные 
# слова должны быть в словаре.

# Пример:
# hit -> hot -> dot -> dog -> cog