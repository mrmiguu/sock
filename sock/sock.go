package sock

import (
	"net/http"
	"os"
	"sync"

	"github.com/gopherjs/gopherjs/js"
	"github.com/gorilla/websocket"
	load "github.com/mrmiguu/Loading"
	"github.com/mrmiguu/jsutil"
)

var (
	Addr     = "localhost:80"
	IsClient = js.Global != nil
	Root     = "www"
	API      = "/abc123"

	start sync.Once

	ws *js.Object

	bytel sync.RWMutex
	bytem = map[string][]wrbyte{}
)

type wrbyte struct {
	key string
	idx int
	n   chan int
	w   chan<- []byte
}

func run() {
	if IsClient {
		ws = js.Global.Get("WebSocket").New("ws://" + Addr + API)
		f, c := jsutil.C()
		ws.Set("onopen", f)
		<-c
		ws.Set("onmessage", func(e *js.Object) {
			pkt := []byte(e.Get("data").String())
			println(`onmessage`, string(pkt))
		})
	} else {
		if _, err := os.Stat(Root); os.IsNotExist(err) {
			panic("root folder missing")
		}
		http.Handle("/", http.FileServer(http.Dir(Root)))
		up := websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool { return true },
		}
		http.HandleFunc(API, func(w http.ResponseWriter, r *http.Request) {
			done := load.New(`Upgrading`)
			conn, err := up.Upgrade(w, r, nil)
			done <- false
			if err != nil {
				println(err.Error())
				return
			}
			defer conn.Close()
			done <- true

			_, pkt, err := conn.ReadMessage()
			if err != nil {
				println(err.Error())
				return
			}
			println(`/`, pkt[0])
		})
		go http.ListenAndServe(Addr, nil)
	}
}

func Byte(key string) (chan<- byte, <-chan byte) {
	start.Do(run)

	w, r := make(chan byte), make(chan byte)

	bytel.Lock()
	B := wrbyte{
		key: key,
		idx: len(bytem[key]),
	}
	if !IsClient {
		B.n = make(chan int)
		B.w = make(chan []byte)
	}
	bytem[key] = append(bytem[key], B)
	bytel.Unlock()

	go func() {
		if IsClient {
			for b := range w {
				ws.Call("send", []byte{b})
			}
		}
		for b := range w {
			for ok := true; ok; ok = (len(B.n) > 0) {
				<-B.n
				B.w <- []byte{b}
			}
		}
	}()

	return w, r
}

func findWRByte(key string, idx int) (*wrbyte, bool) {
	bytel.Lock()
	defer bytel.Unlock()
	v, found := bytem[key]
	if !found {
		return nil, false
	}
	if idx >= len(v) {
		return nil, false
	}
	return &v[idx], true
}
