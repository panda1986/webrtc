package main

import (
    "golang.org/x/net/websocket"
    "log"
    "fmt"
)

var origin = "http://127.0.0.1:8888/"
var url = "ws://127.0.0.1:8888/signal"

func main() {
    ws, err := websocket.Dial(url, "", origin)
    if err != nil {
        log.Fatalf("connect ws server:%v failed, err is %v", url, err)
    }
    defer ws.Close()

    info := &struct {
        Type string `json:"type"`
        Name string `json:"name"`
    }{
        Type: "login",
        Name: "panda",
    }

    if err := websocket.JSON.Send(ws, info); err != nil {
        log.Println(fmt.Sprintf("send ws msg failed, err is %v", err))
        return
    }
    log.Println(fmt.Sprintf("logined as %v", info.Name))

    res := &struct {
        Type string `json:"type"`
        Success bool `json:"success"`
    }{}
    if err := websocket.JSON.Receive(ws, res); err != nil {
        log.Println(fmt.Sprintf("recv ws msg failed, err is %v", err))
    } else {
        log.Println(fmt.Sprintf("got result:%+v", res))
    }
}