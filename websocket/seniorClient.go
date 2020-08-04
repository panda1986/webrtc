package main

import (
    netUrl "net/url"
    "log"
    "fmt"
    "github.com/gorilla/websocket"
    "time"
    "os"
    "os/signal"
)

func main()  {
    interrupt := make(chan os.Signal, 1)
    signal.Notify(interrupt, os.Interrupt, os.Kill)

    u := netUrl.URL{
        Scheme: "ws",
        Host: "localhost:8888",
        Path: "/echo",
    }
    log.Println(fmt.Sprintf("connecting to %v", u.String()))

    c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
    if err != nil {
        log.Println(fmt.Sprintf("dial failed, err is %v", err))
        return
    }
    defer c.Close()

    done := make(chan struct{})
    go func() {
        defer close(done)
        for {
            _, message, err := c.ReadMessage()
            if err != nil {
                log.Println(fmt.Sprintf("read failed, err is %v", err))
                return
            }
            log.Println(fmt.Sprintf("recv %v message", string(message)))
        }
    }()

    ticker := time.NewTicker(time.Second)
    defer ticker.Stop()

    for {
        select {
        case <- done:
            return
        case t := <- ticker.C:
            if err := c.WriteMessage(websocket.TextMessage, []byte(t.String())); err != nil {
                log.Println(fmt.Sprintf("write failed, err is %v", err))
                return
            }
        case <- interrupt:
            log.Println(fmt.Sprintf("interrupt"))
            // Cleanly close the connection by sending a close message and then waiting for the server to close the connection
            if err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "")); err !=  nil {
                log.Println(fmt.Sprintf("write close failed, err is %v", err))
                return
            }
            select {
            case <- done:
            case <- time.After(time.Second):
                return
            }

        }
    }
}