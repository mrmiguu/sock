package sock

import (
	"net/http"
	"os"
	"strings"

	"github.com/gopherjs/gopherjs/js"
	"github.com/gorilla/websocket"
	"github.com/mrmiguu/jsutil"
)

func run() {
	if IsClient {
		runClient()
	} else {
		runServer()
	}
}

func runClient() {
	wsOrWSS := "ws://"
	if strings.Contains(Addr, "https://") {
		wsOrWSS = "wss://"
	}
	Addr = strings.Replace(Addr, "http://", "", 1)
	Addr = strings.Replace(Addr, "https://", "", 1)

	wsync.Lock()
	defer wsync.Unlock()
	ws = js.Global.Get("WebSocket").New(wsOrWSS + Addr + API)
	ws.Set("binaryType", "arraybuffer")

	ws.Set("onclose", func() {
		go runClient()
	})

	f, c := jsutil.C()
	ws.Set("onopen", f)
	<-c

	ws.Set("onmessage", func(e *js.Object) {
		go func(pkt []byte) {
			err := read(pkt)
			if err != nil {
				wsync.Lock()
				ws.Call("close")
				wsync.Unlock()
			}
		}(js.Global.Get("Uint8Array").New(e.Get("data")).Interface().([]byte))
	})
}

func runServer() {
	if len(Root) > 0 {
		if _, err := os.Stat(Root); os.IsNotExist(err) {
			panic("root folder missing")
		}
		http.Handle("/", http.FileServer(http.Dir(Root)))
	}

	up := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}

	http.HandleFunc(API, func(w http.ResponseWriter, r *http.Request) {
		conn, err := up.Upgrade(w, r, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		connl.Lock()
		if len(conns) == 0 {
			reboot.Unlock()
		}
		conns[conn] = true
		connl.Unlock()
		defer func() {
			connl.Lock()
			conn.Close()
			delete(conns, conn)
			if len(conns) == 0 {
				reboot.Lock()
			}
			connl.Unlock()
		}()

		for {
			mt, pkt, err := conn.ReadMessage()
			if err != nil || mt != websocket.BinaryMessage || read(pkt) != nil {
				return
			}
		}
	})

	go http.ListenAndServe(Addr, nil)
}
