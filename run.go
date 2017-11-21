package sock

import (
	"mime"
	"net/http"
	"os"
	"path/filepath"
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
	wsOrWSS := "wss://"
	if !Secure {
		wsOrWSS = "ws://"
	}

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
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			// var p []string

			ext := filepath.Ext(r.URL.Path)
			if ext == ".gz" {
				// p = append(p, "Content-Encoding: gzip")
				w.Header().Add("Content-Encoding", "gzip")
			}

			idx := strings.LastIndex(r.URL.Path, ext)
			if idx != -1 {
				ext = filepath.Ext(r.URL.Path[:idx])
				if t := mime.TypeByExtension(ext); len(t) > 0 {
					// p = append(p, "Content-Type: "+t)
					w.Header().Add("Content-Type", t)
				}
			}

			http.ServeFile(w, r, Root+r.URL.Path)

			// println(strings.Join(p, ", "))
		})
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
