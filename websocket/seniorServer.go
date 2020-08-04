package main

import (
    "github.com/gorilla/websocket"
    "net/http"
    "log"
    "fmt"
)

// Upgrader specifies parameters for upgrading an HTTP connection to a WebSocket connection.
var upgrader = websocket.Upgrader{}

func main()  {
    http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
        log.Println(fmt.Sprintf("got a echo request"))
        c, err := upgrader.Upgrade(w, r, nil)
        if err != nil {
            log.Println(fmt.Sprintf("upgrade failed, err is %v", err))
            return
        }
        defer c.Close()
        for {
            mt, message, err := c.ReadMessage()
            if err != nil {
                log.Println(fmt.Sprintf("read failed, err is %v", err))
                break
            }
            log.Println(fmt.Sprintf("recv message:%v, %v", mt, string(message)))
            if err = c.WriteMessage(mt, message); err != nil {
                log.Println(fmt.Sprintf("write failed, err is %v", err))
                break
            }
        }
    })
    log.Println(http.ListenAndServe(":8888", nil))
}