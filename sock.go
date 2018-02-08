package sock

import "github.com/gopherjs/gopherjs/js"

var (
	Addr = "localhost:80"

	IsClient = js.Global != nil

	Secure = false

	Root = "www"

	Route = "/317d37b0edc7bd7cbd25d97f53a16ce5"
)

func Wbool(key ...string) chan<- bool {
	k := getKey(key...)

	start.Do(run)

	w := make(chan bool)

	wbooll.Lock()
	wb := wboolm[k]
	var B wbool
	B.key = k
	B.idx = wb.i
	B.w = w
	wb.sl = append(wb.sl, B)
	wb.i++
	wboolm[k] = wb
	wbooll.Unlock()

	go func() {
		for b := range w {
			write(Tbool, k, B.idx, bool2bytes(b))
		}
	}()

	return w
}

func Rbool(key ...string) <-chan bool {
	k := getKey(key...)

	start.Do(run)

	r := make(chan bool)

	rbooll.Lock()
	rb := rboolm[k]
	var B rbool
	B.key = k
	B.idx = rb.i
	B.r = r
	rb.sl = append(rb.sl, B)
	rb.i++
	rboolm[k] = rb
	rbooll.Unlock()

	return r
}

func Wstring(key ...string) chan<- string {
	k := getKey(key...)

	start.Do(run)

	w := make(chan string)

	wstringl.Lock()
	ws := wstringm[k]
	var S wstring
	S.key = k
	S.idx = ws.i
	S.w = w
	ws.sl = append(ws.sl, S)
	ws.i++
	wstringm[k] = ws
	wstringl.Unlock()

	go func() {
		for s := range w {
			write(Tstring, k, S.idx, []byte(s))
		}
	}()

	return w
}

func Rstring(key ...string) <-chan string {
	k := getKey(key...)

	start.Do(run)

	r := make(chan string)

	rstringl.Lock()
	rs := rstringm[k]
	var S rstring
	S.key = k
	S.idx = rs.i
	S.r = r
	rs.sl = append(rs.sl, S)
	rs.i++
	rstringm[k] = rs
	rstringl.Unlock()

	return r
}

func Wint(key ...string) chan<- int {
	k := getKey(key...)

	start.Do(run)

	w := make(chan int)

	wintl.Lock()
	wi := wintm[k]
	var I wint
	I.key = k
	I.idx = wi.i
	I.w = w
	wi.sl = append(wi.sl, I)
	wi.i++
	wintm[k] = wi
	wintl.Unlock()

	go func() {
		for i := range w {
			write(Tint, k, I.idx, int2bytes(i))
		}
	}()

	return w
}

func Rint(key ...string) <-chan int {
	k := getKey(key...)

	start.Do(run)

	r := make(chan int)

	rintl.Lock()
	ri := rintm[k]
	var I rint
	I.key = k
	I.idx = ri.i
	I.r = r
	ri.sl = append(ri.sl, I)
	ri.i++
	rintm[k] = ri
	rintl.Unlock()

	return r
}

func Wint8(key ...string) chan<- int8 {
	k := getKey(key...)

	start.Do(run)

	w := make(chan int8)

	wint8l.Lock()
	wi := wint8m[k]
	var I wint8
	I.key = k
	I.idx = wi.i
	I.w = w
	wi.sl = append(wi.sl, I)
	wi.i++
	wint8m[k] = wi
	wint8l.Unlock()

	go func() {
		for i := range w {
			write(Tint8, k, I.idx, int82bytes(i))
		}
	}()

	return w
}

func Rint8(key ...string) <-chan int8 {
	k := getKey(key...)

	start.Do(run)

	r := make(chan int8)

	rint8l.Lock()
	ri := rint8m[k]
	var I rint8
	I.key = k
	I.idx = ri.i
	I.r = r
	ri.sl = append(ri.sl, I)
	ri.i++
	rint8m[k] = ri
	rint8l.Unlock()

	return r
}

func Wint16(key ...string) chan<- int16 {
	k := getKey(key...)

	start.Do(run)

	w := make(chan int16)

	wint16l.Lock()
	wi := wint16m[k]
	var I wint16
	I.key = k
	I.idx = wi.i
	I.w = w
	wi.sl = append(wi.sl, I)
	wi.i++
	wint16m[k] = wi
	wint16l.Unlock()

	go func() {
		for i := range w {
			write(Tint16, k, I.idx, int162bytes(i))
		}
	}()

	return w
}

func Rint16(key ...string) <-chan int16 {
	k := getKey(key...)

	start.Do(run)

	r := make(chan int16)

	rint16l.Lock()
	ri := rint16m[k]
	var I rint16
	I.key = k
	I.idx = ri.i
	I.r = r
	ri.sl = append(ri.sl, I)
	ri.i++
	rint16m[k] = ri
	rint16l.Unlock()

	return r
}

func Wint32(key ...string) chan<- int32 {
	k := getKey(key...)

	start.Do(run)

	w := make(chan int32)

	wint32l.Lock()
	wi := wint32m[k]
	var I wint32
	I.key = k
	I.idx = wi.i
	I.w = w
	wi.sl = append(wi.sl, I)
	wi.i++
	wint32m[k] = wi
	wint32l.Unlock()

	go func() {
		for i := range w {
			write(Tint32, k, I.idx, int322bytes(i))
		}
	}()

	return w
}

func Rint32(key ...string) <-chan int32 {
	k := getKey(key...)

	start.Do(run)

	r := make(chan int32)

	rint32l.Lock()
	ri := rint32m[k]
	var I rint32
	I.key = k
	I.idx = ri.i
	I.r = r
	ri.sl = append(ri.sl, I)
	ri.i++
	rint32m[k] = ri
	rint32l.Unlock()

	return r
}

func Wint64(key ...string) chan<- int64 {
	k := getKey(key...)

	start.Do(run)

	w := make(chan int64)

	wint64l.Lock()
	wi := wint64m[k]
	var I wint64
	I.key = k
	I.idx = wi.i
	I.w = w
	wi.sl = append(wi.sl, I)
	wi.i++
	wint64m[k] = wi
	wint64l.Unlock()

	go func() {
		for i := range w {
			write(Tint64, k, I.idx, int642bytes(i))
		}
	}()

	return w
}

func Rint64(key ...string) <-chan int64 {
	k := getKey(key...)

	start.Do(run)

	r := make(chan int64)

	rint64l.Lock()
	ri := rint64m[k]
	var I rint64
	I.key = k
	I.idx = ri.i
	I.r = r
	ri.sl = append(ri.sl, I)
	ri.i++
	rint64m[k] = ri
	rint64l.Unlock()

	return r
}

func Wuint(key ...string) chan<- uint {
	k := getKey(key...)

	start.Do(run)

	w := make(chan uint)

	wuintl.Lock()
	wu := wuintm[k]
	var U wuint
	U.key = k
	U.idx = wu.i
	U.w = w
	wu.sl = append(wu.sl, U)
	wu.i++
	wuintm[k] = wu
	wuintl.Unlock()

	go func() {
		for u := range w {
			write(Tuint, k, U.idx, uint2bytes(u))
		}
	}()

	return w
}

func Ruint(key ...string) <-chan uint {
	k := getKey(key...)

	start.Do(run)

	r := make(chan uint)

	ruintl.Lock()
	ru := ruintm[k]
	var U ruint
	U.key = k
	U.idx = ru.i
	U.r = r
	ru.sl = append(ru.sl, U)
	ru.i++
	ruintm[k] = ru
	ruintl.Unlock()

	return r
}

func Wuint8(key ...string) chan<- uint8 {
	k := getKey(key...)

	start.Do(run)

	w := make(chan uint8)

	wuint8l.Lock()
	wu := wuint8m[k]
	var U wuint8
	U.key = k
	U.idx = wu.i
	U.w = w
	wu.sl = append(wu.sl, U)
	wu.i++
	wuint8m[k] = wu
	wuint8l.Unlock()

	go func() {
		for u := range w {
			write(Tuint8, k, U.idx, uint82bytes(u))
		}
	}()

	return w
}

func Ruint8(key ...string) <-chan uint8 {
	k := getKey(key...)

	start.Do(run)

	r := make(chan uint8)

	ruint8l.Lock()
	ru := ruint8m[k]
	var U ruint8
	U.key = k
	U.idx = ru.i
	U.r = r
	ru.sl = append(ru.sl, U)
	ru.i++
	ruint8m[k] = ru
	ruint8l.Unlock()

	return r
}

func Wuint16(key ...string) chan<- uint16 {
	k := getKey(key...)

	start.Do(run)

	w := make(chan uint16)

	wuint16l.Lock()
	wu := wuint16m[k]
	var U wuint16
	U.key = k
	U.idx = wu.i
	U.w = w
	wu.sl = append(wu.sl, U)
	wu.i++
	wuint16m[k] = wu
	wuint16l.Unlock()

	go func() {
		for u := range w {
			write(Tuint16, k, U.idx, uint162bytes(u))
		}
	}()

	return w
}

func Ruint16(key ...string) <-chan uint16 {
	k := getKey(key...)

	start.Do(run)

	r := make(chan uint16)

	ruint16l.Lock()
	ru := ruint16m[k]
	var U ruint16
	U.key = k
	U.idx = ru.i
	U.r = r
	ru.sl = append(ru.sl, U)
	ru.i++
	ruint16m[k] = ru
	ruint16l.Unlock()

	return r
}

func Wuint32(key ...string) chan<- uint32 {
	k := getKey(key...)

	start.Do(run)

	w := make(chan uint32)

	wuint32l.Lock()
	wu := wuint32m[k]
	var U wuint32
	U.key = k
	U.idx = wu.i
	U.w = w
	wu.sl = append(wu.sl, U)
	wu.i++
	wuint32m[k] = wu
	wuint32l.Unlock()

	go func() {
		for u := range w {
			write(Tuint32, k, U.idx, uint322bytes(u))
		}
	}()

	return w
}

func Ruint32(key ...string) <-chan uint32 {
	k := getKey(key...)

	start.Do(run)

	r := make(chan uint32)

	ruint32l.Lock()
	ru := ruint32m[k]
	var U ruint32
	U.key = k
	U.idx = ru.i
	U.r = r
	ru.sl = append(ru.sl, U)
	ru.i++
	ruint32m[k] = ru
	ruint32l.Unlock()

	return r
}

func Wuint64(key ...string) chan<- uint64 {
	k := getKey(key...)

	start.Do(run)

	w := make(chan uint64)

	wuint64l.Lock()
	wu := wuint64m[k]
	var U wuint64
	U.key = k
	U.idx = wu.i
	U.w = w
	wu.sl = append(wu.sl, U)
	wu.i++
	wuint64m[k] = wu
	wuint64l.Unlock()

	go func() {
		for u := range w {
			write(Tuint64, k, U.idx, uint642bytes(u))
		}
	}()

	return w
}

func Ruint64(key ...string) <-chan uint64 {
	k := getKey(key...)

	start.Do(run)

	r := make(chan uint64)

	ruint64l.Lock()
	ru := ruint64m[k]
	var U ruint64
	U.key = k
	U.idx = ru.i
	U.r = r
	ru.sl = append(ru.sl, U)
	ru.i++
	ruint64m[k] = ru
	ruint64l.Unlock()

	return r
}

func Wbyte(key ...string) chan<- byte {
	k := getKey(key...)

	start.Do(run)

	w := make(chan byte)

	wbytel.Lock()
	wb := wbytem[k]
	var B wbyte
	B.key = k
	B.idx = wb.i
	B.w = w
	wb.sl = append(wb.sl, B)
	wb.i++
	wbytem[k] = wb
	wbytel.Unlock()

	go func() {
		for b := range w {
			write(Tbyte, k, B.idx, []byte{b})
		}
	}()

	return w
}

func Rbyte(key ...string) <-chan byte {
	k := getKey(key...)

	start.Do(run)

	r := make(chan byte)

	rbytel.Lock()
	rb := rbytem[k]
	var B rbyte
	B.key = k
	B.idx = rb.i
	B.r = r
	rb.sl = append(rb.sl, B)
	rb.i++
	rbytem[k] = rb
	rbytel.Unlock()

	return r
}

func Wbytes(key ...string) chan<- []byte {
	k := getKey(key...)

	start.Do(run)

	w := make(chan []byte)

	wbytesl.Lock()
	wb := wbytesm[k]
	var B wbytes
	B.key = k
	B.idx = wb.i
	B.w = w
	wb.sl = append(wb.sl, B)
	wb.i++
	wbytesm[k] = wb
	wbytesl.Unlock()

	go func() {
		for b := range w {
			write(Tbytes, k, B.idx, b)
		}
	}()

	return w
}

func Rbytes(key ...string) <-chan []byte {
	k := getKey(key...)

	start.Do(run)

	r := make(chan []byte)

	rbytesl.Lock()
	rb := rbytesm[k]
	var B rbytes
	B.key = k
	B.idx = rb.i
	B.r = r
	rb.sl = append(rb.sl, B)
	rb.i++
	rbytesm[k] = rb
	rbytesl.Unlock()

	return r
}

func Wrune(key ...string) chan<- rune {
	k := getKey(key...)

	start.Do(run)

	w := make(chan rune)

	wrunel.Lock()
	wr := wrunem[k]
	var R wrune
	R.key = k
	R.idx = wr.i
	R.w = w
	wr.sl = append(wr.sl, R)
	wr.i++
	wrunem[k] = wr
	wrunel.Unlock()

	go func() {
		for r := range w {
			write(Trune, k, R.idx, rune2bytes(r))
		}
	}()

	return w
}

func Rrune(key ...string) <-chan rune {
	k := getKey(key...)

	start.Do(run)

	r := make(chan rune)

	rrunel.Lock()
	rr := rrunem[k]
	var R rrune
	R.key = k
	R.idx = rr.i
	R.r = r
	rr.sl = append(rr.sl, R)
	rr.i++
	rrunem[k] = rr
	rrunel.Unlock()

	return r
}

func Wfloat32(key ...string) chan<- float32 {
	k := getKey(key...)

	start.Do(run)

	w := make(chan float32)

	wfloat32l.Lock()
	wf := wfloat32m[k]
	var F wfloat32
	F.key = k
	F.idx = wf.i
	F.w = w
	wf.sl = append(wf.sl, F)
	wf.i++
	wfloat32m[k] = wf
	wfloat32l.Unlock()

	go func() {
		for f := range w {
			write(Tfloat32, k, F.idx, float322bytes(f))
		}
	}()

	return w
}

func Rfloat32(key ...string) <-chan float32 {
	k := getKey(key...)

	start.Do(run)

	r := make(chan float32)

	rfloat32l.Lock()
	rf := rfloat32m[k]
	var F rfloat32
	F.key = k
	F.idx = rf.i
	F.r = r
	rf.sl = append(rf.sl, F)
	rf.i++
	rfloat32m[k] = rf
	rfloat32l.Unlock()

	return r
}

func Wfloat64(key ...string) chan<- float64 {
	k := getKey(key...)

	start.Do(run)

	w := make(chan float64)

	wfloat64l.Lock()
	wf := wfloat64m[k]
	var F wfloat64
	F.key = k
	F.idx = wf.i
	F.w = w
	wf.sl = append(wf.sl, F)
	wf.i++
	wfloat64m[k] = wf
	wfloat64l.Unlock()

	go func() {
		for f := range w {
			write(Tfloat64, k, F.idx, float642bytes(f))
		}
	}()

	return w
}

func Rfloat64(key ...string) <-chan float64 {
	k := getKey(key...)

	start.Do(run)

	r := make(chan float64)

	rfloat64l.Lock()
	rf := rfloat64m[k]
	var F rfloat64
	F.key = k
	F.idx = rf.i
	F.r = r
	rf.sl = append(rf.sl, F)
	rf.i++
	rfloat64m[k] = rf
	rfloat64l.Unlock()

	return r
}

func Werror(key ...string) chan<- error {
	k := getKey(key...)

	start.Do(run)

	w := make(chan error)

	werrorl.Lock()
	we := werrorm[k]
	var E werror
	E.key = k
	E.idx = we.i
	E.w = w
	we.sl = append(we.sl, E)
	we.i++
	werrorm[k] = we
	werrorl.Unlock()

	go func() {
		for e := range w {
			write(Terror, k, E.idx, error2bytes(e))
		}
	}()

	return w
}

func Rerror(key ...string) <-chan error {
	k := getKey(key...)

	start.Do(run)

	r := make(chan error)

	rerrorl.Lock()
	re := rerrorm[k]
	var E rerror
	E.key = k
	E.idx = re.i
	E.r = r
	re.sl = append(re.sl, E)
	re.i++
	rerrorm[k] = re
	rerrorl.Unlock()

	return r
}

func getKey(key ...string) string {
	if len(key) == 0 {
		return ""
	}
	if len(key) > 1 {
		panic("too many arguments")
	}
	if len(key[0]) == 0 {
		panic("empty key")
	}
	return key[0]
}
