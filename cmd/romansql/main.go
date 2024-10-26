package main

import (
    "fmt"

    "github.com/bifshteksex/romansql/internal/storage"
    "github.com/bifshteksex/romansql/internal/network"
)

func main() {
    fmt.Println("Запуск RomanSql...")

    store := storage.NewStorage()
    server := network.NewServer(store)

    if err := server.Start(":5432"); err != nil {
        fmt.Printf("Ошибка запуска: %v\n", err)
    }
}
