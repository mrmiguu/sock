package sock

import (
	"sync"

	"github.com/gopherjs/gopherjs/js"
	"github.com/gorilla/websocket"
)

var (
	Addr     = "localhost:80"
	IsClient = js.Global != nil
	Root     = "www"
	API      = "/abc123"
	v        = []byte(V)

	start sync.Once

	ws    *js.Object
	connl sync.RWMutex
	conns = map[*websocket.Conn]bool{}

	bytel sync.RWMutex
	bytem = map[string][]wrbyte{}
)
