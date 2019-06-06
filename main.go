package main

import (
    "flag"
    "os"
    "log"
    "sync"
    "wsserver/ws"
    "wsserver/rest"
)

var (
    h bool
    wsip string
    restip string
    resturl string
    wsurl string
    wscert string
    wskey string

    wg sync.WaitGroup
    wschan chan string
)

func wstask(wschan chan string){
    ws.Server(wsip, wsurl, wscert, wskey, wschan)
}

func resttask(wschan chan string){
    rest.Server(restip, resturl, wschan)
}

func paramProc(){
    flag.BoolVar(&h, "h", false, "cmd -ip \"0.0.0.0:3333\" -url \"ws\"")
    flag.StringVar(&wsip, "wsip", "0.0.0.0:3333", "websocket ip address")
    flag.StringVar(&wsurl, "wsurl", "ws", "ws url")
    flag.StringVar(&restip, "restip", "0.0.0.0:3334", "rest ip address")
    flag.StringVar(&resturl, "resturl", "rest", "rest url")
    flag.StringVar(&wscert, "wscert", "cert/test.pem", "websocket ssl cert")
    flag.StringVar(&wskey, "wskey", "cert/test.key", "websocket ssl key")
    flag.Parse()

    if h{
        flag.PrintDefaults()
        os.Exit(0)
    }
    resturl = "/" + resturl
    wsurl = "/" + wsurl
}

func main(){

    paramProc()
    wg.Add(1)
    wschan = make(chan string, 20)

    go wstask(wschan)
    go resttask(wschan)

    log.Println("running...")
    wg.Wait()
}
