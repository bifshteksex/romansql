package commands

import (
    "fmt"
    "strings"

	"github.com/bifshteksex/romansql/internal/storage"
)

func HandleDelete(args []string, store *storage.Storage) (string, error) {
    if len(args) < 3 || strings.ToUpper(args[0]) != "FROM" {
        return "", fmt.Errorf("Неправильный формат команды DELETE FROM")
    }

    tableName := args[1]
    table, err := store.GetTable(tableName)

    if err != nil {
        return "", fmt.Errorf("Таблица %s не найдена", tableName)
    }

    // Условия для WHERE
    condition := func(row storage.Row) bool { return true } // Пока без условий

    count := table.DeleteRows(condition)
    return fmt.Sprintf("Удалено строк: %d из таблицы %s", count, tableName), nil
}
