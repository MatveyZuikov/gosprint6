package main

import (
	"log"
	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server" 
	"os"
)

func main() {

    logger := log.New(
        os.Stdout,           
        "SERVER: ",          
        log.Ldate|log.Ltime, 
    )


    srv := server.NewServer(logger)

    if err := srv.Run(); err != nil {
        logger.Fatal("Ошибка запуска сервера:", err)
    }
} 

