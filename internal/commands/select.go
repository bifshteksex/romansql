package commands

import (
    "fmt"
    "strings"

    "github.com/bifshteksex/romansql/internal/storage"
)

// parseCondition парсит условие WHERE (например, age > 30)
func parseCondition(condition string) (func(storage.Row) bool, error) {
    parts := strings.Fields(condition)
    if len(parts) != 3 {
        return nil, fmt.Errorf("Неправильный формат условия WHERE: %s", condition)
    }

    column, operator, value := parts[0], parts[1], parts[2]
    return func(row storage.Row) bool {
        rowVal, ok := row[column]
        if !ok {
            return false
        }
        switch operator {
        case "=":
            return fmt.Sprintf("%v", rowVal) == value
        case ">":
            return fmt.Sprintf("%v", rowVal) > value
        case "<":
            return fmt.Sprintf("%v", rowVal) < value
        default:
            return false
        }
    }, nil
}

func HandleSelect(args []string, store *storage.Storage) (string, error) {
    if len(args) < 4 || strings.ToUpper(args[1]) != "FROM" {
        return "", fmt.Errorf("Неправильный формат команды SELECT")
    }

    // Парсим колонки и таблицу
    columns := strings.Split(args[0], ",")
    tableName := args[2]
    table, err := store.GetTable(tableName)

    if err != nil {
        return "", fmt.Errorf("Таблица %s не найдена", tableName)
    }

    // Парсим условие WHERE (если оно есть)
    var condition func(storage.Row) bool
    if len(args) > 4 && strings.ToUpper(args[3]) == "WHERE" {
        condStr := strings.Join(args[4:], " ")
        cond, err := parseCondition(condStr)
        if err != nil {
            return "", err
        }
        condition = cond
    } else {
        // Если нет условия WHERE, выбираем все строки
        condition = func(row storage.Row) bool { return true }
    }

    // Выполняем выборку
    rows, err := table.SelectRows(columns, condition)
    if err != nil {
        return "", err
    }

    // Формируем результат для вывода
    var result strings.Builder
    for _, row := range rows {
        rowStr := []string{}
        for _, col := range columns {
            rowStr = append(rowStr, fmt.Sprintf("%v", row[col]))
        }
        result.WriteString(strings.Join(rowStr, " | "))
        result.WriteString("\n")
    }
    return result.String(), nil
}
