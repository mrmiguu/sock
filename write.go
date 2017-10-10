package sock

import (
	"bytes"

	"github.com/gorilla/websocket"
)

func write(t byte, key string, idx int, body []byte) {
	pkt := bytes.Join([][]byte{[]byte{t}, []byte(key), int2bytes(idx), body}, v)

	if IsClient {
		ws.Call("send", pkt)
	} else {
		connl.RLock()
		for conn := range conns {
			conn.WriteMessage(websocket.BinaryMessage, pkt)
		}
		connl.RUnlock()
	}
}
