package commands

import (
    "fmt"
    "strconv"
    "strings"

    "github.com/bifshteksex/romansql/internal/storage"
)

func HandleInsert(args []string, store *storage.Storage) (string, error) {
    if len(args) < 4 || strings.ToUpper(args[0]) != "INTO" {
        return "", fmt.Errorf("Неправильный формат команды INSERT INTO")
    }

    tableName := args[1]
    table, err := store.GetTable(tableName)

    if err != nil {
        return "", fmt.Errorf("Таблица %s не найдена", tableName)
    }

    // Пример: INSERT INTO users VALUES (1, 'Alice', 30, 'alice@example.com')
    valuesArgs := strings.Trim(strings.Join(args[2:], " "), "()")
    valuesStr := strings.Split(valuesArgs, ",")

    if len(valuesStr) != len(table.Columns) {
        return "", fmt.Errorf("Количество значений не соответствует количеству колонок")
    }

    row := storage.Row{}
    for i, col := range table.Columns {
        valStr := strings.TrimSpace(valuesStr[i])
        switch col.Type {
        case "INT":
            // Преобразуем строку в int
            value, err := strconv.Atoi(valStr)
            if err != nil {
                return "", fmt.Errorf("Неправильный тип данных для колонки %s: ожидается INT", col.Name)
            }
            row[col.Name] = value
        case "STRING":
            // Для строк просто присваиваем
            row[col.Name] = valStr[1 : len(valStr)-1] // Убираем одинарные кавычки
        default:
            return "", fmt.Errorf("Неизвестный тип данных для колонки %s: %s", col.Name, col.Type)
        }
    }

    if err := table.AddRow(row); err != nil {
        return "", err
    }
    return fmt.Sprintf("Строка добавлена в таблицу %s", tableName), nil
}
