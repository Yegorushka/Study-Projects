package main

import (
    "fmt"
    "io"
    "net/http"
    "os"
    "path/filepath"
)

const uploadDir = "./uploads"

func uploadHandler(w http.ResponseWriter, r *http.Request) {
    // Ограничиваем размер
    r.ParseMultipartForm(10 << 20) // 10 MB

    file, handler, err := r.FormFile("file")
    if err != nil {
        http.Error(w, "Ошибка загрузки файла", http.StatusBadRequest)
        return
    }
    defer file.Close()

    dst, err := os.Create(filepath.Join(uploadDir, handler.Filename))
    if err != nil {
        http.Error(w, "Не могу сохранить файл", http.StatusInternalServerError)
        return
    }
    defer dst.Close()

    _, err = io.Copy(dst, file)
    if err != nil {
        http.Error(w, "Ошибка сохранения файла", http.StatusInternalServerError)
        return
    }

    fmt.Fprintf(w, "Файл %s загружен успешно\n", handler.Filename)
}

func downloadHandler(w http.ResponseWriter, r *http.Request) {
    filename := r.URL.Query().Get("name")
    if filename == "" {
        http.Error(w, "Не указано имя файла", http.StatusBadRequest)
        return
    }

    filepath := filepath.Join(uploadDir, filename)
    http.ServeFile(w, r, filepath)
}

func listHandler(w http.ResponseWriter, r *http.Request) {
    files, err := os.ReadDir(uploadDir)
    if err != nil {
        http.Error(w, "Не могу прочитать каталог", http.StatusInternalServerError)
        return
    }

    for _, f := range files {
        fmt.Fprintln(w, f.Name())
    }
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
    filename := r.URL.Query().Get("name")
    if filename == "" {
        http.Error(w, "Не указано имя файла", http.StatusBadRequest)
        return
    }

    err := os.Remove(filepath.Join(uploadDir, filename))
    if err != nil {
        http.Error(w, "Не могу удалить файл", http.StatusInternalServerError)
        return
    }

    fmt.Fprintf(w, "Файл %s удалён\n", filename)
}

func main() {
    // Создаём папку для хранения, если её нет
    os.MkdirAll(uploadDir, 0755)

    http.HandleFunc("/upload", uploadHandler)
    http.HandleFunc("/download", downloadHandler)
    http.HandleFunc("/list", listHandler)
    http.HandleFunc("/delete", deleteHandler)

    fmt.Println("Сервер запущен на: 8080")
    http.ListenAndServe(":8080", nil)
}


/*
5. Файловое хранилище с API
Задача:
Реализуй REST API-сервис, который:
    Принимает файлы на загрузку (multipart/form-data).
    Хранит их на диске с уникальным ID.
    Отдаёт файл по ID.
    Автоматически удаляет файлы, которые старше N дней.
    Поддерживает ограничение размера файла.
Требования:
    Используй стандартную библиотеку и net/http.
    Горутины для фоновой очистки.
*/