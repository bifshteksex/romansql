package commands

import (
    "fmt"
    "strings"

    "github.com/bifshteksex/romansql/internal/storage"
)

func HandleCreateTable(args []string, store *storage.Storage) (string, error) {
    if len(args) < 3 || strings.ToUpper(args[0]) != "TABLE" {
        return "", fmt.Errorf("Неправильный формат команды CREATE TABLE")
    }

    tableName := args[1]
    columnsArgs := strings.Join(args[2:], " ")

    // Парсим колонки (формат: columnName columnType, ...)
    columns := []storage.Column{}
    columnsSpec := strings.Split(columnsArgs, ",")
    for _, col := range columnsSpec {
        colDetails := strings.Fields(strings.TrimSpace(col))
        if len(colDetails) != 2 {
            return "", fmt.Errorf("Неправильное определение колонки: %s", col)
        }
        columns = append(columns, storage.Column{Name: colDetails[0], Type: colDetails[1]})
    }

    return store.CreateTable(tableName, columns)
}
