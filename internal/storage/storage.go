package storage

import (
    "fmt"
)

type Row map[string]interface{}

type Column struct {
    Name string
    Type string
}

type Table struct {
    Name    string
    Columns []Column
    Data    []Row
}

type Storage struct {
    tables map[string]*Table
}

func NewStorage() *Storage {
    return &Storage{
        tables: make(map[string]*Table),
    }
}

func (s *Storage) GetTable(name string) (*Table, error) {
    table, exists := s.tables[name]
    if !exists {
        return nil, fmt.Errorf("Таблица %s не найдена", name)
    }
    return table, nil
}

func (t *Table) SelectRows(columns []string, condition func(Row) bool) ([]Row, error) {
    var result []Row

    for _, row := range t.Data {
        if condition(row) {
            selectedRow := Row{}
            if len(columns) == 1 && columns[0] == "*" { // Выбор всех колонок
                selectedRow = row
            } else {
                for _, col := range columns {
                    if val, ok := row[col]; ok {
                        selectedRow[col] = val
                    } else {
                        return nil, fmt.Errorf("Колонка %s не найдена в таблице %s", col, t.Name)
                    }
                }
            }
            result = append(result, selectedRow)
        }
    }
    return result, nil
}

func (t *Table) AddRow(row Row) error {
    for _, col := range t.Columns {
        value, exists := row[col.Name]
        if !exists {
            return fmt.Errorf("Отсутствует значение для колонки %s", col.Name)
        }

        // Проверка типов
        switch col.Type {
        case "INT":
            if _, ok := value.(int); !ok {
                return fmt.Errorf("Неправильный тип данных для колонки %s: ожидается INT", col.Name)
            }
        case "STRING":
            if _, ok := value.(string); !ok {
                return fmt.Errorf("Неправильный тип данных для колонки %s: ожидается STRING", col.Name)
            }
        default:
            return fmt.Errorf("Неизвестный тип данных для колонки %s: %s", col.Name, col.Type)
        }
    }
    t.Data = append(t.Data, row)
    return nil
}

func (t *Table) DeleteRows(condition func(Row) bool) int {
    var newData []Row
    count := 0
    for _, row := range t.Data {
        if !condition(row) {
            newData = append(newData, row)
        } else {
            count++
        }
    }
    t.Data = newData
    return count
}

func (t *Table) UpdateRows(condition func(Row) bool, updates Row) int {
    count := 0
    for i, row := range t.Data {
        if condition(row) {
            for col, val := range updates {
                row[col] = val
            }
            t.Data[i] = row
            count++
        }
    }
    return count
}

func (s *Storage) CreateTable(name string, columns []Column) (string, error) {
    if _, exists := s.tables[name]; exists {
        return "", fmt.Errorf("Таблица %s уже существует", name)
    }
    s.tables[name] = &Table{
        Name:    name,
        Columns: columns,
    }
    return fmt.Sprintf("Таблица %s успешно создана", name), nil
}

func (s *Storage) DropTable(name string) (string, error) {
    if _, exists := s.tables[name]; !exists {
        return "", fmt.Errorf("Таблица %s не найдена", name)
    }
    delete(s.tables, name)
    return fmt.Sprintf("Таблица %s успешно удалена", name), nil
}
