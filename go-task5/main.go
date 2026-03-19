package main

import (
    "encoding/json"
    "fmt"
    "net"
    "os"
)

type Request struct {
    Numbers []int `json:"numbers"`
}

type Response struct {
    Sum int `json:"sum"`
}

func handleConnection(conn net.Conn) {
    defer conn.Close()
    
    var req Request
    decoder := json.NewDecoder(conn)
    if err := decoder.Decode(&req); err != nil {
        return
    }
    
    sum := 0
    for _, n := range req.Numbers {
        sum += n * n
    }
    
    resp := Response{Sum: sum}
    encoder := json.NewEncoder(conn)
    encoder.Encode(resp)
}

func main() {
    if len(os.Args) > 1 && os.Args[1] == "server" {
        listener, err := net.Listen("tcp", ":9000")
        if err != nil {
            fmt.Println("Ошибка:", err)
            return
        }
        defer listener.Close()
        
        fmt.Println("TCP сервер запущен на порту 9000")
        
        for {
            conn, err := listener.Accept()
            if err != nil {
                continue
            }
            go handleConnection(conn)
        }
    } else {
        var req Request
        json.NewDecoder(os.Stdin).Decode(&req)
        
        sum := 0
        for _, n := range req.Numbers {
            sum += n * n
        }
        
        resp := Response{Sum: sum}
        json.NewEncoder(os.Stdout).Encode(resp)
    }
}