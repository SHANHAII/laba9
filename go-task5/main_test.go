package main

import (
    "encoding/json"
    "net"
    "testing"
    "time"
)

func TestTCPServer(t *testing.T) {
    go func() {
        listener, _ := net.Listen("tcp", ":9000")
        defer listener.Close()
        
        for {
            conn, _ := listener.Accept()
            go handleConnection(conn)
        }
    }()
    
    time.Sleep(100 * time.Millisecond)
    
    conn, err := net.Dial("tcp", "localhost:9000")
    if err != nil {
        t.Fatal("Не удалось подключиться:", err)
    }
    defer conn.Close()
    
    req := Request{Numbers: []int{1, 2, 3, 4, 5}}
    encoder := json.NewEncoder(conn)
    if err := encoder.Encode(req); err != nil {
        t.Fatal("Ошибка отправки:", err)
    }
    
    var resp Response
    decoder := json.NewDecoder(conn)
    if err := decoder.Decode(&resp); err != nil {
        t.Fatal("Ошибка получения:", err)
    }
    
    if resp.Sum != 55 {
        t.Errorf("Ожидалось 55, получено %d", resp.Sum)
    }
}