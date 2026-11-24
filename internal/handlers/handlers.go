package handlers

import (
	"net/http"
	"os"
	"io"
	"path/filepath"
    "time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service" 
)

func MainHandler(res http.ResponseWriter, req *http.Request) {

	file, err := os.Open("index.html")

	if err != nil {
		http.Error(res, "файл не найден", http.StatusInternalServerError )
		return
	}

	defer file.Close()

	_, err = io.Copy(res, file)
	if err != nil {
		http.Error(res, "Ошибка при чтении файла", http.StatusInternalServerError )
		return
	}
}

func UploadHandler(res http.ResponseWriter, req *http.Request) {

    err := req.ParseMultipartForm(32 << 20) 
    if err != nil {
        http.Error(res, "Ошибка парсинга формы", http.StatusInternalServerError)
        return
    }

    file, handler, err := req.FormFile("myFile")
    if err != nil {
        http.Error(res, "Файл не найден в форме", http.StatusInternalServerError)
        return
    }
    defer file.Close()

    data, err := io.ReadAll(file)
    if err != nil {
        http.Error(res, "Ошибка чтения файла", http.StatusInternalServerError)
        return
    }

    input := string(data)

    converted, err := service.ConvertAuto(input)
    if err != nil {
        http.Error(res, "Ошибка конвертации", http.StatusInternalServerError)
        return
    }

    timestamp := time.Now().UTC().String()
    ext := filepath.Ext(handler.Filename)
    newFilename := timestamp + ext

    newFile, err := os.Create(newFilename)
    if err != nil {
        http.Error(res, "Ошибка создания файла", http.StatusInternalServerError)
        return
    }
    defer newFile.Close()

    _, err = newFile.WriteString(converted)
    if err != nil {
        http.Error(res, "Ошибка записи в файл", http.StatusInternalServerError)
        return
    }

    res.WriteHeader(http.StatusOK)
    res.Write([]byte(converted))
} 