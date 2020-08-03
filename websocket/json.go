package main

import (
    "golang.org/x/net/websocket"
    "log"
    "net/http"
    "net/http/httptest"
    "strings"
    "net"
    "fmt"
)

var serverAddr string

type Count struct {
    S string
    N int
}

func countServer(ws *websocket.Conn) {
    defer ws.Close()
    for {
        var count Count
        err := websocket.JSON.Receive(ws, &count)
        if err != nil {
            return
        }
        count.N++
        count.S = strings.Repeat(count.S, count.N)
        err = websocket.JSON.Send(ws, count)
        if err != nil {
            return
        }
    }
}

func startServer() {
    http.Handle("/count", websocket.Handler(countServer))
    server := httptest.NewServer(nil)
    serverAddr = server.Listener.Addr().String()
    log.Print("Test WebSocket server listening on ", serverAddr)
}

func main()  {
    startServer()

    // websocket.Dial()
    client, err := net.Dial("tcp", serverAddr)
    if err != nil {
        log.Fatal("dialing", err)
    }

    config, _ :=  websocket.NewConfig(fmt.Sprintf("ws://%s/count", serverAddr), "http://localhost")
    conn, err := websocket.NewClient(config, client)
    if err != nil {
        log.Printf("WebSocket handshake error: %v", err)
        return
    }

    var count Count
    count.S = "hello"
    if err := websocket.JSON.Send(conn, count); err != nil {
        log.Printf("Write: %v", err)
    }
    if err := websocket.JSON.Receive(conn, &count); err != nil {
        log.Printf("Read: %v", err)
    }
    if count.N != 1 {
        log.Printf("count: expected %d got %d", 1, count.N)
    }
    if count.S != "hello" {
        log.Printf("count: expected %q got %q", "hello", count.S)
    }
    if err := websocket.JSON.Send(conn, count); err != nil {
        log.Printf("Write: %v", err)
    }
    if err := websocket.JSON.Receive(conn, &count); err != nil {
        log.Printf("Read: %v", err)
    }
    if count.N != 2 {
        log.Printf("count: expected %d got %d", 2, count.N)
    }
    if count.S != "hellohello" {
        log.Printf("count: expected %q got %q", "hellohello", count.S)
    }
    conn.Close()
}