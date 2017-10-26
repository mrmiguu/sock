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
	wboolm    = make(map[string]wboolx)
	rbooll    sync.RWMutex
	rboolm    = make(map[string]rboolx)
	wstringl  sync.Mutex
	wstringm  = make(map[string]wstringx)
	rstringl  sync.RWMutex
	rstringm  = make(map[string]rstringx)
	wintl     sync.Mutex
	wintm     = make(map[string]wintx)
	rintl     sync.RWMutex
	rintm     = make(map[string]rintx)
	wint8l    sync.Mutex
	wint8m    = make(map[string]wint8x)
	rint8l    sync.RWMutex
	rint8m    = make(map[string]rint8x)
	wint16l   sync.Mutex
	wint16m   = make(map[string]wint16x)
	rint16l   sync.RWMutex
	rint16m   = make(map[string]rint16x)
	wint32l   sync.Mutex
	wint32m   = make(map[string]wint32x)
	rint32l   sync.RWMutex
	rint32m   = make(map[string]rint32x)
	wint64l   sync.Mutex
	wint64m   = make(map[string]wint64x)
	rint64l   sync.RWMutex
	rint64m   = make(map[string]rint64x)
	wuintl    sync.Mutex
	wuintm    = make(map[string]wuintx)
	ruintl    sync.RWMutex
	ruintm    = make(map[string]ruintx)
	wuint8l   sync.Mutex
	wuint8m   = make(map[string]wuint8x)
	ruint8l   sync.RWMutex
	ruint8m   = make(map[string]ruint8x)
	wuint16l  sync.Mutex
	wuint16m  = make(map[string]wuint16x)
	ruint16l  sync.RWMutex
	ruint16m  = make(map[string]ruint16x)
	wuint32l  sync.Mutex
	wuint32m  = make(map[string]wuint32x)
	ruint32l  sync.RWMutex
	ruint32m  = make(map[string]ruint32x)
	wuint64l  sync.Mutex
	wuint64m  = make(map[string]wuint64x)
	ruint64l  sync.RWMutex
	ruint64m  = make(map[string]ruint64x)
	wbytel    sync.Mutex
	wbytem    = make(map[string]wbytex)
	rbytel    sync.RWMutex
	rbytem    = make(map[string]rbytex)
	wbytesl   sync.Mutex
	wbytesm   = make(map[string]wbytesx)
	rbytesl   sync.RWMutex
	rbytesm   = make(map[string]rbytesx)
	wrunel    sync.Mutex
	wrunem    = make(map[string]wrunex)
	rrunel    sync.RWMutex
	rrunem    = make(map[string]rrunex)
	wfloat32l sync.Mutex
	wfloat32m = make(map[string]wfloat32x)
	rfloat32l sync.RWMutex
	rfloat32m = make(map[string]rfloat32x)
	wfloat64l sync.Mutex
	wfloat64m = make(map[string]wfloat64x)
	rfloat64l sync.RWMutex
	rfloat64m = make(map[string]rfloat64x)
	werrorl   sync.Mutex
	werrorm   = make(map[string]werrorx)
	rerrorl   sync.RWMutex
	rerrorm   = make(map[string]rerrorx)
)

func init() {
	reboot.Lock()
}
