package sock

import (
	"sync"
	"time"

	"github.com/gopherjs/gopherjs/js"
)

var (
	Addr       string
	IsClient   = js.Global != nil
	ErrorDelay = 2 * time.Second

	v = []byte(V)

	started sync.Once

	errorDict struct {
		sync.RWMutex
		m map[string][]*terror
	}
	boolDict struct {
		sync.RWMutex
		m map[string][]*tbool
	}
	stringDict struct {
		sync.RWMutex
		m map[string][]*tstring
	}
	intDict struct {
		sync.RWMutex
		m map[string][]*tint
	}
	// int8Dict struct {
	// 	sync.RWMutex
	// 	m map[string][]*tint8
	// }
	// int16Dict struct {
	// 	sync.RWMutex
	// 	m map[string][]*tint16
	// }
	// int32Dict struct {
	// 	sync.RWMutex
	// 	m map[string][]*tint32
	// }
	// int64Dict struct {
	// 	sync.RWMutex
	// 	m map[string][]*tint64
	// }
	// uintDict struct {
	// 	sync.RWMutex
	// 	m map[string][]*tuint
	// }
	// uint8Dict struct {
	// 	sync.RWMutex
	// 	m map[string][]*tuint8
	// }
	// uint16Dict struct {
	// 	sync.RWMutex
	// 	m map[string][]*tuint16
	// }
	// uint32Dict struct {
	// 	sync.RWMutex
	// 	m map[string][]*tuint32
	// }
	// uint64Dict struct {
	// 	sync.RWMutex
	// 	m map[string][]*tuint64
	// }
	// uintptrDict struct {
	// 	sync.RWMutex
	// 	m map[string][]*tuintptr
	// }
	// byteDict struct {
	// 	sync.RWMutex
	// 	m map[string][]*tbyte
	// }
	bytesDict struct {
		sync.RWMutex
		m map[string][]*tbytes
	}
	runeDict struct {
		sync.RWMutex
		m map[string][]*trune
	}
	// float32Dict struct {
	// 	sync.RWMutex
	// 	m map[string][]*tfloat32
	// }
	float64Dict struct {
		sync.RWMutex
		m map[string][]*tfloat64
	}
	// complex64Dict struct {
	// 	sync.RWMutex
	// 	m map[string][]*tcomplex64
	// }
	// complex128Dict struct {
	// 	sync.RWMutex
	// 	m map[string][]*tcomplex128
	// }
)
