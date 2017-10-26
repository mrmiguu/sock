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
	wboolm    = make(map[string][]wbool)
	rbooll    sync.RWMutex
	rboolm    = make(map[string][]rbool)
	wstringl  sync.Mutex
	wstringm  = make(map[string][]wstring)
	rstringl  sync.RWMutex
	rstringm  = make(map[string][]rstring)
	wintl     sync.Mutex
	wintm     = make(map[string][]wint)
	rintl     sync.RWMutex
	rintm     = make(map[string][]rint)
	wint8l    sync.Mutex
	wint8m    = make(map[string][]wint8)
	rint8l    sync.RWMutex
	rint8m    = make(map[string][]rint8)
	wint16l   sync.Mutex
	wint16m   = make(map[string][]wint16)
	rint16l   sync.RWMutex
	rint16m   = make(map[string][]rint16)
	wint32l   sync.Mutex
	wint32m   = make(map[string][]wint32)
	rint32l   sync.RWMutex
	rint32m   = make(map[string][]rint32)
	wint64l   sync.Mutex
	wint64m   = make(map[string][]wint64)
	rint64l   sync.RWMutex
	rint64m   = make(map[string][]rint64)
	wuintl    sync.Mutex
	wuintm    = make(map[string][]wuint)
	ruintl    sync.RWMutex
	ruintm    = make(map[string][]ruint)
	wuint8l   sync.Mutex
	wuint8m   = make(map[string][]wuint8)
	ruint8l   sync.RWMutex
	ruint8m   = make(map[string][]ruint8)
	wuint16l  sync.Mutex
	wuint16m  = make(map[string][]wuint16)
	ruint16l  sync.RWMutex
	ruint16m  = make(map[string][]ruint16)
	wuint32l  sync.Mutex
	wuint32m  = make(map[string][]wuint32)
	ruint32l  sync.RWMutex
	ruint32m  = make(map[string][]ruint32)
	wuint64l  sync.Mutex
	wuint64m  = make(map[string][]wuint64)
	ruint64l  sync.RWMutex
	ruint64m  = make(map[string][]ruint64)
	wbytel    sync.Mutex
	wbytem    = make(map[string][]wbyte)
	rbytel    sync.RWMutex
	rbytem    = make(map[string][]rbyte)
	wbytesl   sync.Mutex
	wbytesm   = make(map[string][]wbytes)
	rbytesl   sync.RWMutex
	rbytesm   = make(map[string][]rbytes)
	wrunel    sync.Mutex
	wrunem    = make(map[string][]wrune)
	rrunel    sync.RWMutex
	rrunem    = make(map[string][]rrune)
	wfloat32l sync.Mutex
	wfloat32m = make(map[string][]wfloat32)
	rfloat32l sync.RWMutex
	rfloat32m = make(map[string][]rfloat32)
	wfloat64l sync.Mutex
	wfloat64m = make(map[string][]wfloat64)
	rfloat64l sync.RWMutex
	rfloat64m = make(map[string][]rfloat64)
	werrorl   sync.Mutex
	werrorm   = make(map[string][]werror)
	rerrorl   sync.RWMutex
	rerrorm   = make(map[string][]rerror)
)

func init() {
	reboot.Lock()
}
