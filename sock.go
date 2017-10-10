package sock

import "github.com/gopherjs/gopherjs/js"

var (
	Addr     = "localhost:80"
	IsClient = js.Global != nil
	Root     = "www"
	API      = "/317d37b0edc7bd7cbd25d97f53a16ce5"
)

func MakeBool(key ...string) (chan<- bool, <-chan bool) {
	k := getKey(key...)

	start.Do(run)

	w, r := make(chan bool), make(chan bool)

	booll.Lock()
	var B wrbool
	B.key = k
	B.idx = len(boolm[k])
	B.r = r
	boolm[k] = append(boolm[k], B)
	booll.Unlock()

	go func() {
		for b := range w {
			write(Tbool, B.key, B.idx, bool2bytes(b))
		}
	}()

	return w, r
}

func MakeString(key ...string) (chan<- string, <-chan string) {
	k := getKey(key...)

	start.Do(run)

	w, r := make(chan string), make(chan string)

	stringl.Lock()
	var S wrstring
	S.key = k
	S.idx = len(stringm[k])
	S.r = r
	stringm[k] = append(stringm[k], S)
	stringl.Unlock()

	go func() {
		for s := range w {
			write(Tstring, S.key, S.idx, []byte(s))
		}
	}()

	return w, r
}

func MakeInt(key ...string) (chan<- int, <-chan int) {
	k := getKey(key...)

	start.Do(run)

	w, r := make(chan int), make(chan int)

	intl.Lock()
	var I wrint
	I.key = k
	I.idx = len(intm[k])
	I.r = r
	intm[k] = append(intm[k], I)
	intl.Unlock()

	go func() {
		for i := range w {
			write(Tint, I.key, I.idx, int2bytes(i))
		}
	}()

	return w, r
}

func MakeInt8(key ...string) (chan<- int8, <-chan int8) {
	k := getKey(key...)

	start.Do(run)

	w, r := make(chan int8), make(chan int8)

	int8l.Lock()
	var I wrint8
	I.key = k
	I.idx = len(int8m[k])
	I.r = r
	int8m[k] = append(int8m[k], I)
	int8l.Unlock()

	go func() {
		for i := range w {
			write(Tint8, I.key, I.idx, int82bytes(i))
		}
	}()

	return w, r
}

func MakeInt16(key ...string) (chan<- int16, <-chan int16) {
	k := getKey(key...)

	start.Do(run)

	w, r := make(chan int16), make(chan int16)

	int16l.Lock()
	var I wrint16
	I.key = k
	I.idx = len(int16m[k])
	I.r = r
	int16m[k] = append(int16m[k], I)
	int16l.Unlock()

	go func() {
		for i := range w {
			write(Tint16, I.key, I.idx, int162bytes(i))
		}
	}()

	return w, r
}

func MakeInt32(key ...string) (chan<- int32, <-chan int32) {
	k := getKey(key...)

	start.Do(run)

	w, r := make(chan int32), make(chan int32)

	int32l.Lock()
	var I wrint32
	I.key = k
	I.idx = len(int32m[k])
	I.r = r
	int32m[k] = append(int32m[k], I)
	int32l.Unlock()

	go func() {
		for i := range w {
			write(Tint32, I.key, I.idx, int322bytes(i))
		}
	}()

	return w, r
}

func MakeInt64(key ...string) (chan<- int64, <-chan int64) {
	k := getKey(key...)

	start.Do(run)

	w, r := make(chan int64), make(chan int64)

	int64l.Lock()
	var I wrint64
	I.key = k
	I.idx = len(int64m[k])
	I.r = r
	int64m[k] = append(int64m[k], I)
	int64l.Unlock()

	go func() {
		for i := range w {
			write(Tint64, I.key, I.idx, int642bytes(i))
		}
	}()

	return w, r
}

func MakeUint(key ...string) (chan<- uint, <-chan uint) {
	k := getKey(key...)

	start.Do(run)

	w, r := make(chan uint), make(chan uint)

	uintl.Lock()
	var U wruint
	U.key = k
	U.idx = len(uintm[k])
	U.r = r
	uintm[k] = append(uintm[k], U)
	uintl.Unlock()

	go func() {
		for u := range w {
			write(Tuint, U.key, U.idx, uint2bytes(u))
		}
	}()

	return w, r
}

func MakeUint8(key ...string) (chan<- uint8, <-chan uint8) {
	k := getKey(key...)

	start.Do(run)

	w, r := make(chan uint8), make(chan uint8)

	uint8l.Lock()
	var U wruint8
	U.key = k
	U.idx = len(uint8m[k])
	U.r = r
	uint8m[k] = append(uint8m[k], U)
	uint8l.Unlock()

	go func() {
		for u := range w {
			write(Tuint8, U.key, U.idx, uint82bytes(u))
		}
	}()

	return w, r
}

func MakeUint16(key ...string) (chan<- uint16, <-chan uint16) {
	k := getKey(key...)

	start.Do(run)

	w, r := make(chan uint16), make(chan uint16)

	uint16l.Lock()
	var U wruint16
	U.key = k
	U.idx = len(uint16m[k])
	U.r = r
	uint16m[k] = append(uint16m[k], U)
	uint16l.Unlock()

	go func() {
		for u := range w {
			write(Tuint16, U.key, U.idx, uint162bytes(u))
		}
	}()

	return w, r
}

func MakeUint32(key ...string) (chan<- uint32, <-chan uint32) {
	k := getKey(key...)

	start.Do(run)

	w, r := make(chan uint32), make(chan uint32)

	uint32l.Lock()
	var U wruint32
	U.key = k
	U.idx = len(uint32m[k])
	U.r = r
	uint32m[k] = append(uint32m[k], U)
	uint32l.Unlock()

	go func() {
		for u := range w {
			write(Tuint32, U.key, U.idx, uint322bytes(u))
		}
	}()

	return w, r
}

func MakeUint64(key ...string) (chan<- uint64, <-chan uint64) {
	k := getKey(key...)

	start.Do(run)

	w, r := make(chan uint64), make(chan uint64)

	uint64l.Lock()
	var U wruint64
	U.key = k
	U.idx = len(uint64m[k])
	U.r = r
	uint64m[k] = append(uint64m[k], U)
	uint64l.Unlock()

	go func() {
		for u := range w {
			write(Tuint64, U.key, U.idx, uint642bytes(u))
		}
	}()

	return w, r
}

func MakeByte(key ...string) (chan<- byte, <-chan byte) {
	k := getKey(key...)

	start.Do(run)

	w, r := make(chan byte), make(chan byte)

	bytel.Lock()
	var B wrbyte
	B.key = k
	B.idx = len(bytem[k])
	B.r = r
	bytem[k] = append(bytem[k], B)
	bytel.Unlock()

	go func() {
		for b := range w {
			write(Tbyte, B.key, B.idx, []byte{b})
		}
	}()

	return w, r
}

func MakeBytes(key ...string) (chan<- []byte, <-chan []byte) {
	k := getKey(key...)

	start.Do(run)

	w, r := make(chan []byte), make(chan []byte)

	bytesl.Lock()
	var B wrbytes
	B.key = k
	B.idx = len(bytesm[k])
	B.r = r
	bytesm[k] = append(bytesm[k], B)
	bytesl.Unlock()

	go func() {
		for b := range w {
			write(Tbytes, B.key, B.idx, b)
		}
	}()

	return w, r
}

func MakeRune(key ...string) (chan<- rune, <-chan rune) {
	k := getKey(key...)

	start.Do(run)

	w, r := make(chan rune), make(chan rune)

	runel.Lock()
	var R wrrune
	R.key = k
	R.idx = len(runem[k])
	R.r = r
	runem[k] = append(runem[k], R)
	runel.Unlock()

	go func() {
		for r := range w {
			write(Trune, R.key, R.idx, rune2bytes(r))
		}
	}()

	return w, r
}

func MakeFloat32(key ...string) (chan<- float32, <-chan float32) {
	k := getKey(key...)

	start.Do(run)

	w, r := make(chan float32), make(chan float32)

	float32l.Lock()
	var F wrfloat32
	F.key = k
	F.idx = len(float32m[k])
	F.r = r
	float32m[k] = append(float32m[k], F)
	float32l.Unlock()

	go func() {
		for f := range w {
			write(Tfloat32, F.key, F.idx, float322bytes(f))
		}
	}()

	return w, r
}

func MakeFloat64(key ...string) (chan<- float64, <-chan float64) {
	k := getKey(key...)

	start.Do(run)

	w, r := make(chan float64), make(chan float64)

	float64l.Lock()
	var F wrfloat64
	F.key = k
	F.idx = len(float64m[k])
	F.r = r
	float64m[k] = append(float64m[k], F)
	float64l.Unlock()

	go func() {
		for f := range w {
			write(Tfloat64, F.key, F.idx, float642bytes(f))
		}
	}()

	return w, r
}

func MakeError(key ...string) (chan<- error, <-chan error) {
	k := getKey(key...)

	start.Do(run)

	w, r := make(chan error), make(chan error)

	errorl.Lock()
	var E wrerror
	E.key = k
	E.idx = len(errorm[k])
	E.r = r
	errorm[k] = append(errorm[k], E)
	errorl.Unlock()

	go func() {
		for e := range w {
			write(Terror, E.key, E.idx, error2bytes(e))
		}
	}()

	return w, r
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
