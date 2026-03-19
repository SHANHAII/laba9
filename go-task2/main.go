package main

import (
    "encoding/json"
    "fmt"
    "os"
)

type Request struct {
    Numbers []int `json:"numbers"`
}

type Response struct {
    Sum int `json:"sum"`
}

func main() {
    if len(os.Args) > 1 && os.Args[1] == "server" {
        // Режим сервера
        startServer()
        return
    }

    var req Request
    err := json.NewDecoder(os.Stdin).Decode(&req)
    if err != nil {
        return
    }
    
    sum := 0
    for _, n := range req.Numbers {
        sum += n * n
    }
    
    resp := Response{Sum: sum}
    json.NewEncoder(os.Stdout).Encode(resp)
}

func startServer() {

    fmt.Println("Сервер запущен на :8080")

}