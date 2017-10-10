package sock

import (
	"sync"

	"github.com/gopherjs/gopherjs/js"
	"github.com/gorilla/websocket"
)

var (
	v = []byte(V)

	start sync.Once

	ws    *js.Object
	connl sync.RWMutex
	conns = map[*websocket.Conn]bool{}

	booll    sync.RWMutex
	boolm    = map[string][]wrbool{}
	stringl  sync.RWMutex
	stringm  = map[string][]wrstring{}
	intl     sync.RWMutex
	intm     = map[string][]wrint{}
	int8l    sync.RWMutex
	int8m    = map[string][]wrint8{}
	int16l   sync.RWMutex
	int16m   = map[string][]wrint16{}
	int32l   sync.RWMutex
	int32m   = map[string][]wrint32{}
	int64l   sync.RWMutex
	int64m   = map[string][]wrint64{}
	uintl    sync.RWMutex
	uintm    = map[string][]wruint{}
	uint8l   sync.RWMutex
	uint8m   = map[string][]wruint8{}
	uint16l  sync.RWMutex
	uint16m  = map[string][]wruint16{}
	uint32l  sync.RWMutex
	uint32m  = map[string][]wruint32{}
	uint64l  sync.RWMutex
	uint64m  = map[string][]wruint64{}
	bytel    sync.RWMutex
	bytem    = map[string][]wrbyte{}
	bytesl   sync.RWMutex
	bytesm   = map[string][]wrbytes{}
	runel    sync.RWMutex
	runem    = map[string][]wrrune{}
	float32l sync.RWMutex
	float32m = map[string][]wrfloat32{}
	float64l sync.RWMutex
	float64m = map[string][]wrfloat64{}
	errorl   sync.RWMutex
	errorm   = map[string][]wrerror{}
)
