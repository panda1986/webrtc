package main

import (
    "golang.org/x/net/websocket"
    "fmt"
    "log"
    "net/http"
    "encoding/json"
)

var users map[string]string

type BasicSignalInfo struct {
    Type string `json:"type"`
}

type Reader struct {
    name string
    otherName string
    conn *websocket.Conn
}

func (v *Reader) dealLogin(data []byte) (err error) {
    m := &struct {
        Name string `json:"name"`
    }{}

    res := &struct {
        Type string `json:"type"`
        Success bool `json:"success"`
    }{
        Type: "login",
    }

    defer func() {
        if err != nil {
            res.Success  =false
        } else {
            res.Success = true
        }
        body, _ := json.Marshal(res)
        v.conn.Write(body)
    }()

    if err = json.Unmarshal([]byte(data), m); err != nil {
        log.Println(fmt.Sprintf("decode login name failed, err is %v", err))
        return
    }

    if _, ok := users[m.Name]; ok {
        err = fmt.Errorf("name:%v already exist", m.Name)
        log.Println(err.Error())
        return
    }

    users[m.Name] = m.Name
    v.name = m.Name
    return
}

func (v *Reader) serve() (err error) {
    msg := make([]byte, 512)
    n, err := v.conn.Read(msg)
    if err != nil {
        log.Println(fmt.Sprintf("read msg failed, err is %v", err))
        return
    }

    data := msg[:n]
    log.Println(fmt.Printf("got signal info:%v", string(data)))

    si := &BasicSignalInfo{}
    if err = json.Unmarshal(data, si); err != nil {
        log.Println(fmt.Sprintf("decode signal info failed, err is %v", err))
        return
    }

    switch si.Type {
    case "login":
        log.Println(fmt.Sprintf("got login signal, come to deal it"))
        v.dealLogin(data)
    }
    return
}

func signalHandler(ws *websocket.Conn) {
    log.Println(fmt.Sprintf("signal handler connected"))
    defer ws.Close()
    r := &Reader{
        conn: ws,
    }
    for {
        if err := r.serve(); err != nil {
            return
        }
    }
}

func main() {
    users = make(map[string]string)

    log.SetFlags(log.Ldate | log.Ltime)
    http.Handle("/signal", websocket.Handler(signalHandler))
    err := http.ListenAndServe(":8888", nil)
    if err != nil {
        panic("ListenAndServe: " + err.Error())
    }
}