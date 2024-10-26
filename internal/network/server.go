package network

import (
    "bufio"
    "fmt"
    "net"

    "github.com/bifshteksex/romansql/internal/parser"
    "github.com/bifshteksex/romansql/internal/storage"
)

type Server struct {
    Storage *storage.Storage
}

func NewServer(storage *storage.Storage) *Server {
    return &Server{Storage: storage}
}

func (s *Server) Start(address string) error {
    listener, err := net.Listen("tcp", address)
    if err != nil {
        return fmt.Errorf("не удалось запустить сервер: %v", err)
    }
    defer listener.Close()
    fmt.Printf("Сервер запущен на %s\n", address)

    for {
        conn, err := listener.Accept()
        if err != nil {
            return fmt.Errorf("ошибка подключения: %v", err)
        }
        go s.handleConnection(conn)
    }
}

func (s *Server) handleConnection(conn net.Conn) {
    defer conn.Close()
    reader := bufio.NewReader(conn)
    for {
        request, err := reader.ReadString('\n')
        if err != nil {
            fmt.Println("Ошибка чтения запроса:", err)
            return
        }
        response := s.handleRequest(request)
        conn.Write([]byte(response + "\n"))
    }
}

func (s *Server) handleRequest(request string) string {
    response, err := parser.ParseCommand(request, s.Storage)
    if err != nil {
        return fmt.Sprintf("Ошибка обработки запроса: %v", err)
    }
    return response
}
