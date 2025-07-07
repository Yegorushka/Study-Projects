import re

# Исходный текст
with open('text.txt', "r", encoding="utf-8") as file:
    text = file.read().strip()

# Извлекаем все числа
numbers = re.findall(r'\d+', text)
unique_numbers = sorted(set(numbers), key=int)

# Словарь замен
replacements = {}

def replace_numbers(text, replacements):
    result = text
    for num, char in replacements.items():
        result = re.sub(rf'\b{num}\b', char, result)
    return result

print("Интерактивная замена чисел на буквы:")
print("Введите команду:")
print(" - любое число — начать замену")
print(" - 'show' — показать текущий результат")
print(" - 'list' — показать все числа")
print(" - 'save' — сохранить в result.txt")
print(" - 'exit' — выйти без сохранения\n")

while True:
    command = input(">>> ").strip().lower()

    if command == "exit":
        print("Выход без сохранения.")
        break

    elif command == "save":
        final_text = replace_numbers(text, replacements)
        with open("result.txt", "w", encoding="utf-8") as f:
            f.write(final_text)
        print("✅ Сохранено в result.txt")
        break

    elif command == "show":
        preview = replace_numbers(text, replacements)
        print("\n🔍 Текущий результат:")
        print(preview, "\n")

    elif command == "list":
        print("📋 Уникальные числа:")
        print(", ".join(unique_numbers))

    elif command.isdigit():
        if command not in unique_numbers:
            print("❌ Это число не найдено в тексте.")
            continue

        preview = replace_numbers(text, {**replacements, command: '?'})
        print("\nПример с '?':")
        print(preview)

        letter = input(f"Введите букву для числа {command}: ").strip().lower()
        if len(letter) == 1 and letter.isalpha():
            replacements[command] = letter
        else:
            print("⛔ Введите одну букву.\n")

    else:
        print("⛔ Неизвестная команда. Попробуй снова.\n")
