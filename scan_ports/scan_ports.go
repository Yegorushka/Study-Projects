package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main(){
    var arr []string

    var IPs []string
    var startPorts []string
    var endPorts []string 


    // Чтение файла
    file, err := os.Open("diaposon.txt")
    if err != nil {
        fmt.Println("Ошибка при чтении:", err)
        return
    }
    defer file.Close()


    // Сканирует файл по строчно и записует в срез
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        arr = append(arr, line)
    }
    if err := scanner.Err(); err != nil {
        fmt.Println("Ошибка при сканирование:", err)
    }


    // Деление строки на айпи адрес и порты
    for i, _ := range arr {
        parts := strings.Split(arr[i], " ")
        if len(parts) != 3 {
            fmt.Println("Неверный формат строки")
            return
        }

        IPs = append(IPs, parts[0])
        startPorts = append(startPorts, parts[1])
        endPorts = append(endPorts, parts[2])
    }

    fmt.Println("IP адреса:", IPs)
    fmt.Println("Начальный порт:", startPorts)
    fmt.Println("Последний порт:", endPorts)

}

    // Вывод среза по строчно
    // for i, v := range arr {
    //     if i > 0 {
    //         fmt.Println()
    //     }
    //     fmt.Print(v)
    // }


// Асинхронный веб-сканер портов
// Задача:
// Напиши программу, которая принимает список хостов и диапазон портов,
// проверяет их доступность параллельно (goroutines), и выдаёт JSON-отчёт.
// Условия:
//     Максимальное количество одновременных проверок (воркеров) задаётся параметром.
//     Время таймаута подключения – настраиваемое.
//     В отчёте указать статус порта (открыт/закрыт), время отклика.