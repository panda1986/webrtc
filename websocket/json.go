package main

import (
    "golang.org/x/net/websocket"
    "log"
    "net/http"
    "net/http/httptest"
    "net"
    "fmt"
    "encoding/json"
)

var serverAddr string

type Count struct {
    S interface{}
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
    count.S = "90:83:89:DC:5D:3A:C7:DE:13:FB:CB:E1:1A:10:9B:29:76:98:E5:F5:4F:AF:16:D7:EA:A7:05:53:CB:AB:9B:62a=fingerprint:sha-256 90:83:89:DC:5D:3A:C7:DE:13:FB:CB:E1:1A:10:9B:29:76:98:E5:F5:4F:AF:16:D7:EA:A7:05:53:CB:AB:9B:62a=fingerprint:sha-256 90:83:89:DC:5D:3A:C7:DE:13:FB:CB:E1:1A:10:9B:29:76:98:E5:F5:4F:AF:16:D7:EA:A7:05:53:CB:AB:9B:62a=fingerprint:sha-256 90:83:89:DC:5D:3A:C7:DE:13:FB:CB:E1:1A:10:9B:29:76:98:E5:F5:4F:AF:16:D7:EA:A7:05:53:CB:AB:9B:62a=fingerprint:sha-256 90:83:89:DC:5D:3A:C7:DE:13:FB:CB:E1:1A:10:9B:29:76:98:E5:F5:4F:AF:16:D7:EA:A7:05:53:CB:AB:9B:62a=fingerprint:sha-256 90:83:89:DC:5D:3A:C7:DE:13:FB:CB:E1:1A:10:9B:29:76:98:E5:F5:4F:AF:16:D7:EA:A7:05:53:CB:AB:9B:62a=fingerprint:sha-256 90:83:89:DC:5D:3A:C7:DE:13:FB:CB:E1:1A:10:9B:29:76:98:E5:F5:4F:AF:16:D7:EA:A7:05:53:CB:AB:9B:62a=fingerprint:sha-256 90:83:89:DC:5D:3A:C7:DE:13:FB:CB:E1:1A:10:9B:29:76:98:E5:F5:4F:AF:16:D7:EA:A7:05:53:CB:AB:9B:62a=fingerprint:sha-256 90:83:89:DC:5D:3A:C7:DE:13:FB:CB:E1:1A:10:9B:29:76:98:E5:F5:4F:AF:16:D7:EA:A7:05:53:CB:AB:9B:62a=fingerprint:sha-256 90:83:89:DC:5D:3A:C7:DE:13:FB:CB:E1:1A:10:9B:29:76:98:E5:F5:4F:AF:16:D7:EA:A7:05:53:CB:AB:9B:62a=fingerprint:sha-256 90:83:89:DC:5D:3A:C7:DE:13:FB:CB:E1:1A:10:9B:29:76:98:E5:F5:4F:AF:16:D7:EA:A7:05:53:CB:AB:9B:62a=fingerprint:sha-256 90:83:89:DC:5D:3A:C7:DE:13:FB:CB:E1:1A:10:9B:29:76:98:E5:F5:4F:AF:16:D7:EA:A7:05:53:CB:AB:9B:62a=fingerprint:sha-256 90:83:89:DC:5D:3A:C7:DE:13:FB:CB:E1:1A:10:9B:29:76:98:E5:F5:4F:AF:16:D7:EA:A7:05:53:CB:AB:9B:62a=fingerprint:sha-256 90:83:89:DC:5D:3A:C7:DE:13:FB:CB:E1:1A:10:9B:29:76:98:E5:F5:4F:AF:16:D7:EA:A7:05:53:CB:AB:9B:62a=fingerprint:sha-256 90:83:89:DC:5D:3A:C7:DE:13:FB:CB:E1:1A:10:9B:29:76:98:E5:F5:4F:AF:16:D7:EA:A7:05:53:CB:AB:9B:62a=fingerprint:sha-256 90:83:89:DC:5D:3A:C7:DE:13:FB:CB:E1:1A:10:9B:29:76:98:E5:F5:4F:AF:16:D7:EA:A7:05:53:CB90:83:89:DC:5D:3A:C7:DE:13:FB:CB:E1:1A:10:9B:29:76:98:E5:F5:4F:AF:16:D7:EA:A7:05:53:CB:AB:9B:62a=fingerprint:sha-256 90:83:89:DC:5D:3A:C7:DE:13:FB:CB:E1:1A:10:9B:29:76:98:E5:F5:4F:AF:16:D7:EA:A7:05:53:CB:AB:9B:62a=fingerprint:sha-256 90:83:89:DC:5D:3A:C7:DE:13:FB:CB:E1:1A:10:9B:29:76:98:E5:F5:4F:AF:16:D7:EA:A7:05:53:CB:AB:9B:62a=fingerprint:sha-256 90:83:89:DC:5D:3A:C7:DE:13:FB:CB:E1:1A:10:9B:29:76:98:E5:F5:4F:AF:16:D7:EA:A7:05:53:CB:AB:9B:62a=fingerprint:sha-256 90:83:89:DC:5D:3A:C7:DE:13:FB:CB:E1:1A:10:9B:29:76:98:E5:F5:4F:AF:16:D7:EA:A7:05:53:CB:AB:9B:62a=fingerprint:sha-256 90:83:89:DC:5D:3A:C7:DE:13:FB:CB:E1:1A:10:9B:29:76:98:E5:F5:4F:AF:16:D7:EA:A7:05:53:CB:AB:9B:62a=fingerprint:sha-256 90:83:89:DC:5D:3A:C7:DE:13:FB:CB:E1:1A:10:9B:29:76:98:E5:F5:4F:AF:16:D7:EA:A7:05:53:CB:AB:9B:62a=fingerprint:sha-256 90:83:89:DC:5D:3A:C7:DE:13:FB:CB:E1:1A:10:9B:29:76:98:E5:F5:4F:AF:16:D7:EA:A7:05:53:CB:AB:9B:62a=fingerprint:sha-256 90:83:89:DC:5D:3A:C7:DE:13:FB:CB:E1:1A:10:9B:29:76:98:E5:F5:4F:AF:16:D7:EA:A7:05:53:CB:AB:9B:62a=fingerprint:sha-256 90:83:89:DC:5D:3A:C7:DE:13:FB:CB:E1:1A:10:9B:29:76:98:E5:F5:4F:AF:16:D7:EA:A7:05:53:CB:AB:9B:62a=fingerprint:sha-256 90:83:89:DC:5D:3A:C7:DE:13:FB:CB:E1:1A:10:9B:29:76:98:E5:F5:4F:AF:16:D7:EA:A7:05:53:CB:AB:9B:62a=fingerprint:sha-256 90:83:89:DC:5D:3A:C7:DE:13:FB:CB:E1:1A:10:9B:29:76:98:E5:F5:4F:AF:16:D7:EA:A7:05:53:CB:AB:9B:62a=fingerprint:sha-256 90:83:89:DC:5D:3A:C7:DE:13:FB:CB:E1:1A:10:9B:29:76:98:E5:F5:4F:AF:16:D7:EA:A7:05:53:CB:AB:9B:62a=fingerprint:sha-256 90:83:89:DC:5D:3A:C7:DE:13:FB:CB:E1:1A:10:9B:29:76:98:E5:F5:4F:AF:16:D7:EA:A7:05:53:CB:AB:9B:62a=fingerprint:sha-256 90:83:89:DC:5D:3A:C7:DE:13:FB:CB:E1:1A:10:9B:29:76:98:E5:F5:4F:AF:16:D7:EA:A7:05:53:CB:AB:9B:62a=fingerprint:sha-256 90:83:89:DC:5D:3A:C7:DE:13:FB:CB:E1:1A:10:9B:29:76:98:E5:F5:4F:AF:19:DC:5D:3A:C7:DE:13:FB:CB:E1:1A:10:9B:29:76:98:E5:F5:4F:AF:16:D7:EA:A7:05:53:CB:AB:9B:62a=fingerprint:sha-256 90:83:89:DC:5D:3A:C7:DE:13:FB:CB:E1:1A:10:9B:29:76:98:E5:F5:4F:AF:16:D7:EA:A7:05:53:CB:AB:9B:62a=fingerprint:sha-256 90:83:89:DC:5D:3A:C7:DE:13:FB:CB:E1:1A:10:9B:29:76:98:E5:F5:4F:AF:16:D7:EA:A7:05:53:CB:AB:9B:62a=fingerprint:sha-256 90:83:89:DC:5D:3A:C7:DE:13:FB:CB:E1:1A:10:9B:29:76:98:E5:F5:4F:AF:16:D7:EA:A7:05:53:CB:AB:9B:62a=fingerprint:sha-256 90:83:89:DC:5D:3A:C7:DE:13:FB:CB:E1:1A:10:9B:29:76:98:E5:F5:4F:AF:16:D7:EA:A7:05:53:CB:AB:9B:62a=fingerprint:sha-256 90:83:89:DC:5D:3A:C7:DE:13:FB:CB:E1:1A:10:9B:29:76:98:E5:F5:4F:AF:16:D7:EA:A7:05:53:CB90:83:89:DC:5D:3A:C7:DE:13:FB:CB:E1:1A:10:9B:29:76:98:E5:F5:4F:AF:16:D7:EA:A7:05:53:CB:AB:9B:62a=fingerprint:sha-256 90:83:89:DC:5D:3A:C7:DE:13:FB:CB:E1:1A:10:9B:29:76:98:E5:F5:4F:AF:16:D7:EA:A7:05:53:CB:AB:9B:62a=fingerprint:sha-256 90:83:89:DC:5D:3A:C7:DE:13:FB:CB:E1:1A:10:9B:29:76:98:E5:F5:4F:AF:16:D7:EA:A7:05:53:CB:AB:9B:62a=fingerprint:sha-256 90:83:89:DC:5D:3A:C7:DE:13:FB:CB:E1:1A:10:9B:29:76:98:E5:F5:4F:AF:16:D7:EA:A7:05:53:CB:AB:9B:62a=fingerprint:sha-256 90:83:89:DC:5D:3A:C7:DE:13:FB:CB:E1:1A:10:9B:29:76:98:E5:F5:4F:AF:16:D7:EA:A7:05:53:CB:AB:9B:62a=fingerprint:sha-256 90:83:89:DC:5D:3A:C7:DE:13:FB:CB:E1:1A:10:9B:29:76:98:E5:F5:4F:AF:16:D7:EA:A7:05:53:CB:AB:9B:62a=fingerprint:sha-256 90:83:89:DC:5D:3A:C7:DE:13:FB:CB:E1:1A:10:9B:29:76:98:E5:F5:4F:AF:16:D7:EA:A7:05:53:CB:AB:9B:62a=fingerprint:sha-256 90:83:89:DC:5D:3A:C7:DE:13:FB:CB:E1:1A:10:9B:29:76:98:E5:F5:4F:AF:16:D7:EA:A7:05:53:CB:AB:9B:62a=fingerprint:sha-256 90:83:89:DC:5D:3A:C7:DE:13:FB:CB:E1:1A:10:9B:29:76:98:E5:F5:4F:AF:16:D7:EA:A7:05:53:CB:AB:9B:62a=fingerprint:sha-256 90:83:89:DC:5D:3A:C7:DE:13:FB:CB:E1:1A:10:9B:29:76:98:E5:F5:4F:AF:16:D7:EA:A7:05:53:CB:AB:9B:62a=fingerprint:sha-256 90:83:89:DC:5D:3A:C7:DE:13:FB:CB:E1:1A:10:9B:29:76:98:E5:F5:4F:AF:16:D7:EA:A7:05:53:CB:AB:9B:62a=fingerprint:sha-256 90:83:89:DC:5D:3A:C7:DE:13:FB:CB:E1:1A:10:9B:29:76:98:E5:F5:4F:AF:16:D7:EA:A7:05:53:CB:AB:9B:62a=fingerprint:sha-256 90:83:89:DC:5D:3A:C7:DE:13:FB:CB:E1:1A:10:9B:29:76:98:E5:F5:4F:AF:16:D7:EA:A7:05:53:CB:AB:9B:62a=fingerprint:sha-256 90:83:89:DC:5D:3A:C7:D"
    body, _ := json.Marshal(count)
    n, err := conn.Write(body)
    if err != nil {
        log.Println(fmt.Sprintf("write failed, err is %v", err))
        return
    }
    log.Println(fmt.Sprintf("write %v bytes", n))
    //if err := websocket.JSON.Send(conn, count); err != nil {
    //    log.Printf("Write: %v", err)
    //}

    msg := make([]byte, 6000)
    n1, err := conn.Read(msg)
    if err != nil {
        log.Println(fmt.Sprintf("read failed, err is %v", err))
        return
    }
    log.Println(fmt.Sprintf("read %v bytes", n1))

    //newc := &Count{}
    //if err := websocket.JSON.Receive(conn, newc); err != nil {
    //    log.Printf("Read: %v", err)
    //}
    //log.Println(fmt.Sprintf("got count souce:%v", newc.S))
    //if count.N != 1 {
    //    log.Printf("count: expected %d got %d", 1, count.N)
    //}
    //if count.S != "hello" {
    //    log.Printf("count: expected %q got %q", "hello", count.S)
    //}
    //if err := websocket.JSON.Send(conn, count); err != nil {
    //    log.Printf("Write: %v", err)
    //}
    //if err := websocket.JSON.Receive(conn, &count); err != nil {
    //    log.Printf("Read: %v", err)
    //}
    //if count.N != 2 {
    //    log.Printf("count: expected %d got %d", 2, count.N)
    //}
    //if count.S != "hellohello" {
    //    log.Printf("count: expected %q got %q", "hellohello", count.S)
    //}
    conn.Close()
}