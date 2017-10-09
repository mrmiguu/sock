package sock

import (
	"bytes"
	"log"
	"net/http"
	"os"

	"github.com/gopherjs/gopherjs/js"
	"github.com/gorilla/websocket"
)

func init() {
	if IsClient {
		Addr = DefaultClientAddr
	} else {
		Addr = DefaultServerAddr
	}
}

func wAndOrRIfServer() {
	if IsClient {
		return
	}

	if _, err := os.Stat(ClientFolder); len(ClientFolder) > 0 && os.IsNotExist(err) {
		panic("client folder not found")
	}
	http.Handle("/", http.FileServer(http.Dir(ClientFolder)))

	up := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}
	http.HandleFunc("/"+SOCK, func(w http.ResponseWriter, r *http.Request) {
		conn, err := up.Upgrade(w, r, nil)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		defer conn.Close()

		for {
			_, pkt, err := conn.ReadMessage()
			if err != nil {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}

			parts := bytes.Split(pkt, v)
			t, name, idx := parts[0][0], string(parts[1]), bytes2int(parts[2])
			if len(parts) > 3 {
				go handlePOST(conn, t, name, idx, parts[3])
			} else {
				go handleGET(conn, t, name, idx)
			}
		}
	})

	log.Fatal(http.ListenAndServe(Addr, nil))
}

func handlePOST(conn *websocket.Conn, t byte, name string, idx int, body []byte) {
	switch t {
	case Tbyte:
		B, ok := findbyte(name, idx)
		if !ok {
			conn.Close()
			return
		}
		B.r <- body

	default:
		conn.Close()
		return
	}
}

func handleGET(conn *websocket.Conn, t byte, name string, idx int) {
	var b []byte

	switch t {
	case Tbyte:
		B, ok := findbyte(name, idx)
		if !ok {
			conn.Close()
			return
		}
		B.n <- 1
		b = <-B.w

	default:
		conn.Close()
		return
	}

	pkt := bytes.Join([][]byte{[]byte{t}, []byte(name), int2bytes(idx), b}, v)
	conn.WriteMessage(websocket.TextMessage, pkt)
}

func wAndOrR() {
	wAndOrRIfServer()
	wAndOrRIfClient()
}

func wAndOrRIfClient() {
	if !IsClient {
		return
	}

	if len(Addr) == 0 || Addr[len(Addr)-1] != '/' {
		Addr += "/"
	}

	ws = js.Global.Get("WebSocket").New("ws://" + Addr + SOCK)
	ws.Set("onmessage", func(e *js.Object) {
		pkt := []byte(e.Get("data").String())

		parts := bytes.Split(pkt, v)
		t, name, idx := parts[0][0], string(parts[1]), bytes2int(parts[2])

		if len(parts) > 3 {
			go handleSVRWR(t, name, idx, parts[3])
		} else {
			go handleSVRRD(t, name, idx)
		}
	})
}

func handleSVRWR(t byte, name string, idx int, body []byte) {
	switch t {
	case Tbyte:
		B, ok := findbyte(name, idx)
		if !ok {
			ws.Call("close")
			return
		}
		B.r <- body

	default:
		ws.Call("close")
	}
}

func handleSVRRD(t byte, name string, idx int) {
	var b []byte

	switch t {
	case Tbyte:
		B, ok := findbyte(name, idx)
		if !ok {
			ws.Call("close")
			return
		}
		b = <-B.w

	default:
		ws.Call("close")
		return
	}

	pkt := bytes.Join([][]byte{[]byte{t}, []byte(name), int2bytes(idx), b}, v)
	ws.Call("send", string(pkt))
}

// func wIfClient(w chan []byte, t byte, name string, idx int) {
// 	if !IsClient {
// 		return
// 	}

// 	for {
// 		pkt := bytes.Join([][]byte{[]byte{t}, []byte(name), int2bytes(idx), <-w}, v)

// 		// for {
// 		// 	resp, err := http.Post(Addr+SOCK, "text/plain", bytes.NewReader(pkt))
// 		// 	if err == nil && resp.StatusCode < 300 {
// 		// 		break
// 		// 	}
// 		// }
// 	}
// }

// func rIfClient(r chan []byte, t byte, name string, idx int) {
// 	if !IsClient {
// 		return
// 	}

// 	for {
// 		pkt := bytes.Join([][]byte{[]byte{t}, []byte(name), int2bytes(idx)}, v)
// 		for {
// 			resp, err := http.Post(Addr+SOCK, "text/plain", bytes.NewReader(pkt))
// 			if err != nil || resp.StatusCode > 299 {
// 				continue
// 			}
// 			b, err := ioutil.ReadAll(resp.Body)
// 			resp.Body.Close()
// 			if err == nil {
// 				r <- b
// 				break
// 			}
// 		}
// 	}
// }

// func delayedError(w http.ResponseWriter, code int) {
// 	time.Sleep(Timeout)
// 	http.Error(w, "", code)
// }
