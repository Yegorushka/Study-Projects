package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "strconv"
    "net"
    "time"
)

func main(){
    var arr []string

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


    var IPs []string
    var startPorts []int
    var endPorts []int 

    // Деление строки на айпи адрес и порты
    for i, _ := range arr {
        parts := strings.Split(arr[i], " ")
        if len(parts) != 3 {
            fmt.Println("Неверный формат строки")
            return
        }

        IPs = append(IPs, parts[0])
        p1, err := strconv.Atoi(parts[1])
        startPorts = append(startPorts, p1)
        p2, err := strconv.Atoi(parts[2])
        endPorts = append(endPorts, p2)

        if err != nil {
            fmt.Println("Ошибка:", err)
            return
        }
    }

    // Задержка в мс
    var num int
    fmt.Print("Введите время задержки в миллисекундах (оптимально 500):")
    _, err = fmt.Scanf("%d\n", &num)
    if err != nil {
        fmt.Println("Ошибка:", err)
        return
    }
    timeout := time.Duration(num) * time.Millisecond

    // Канал для результатов
    var results []chan int

    var timings []time.Duration

    for i, _ := range IPs {
        ch := make(chan int, endPorts[i]-startPorts[i]+1)
        results = append(results, ch)

        for port := startPorts[i]; port <= endPorts[i]; port++ {
            go func(p int) {
                address := IPs[i] + ":" + strconv.Itoa(p)

                start_time := time.Now()
                conn, err := net.DialTimeout("tcp", address, timeout)
                elapsed_time := time.Since(start_time)
                timings = append(timings, elapsed_time)

                if err != nil {
                    results[i] <- 0
                    return
                }
                conn.Close()
                results[i] <- p
            }(port)
        }

        // Собираем результаты
        fmt.Println("IP адрес:", IPs[i], "диапозон сканирования:", startPorts[i], "-", endPorts[i])
        for j := startPorts[i]; j <= endPorts[i]; j++ {
            port := <-results[i]
            if port != 0 {
                fmt.Println("Open port:", port, "| Time:", timings[j])
            }
        }
        fmt.Println()   
    }

}



// Асинхронный веб-сканер портов
// Задача:
// Напиши программу, которая принимает список хостов и диапазон портов,
// проверяет их доступность параллельно (goroutines), и выдаёт JSON-отчёт.
// Условия:
//     Максимальное количество одновременных проверок (воркеров) задаётся параметром.
//     Время таймаута подключения – настраиваемое.
//     В отчёте указать статус порта (открыт/закрыт), время отклика.