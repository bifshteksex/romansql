package commands

import (
    "fmt"
    "strings"

	"github.com/bifshteksex/romansql/internal/storage"
)

func HandleUpdate(args []string, store *storage.Storage) (string, error) {
    if len(args) < 5 || strings.ToUpper(args[1]) != "SET" {
        return "", fmt.Errorf("Неправильный формат команды UPDATE")
    }

    tableName := args[0]
    table, err := store.GetTable(tableName)

    if err != nil {
        return "", fmt.Errorf("Таблица %s не найдена", tableName)
    }

    updates := storage.Row{}
    // Пример: UPDATE users SET name = 'Jane', age = 30
    setParts := strings.Split(strings.Join(args[2:], " "), ",")
    for _, part := range setParts {
        kv := strings.Split(part, "=")
        if len(kv) != 2 {
            return "", fmt.Errorf("Неправильный синтаксис SET: %s", part)
        }
        updates[strings.TrimSpace(kv[0])] = strings.TrimSpace(kv[1])
    }

    condition := func(row storage.Row) bool { return true } // Пока без условий

    count := table.UpdateRows(condition, updates)
    return fmt.Sprintf("Обновлено строк: %d в таблице %s", count, tableName), nil
}
