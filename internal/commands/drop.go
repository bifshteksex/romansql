package commands

import (
    "fmt"
    "strings"

    "github.com/bifshteksex/romansql/internal/storage"
)

func HandleDropTable(args []string, store *storage.Storage) (string, error) {
    if len(args) < 2 || strings.ToUpper(args[0]) != "TABLE" {
        return "", fmt.Errorf("Неправильный формат команды DROP TABLE")
    }

    tableName := args[1]

    return store.DropTable(tableName)
}
