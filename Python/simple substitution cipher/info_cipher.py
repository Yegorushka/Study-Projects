from collections import Counter
import re

with open('text.txt', "r", encoding="utf-8") as file:
    text = file.read().strip()

russian_letters = [	# Частотность букв русского алфавита
    "о", "е", "а", "и", "н", "т", "с", "р", "в", "л", "к", "м", "д", "п", "у",
    "я", "ы", "ь", "г", "з", "б", "ч", "й", "х", "ж", "ш", "ю", "ц", "щ", "э", "ф", "ъ", "ё"
]

numbers = list(map(int, re.findall(r"\d+", text)))                              # Извлекаем все числа
counter = Counter(numbers)                                                      # Подсчитываем количество каждого числа
sorted_counts = sorted(counter.items(), key=lambda x: x[1], reverse=True)       # Сортируем по убыванию количества
total  = len(numbers)                                                           # Общее количество чисел

print("Сумма всех встречаемых знаков (чисел):",total )

print("\nЧисло  -> Псевдо-буква -> % Частоты")
for i, (num, count) in enumerate(sorted_counts):
    percent = (count / total) * 100
    letter = russian_letters[i] if i < len(russian_letters) else "?"
    print(f"{num:>5}  ->   {letter}   ->  {percent:5.2f}%")
