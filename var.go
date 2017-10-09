package sock

import (
	"sync"

	"github.com/gopherjs/gopherjs/js"
)

var (
	Addr         string
	IsClient     = js.Global != nil
	ClientFolder = "www"

	v = []byte(V)

	started sync.Once

	ws *js.Object

	byteDict struct {
		sync.RWMutex
		m map[string][]*tbyte
	}
)
