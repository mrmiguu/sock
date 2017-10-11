package sock

import (
	"bytes"

	"github.com/gorilla/websocket"
)

func write(t byte, key string, idx int, body []byte) {
	pkt := bytes.Join([][]byte{[]byte{t}, []byte(key), int2bytes(idx), body}, v)

	if IsClient {
		ws.Call("send", pkt)
		return
	}
	for {
		connl.RLock()
		cnt := len(conns)
		for conn := range conns {
			conn.WriteMessage(websocket.BinaryMessage, pkt)
		}
		connl.RUnlock()
		if cnt > 0 {
			break
		}
		reboot.RLock()
		reboot.RUnlock()
	}
}
