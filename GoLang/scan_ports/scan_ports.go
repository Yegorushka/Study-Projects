package main

import (
    "bufio"
    "fmt"
    "flag"  // Если воркеры не работают, коменти флаг
    "os"
    "strings"
    "strconv"
    "net"
    "time"
)


func worker(id int, jobs <-chan int, ip string, results chan<- int, timings chan<- time.Duration, timeout time.Duration){
    for port := range jobs {
        address := ip + ":" + strconv.Itoa(port)

        start_time := time.Now()
        conn, err := net.DialTimeout("tcp", address, timeout)
        elapsed_time := time.Since(start_time)

        timings <- elapsed_time

        if err != nil{
            results <- 0
        } else {
            conn.Close()
            results <- port
        }
    }
}


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


    // Деление строки на айпи адрес и порты
    var IPs []string
    var startPorts []int
    var endPorts []int 
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


    // Задаем количество горутин
    // Этот код можно заменить кодом ниже без установки количества воркеров
    workerCount := flag.Int("workers", 1000, "количество воркеров")
    flag.Parse()

    fmt.Println("Количество горутин:", *workerCount)


    for i, ip := range IPs {
        jobChan := make(chan int, 100)
        resultChan := make(chan int, 100)
        timingChan := make(chan time.Duration, 100)

        // Запускаем воркеров
        for w := 0; w < *workerCount; w++{
            go worker(w, jobChan, ip, resultChan, timingChan, timeout)
        }

        // Кидаем задачи
        for port := startPorts[i]; port <= endPorts[i]; port++{
            jobChan <- port
        }
        close(jobChan)

        // Считываем результат
        fmt.Println("IP адрес:", ip, "диапозон сканирования:", startPorts[i], "-", endPorts[i])
        for j := startPorts[i]; j <= endPorts[i]; j++ {
            port := <- resultChan
            elapsed := <- timingChan

            if port != 0 {
                fmt.Println("Open port:", port, "| Time:", elapsed)
            }
        }

    }
    // Конец 


    // // Канал для результатов
    // var results []chan int
    // var timings []time.Duration
    // for i, ip := range IPs {
    //     ch := make(chan int, endPorts[i]-startPorts[i]+1)
    //     results = append(results, ch)

    //     for port := startPorts[i]; port <= endPorts[i]; port++ {
    //         go func(p int) {
    //             address := ip + ":" + strconv.Itoa(p)

    //             start_time := time.Now()
    //             conn, err := net.DialTimeout("tcp", address, timeout)
    //             elapsed_time := time.Since(start_time)
    //             timings = append(timings, elapsed_time)

    //             if err != nil {
    //                 results[i] <- 0
    //                 return
    //             }
    //             conn.Close()
    //             results[i] <- p
    //         }(port)
    //     }


    //     // Собираем результаты
    //     fmt.Println("IP адрес:", ip, "диапозон сканирования:", startPorts[i], "-", endPorts[i])
    //     for j := startPorts[i]; j <= endPorts[i]; j++ {
    //         port := <-results[i]
    //         if port != 0 {
    //             fmt.Println("Open port:", port, "| Time:", timings[j])
    //         }
    //     }
    //     fmt.Println()   
    // }

}



// Асинхронный веб-сканер портов
// Задача:
// Напиши программу, которая принимает список хостов и диапазон портов,
// проверяет их доступность параллельно (goroutines), и выдаёт JSON-отчёт.
// Условия:
//     Максимальное количество одновременных проверок (воркеров) задаётся параметром.
//     Время таймаута подключения – настраиваемое.
//     В отчёте указать статус порта (открыт/закрыт), время отклика.