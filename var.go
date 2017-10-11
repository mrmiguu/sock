package sock

import (
	"sync"

	"github.com/gopherjs/gopherjs/js"
	"github.com/gorilla/websocket"
)

var (
	v = []byte(V)

	start sync.Once

	ws     *js.Object
	connl  sync.RWMutex
	conns  = map[*websocket.Conn]bool{}
	reboot sync.RWMutex
	wsync  sync.Mutex

	wbooll    sync.Mutex
	wboolm    map[string][]interface{}
	rbooll    sync.RWMutex
	rboolm    map[string][]rbool
	wstringl  sync.Mutex
	wstringm  map[string][]interface{}
	rstringl  sync.RWMutex
	rstringm  map[string][]rstring
	wintl     sync.Mutex
	wintm     map[string][]interface{}
	rintl     sync.RWMutex
	rintm     map[string][]rint
	wint8l    sync.Mutex
	wint8m    map[string][]interface{}
	rint8l    sync.RWMutex
	rint8m    map[string][]rint8
	wint16l   sync.Mutex
	wint16m   map[string][]interface{}
	rint16l   sync.RWMutex
	rint16m   map[string][]rint16
	wint32l   sync.Mutex
	wint32m   map[string][]interface{}
	rint32l   sync.RWMutex
	rint32m   map[string][]rint32
	wint64l   sync.Mutex
	wint64m   map[string][]interface{}
	rint64l   sync.RWMutex
	rint64m   map[string][]rint64
	wuintl    sync.Mutex
	wuintm    map[string][]interface{}
	ruintl    sync.RWMutex
	ruintm    map[string][]ruint
	wuint8l   sync.Mutex
	wuint8m   map[string][]interface{}
	ruint8l   sync.RWMutex
	ruint8m   map[string][]ruint8
	wuint16l  sync.Mutex
	wuint16m  map[string][]interface{}
	ruint16l  sync.RWMutex
	ruint16m  map[string][]ruint16
	wuint32l  sync.Mutex
	wuint32m  map[string][]interface{}
	ruint32l  sync.RWMutex
	ruint32m  map[string][]ruint32
	wuint64l  sync.Mutex
	wuint64m  map[string][]interface{}
	ruint64l  sync.RWMutex
	ruint64m  map[string][]ruint64
	wbytel    sync.Mutex
	wbytem    map[string][]interface{}
	rbytel    sync.RWMutex
	rbytem    map[string][]rbyte
	wbytesl   sync.Mutex
	wbytesm   map[string][]interface{}
	rbytesl   sync.RWMutex
	rbytesm   map[string][]rbytes
	wrunel    sync.Mutex
	wrunem    map[string][]interface{}
	rrunel    sync.RWMutex
	rrunem    map[string][]rrune
	wfloat32l sync.Mutex
	wfloat32m map[string][]interface{}
	rfloat32l sync.RWMutex
	rfloat32m map[string][]rfloat32
	wfloat64l sync.Mutex
	wfloat64m map[string][]interface{}
	rfloat64l sync.RWMutex
	rfloat64m map[string][]rfloat64
	werrorl   sync.Mutex
	werrorm   map[string][]interface{}
	rerrorl   sync.RWMutex
	rerrorm   map[string][]rerror
)

func init() {
	reboot.Lock()
}
