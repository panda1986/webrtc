package main

import (
    "golang.org/x/net/websocket"
    "fmt"
    "log"
    "net/http"
    "encoding/json"
)

var users map[string]*Reader

type BasicSignalInfo struct {
    Type string `json:"type"`
    Name string `json:"name"`
    Body interface{} `json:"body"`
}

type Reader struct {
    name string
    otherName string
    conn *websocket.Conn
}

func (v *Reader) dealLogin(si *BasicSignalInfo) (err error) {
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

    if _, ok := users[si.Name]; ok {
        err = fmt.Errorf("name:%v already exist", si.Name)
        log.Println(err.Error())
        return
    }

    users[si.Name] = v
    v.name = si.Name
    return
}

func (v *Reader) dealOffer(si *BasicSignalInfo) (err error) {
    m := &struct {
        Offer interface{} `json:"offer"`
    }{}

    body, _ := json.Marshal(si.Body)
    if err = json.Unmarshal(body, m); err != nil {
        log.Println(fmt.Sprintf("decode offer failed, err is %v", err))
        return
    }

    if r, ok := users[si.Name]; !ok {
        err = fmt.Errorf("deal user:%v offer failed, not exist", si.Name)
        log.Println(err.Error())
        return
    } else {
        v.otherName = si.Name
        res := &struct {
            Type string `json:"type"`
            Offer interface{} `json:"offer"`
            Name string `json:"name"`
        }{
            Type: "offer",
            Offer: m.Offer,
            Name: v.name,
        }
        body, _ := json.Marshal(res)
        r.conn.Write(body)
    }
    return
}

func (v *Reader) dealAnswer(si *BasicSignalInfo) (err error) {
    m := &struct {
        Answer interface{} `json:"answer"`
    }{}

    body, _ := json.Marshal(si.Body)
    if err = json.Unmarshal(body, m); err != nil {
        log.Println(fmt.Sprintf("decode answer failed, err is %v", err))
        return
    }

    r, ok := users[si.Name]
    if !ok {
        err = fmt.Errorf("deal user:%v answer failed, not exist", si.Name)
        log.Println(err.Error())
        return
    }

    v.otherName = si.Name
    res := &struct {
        Type string `json:"type"`
        Answer interface{} `json:"answer"`
    }{
        Type: "answer",
        Answer: m.Answer,
    }
    body, _ = json.Marshal(res)
    r.conn.Write(body)
    return
}

func (v *Reader) dealCandidate(si *BasicSignalInfo) (err error) {
    m := &struct {
        Candidate interface{} `json:"candidate"`
    }{}

    body, _ := json.Marshal(si.Body)
    if err = json.Unmarshal(body, m); err != nil {
        log.Println(fmt.Sprintf("decode answer failed, err is %v", err))
        return
    }

    r, ok := users[si.Name]
    if !ok {
        err = fmt.Errorf("deal user:%v candidate failed, not exist", si.Name)
        log.Println(err.Error())
        return
    }
    v.otherName = si.Name

    res := &struct {
        Type string `json:"type"`
        Candidate interface{} `json:"candidate"`
    }{
        Type: "candidate",
        Candidate: m.Candidate,
    }
    body, _ = json.Marshal(res)
    r.conn.Write(body)
    return
}

func (v *Reader) dealLeave(si *BasicSignalInfo) (err error) {
    r, ok := users[si.Name]
    if !ok {
        err = fmt.Errorf("deal user:%v leave failed, not exist", si.Name)
        log.Println(err.Error())
        return
    }

    r.otherName = ""

    res := &struct {
        Type string `json:"type"`
    }{
        Type: "leave",
    }
    body, _ := json.Marshal(res)
    r.conn.Write(body)
    return
}

func (v *Reader) serve() (err error) {
    si := &BasicSignalInfo{}
    if err = websocket.JSON.Receive(v.conn, si); err != nil {
        log.Println(fmt.Sprintf("decode signal info failed, err is %v", err))
        return
    }
    log.Println(fmt.Sprintf("got signal info:%+v", si))

    switch si.Type {
    case "login":
        log.Println(fmt.Sprintf("got login signal, come to deal it"))
        v.dealLogin(si)
    case "offer":
        log.Println(fmt.Sprintf("got offer signal, come to deal it"))
        v.dealOffer(si)
    case "answer":
        log.Println(fmt.Sprintf("got answer signal, come to deal it"))
        v.dealAnswer(si)
    case "candidate":
        log.Println(fmt.Sprintf("got candidate signal, come to deal it"))
        v.dealCandidate(si)
    case "leave":
        log.Println(fmt.Sprintf("got leave signal, come to deal it"))
        v.dealLeave(si)
    default:
        res := &struct {
            Type string `json:"type"`
            Message string `json:"message"`
        }{
            Type: "error",
            Message: "unrecognized command" + si.Type,
        }
        body, _ := json.Marshal(res)
        v.conn.Write(body)
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
    users = make(map[string]*Reader)

    log.SetFlags(log.Ldate | log.Ltime)
    http.Handle("/signal", websocket.Handler(signalHandler))
    err := http.ListenAndServe(":8888", nil)
    if err != nil {
        panic("ListenAndServe: " + err.Error())
    }
}