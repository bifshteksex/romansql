package parser

import (
    "strings"
    
    "github.com/bifshteksex/romansql/internal/commands"
    "github.com/bifshteksex/romansql/internal/storage"
)

func ParseCommand(query string, store *storage.Storage) (string, error) {
    parts := strings.Fields(query)
    if len(parts) == 0 {
        return "", nil
    }

    switch strings.ToUpper(parts[0]) {
    case "CREATE":
        if len(parts) > 1 && strings.ToUpper(parts[1]) == "TABLE" {
            return commands.HandleCreateTable(parts[1:], store)
        }
    case "DROP":
        if len(parts) > 1 && strings.ToUpper(parts[1]) == "TABLE" {
            return commands.HandleDropTable(parts[1:], store)
        }
    case "INSERT":
        if len(parts) > 1 && strings.ToUpper(parts[1]) == "INTO" {
            return commands.HandleInsert(parts[1:], store)
        }
    case "DELETE":
        if len(parts) > 1 && strings.ToUpper(parts[1]) == "FROM" {
            return commands.HandleDelete(parts[1:], store)
        }
    case "UPDATE":
        return commands.HandleUpdate(parts[1:], store)
    case "SELECT":
        return commands.HandleSelect(parts[1:], store)
    default:
        return "Неизвестная команда", nil
    }
    return "Неправильный запрос", nil
}
