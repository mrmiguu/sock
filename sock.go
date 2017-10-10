package sock

import "github.com/gopherjs/gopherjs/js"

var (
	Addr     = "localhost:80"
	IsClient = js.Global != nil
	Root     = "www"
	API      = "/abc123"
)

func MakeBool(key string) (chan<- bool, <-chan bool) {
	start.Do(run)

	w, r := make(chan bool), make(chan bool)

	booll.Lock()
	var B wrbool
	B.key = key
	B.idx = len(boolm[key])
	B.r = r
	boolm[key] = append(boolm[key], B)
	booll.Unlock()

	go func() {
		for b := range w {
			write(Tbool, B.key, B.idx, bool2bytes(b))
		}
	}()

	return w, r
}

func MakeString(key string) (chan<- string, <-chan string) {
	start.Do(run)

	w, r := make(chan string), make(chan string)

	stringl.Lock()
	var S wrstring
	S.key = key
	S.idx = len(stringm[key])
	S.r = r
	stringm[key] = append(stringm[key], S)
	stringl.Unlock()

	go func() {
		for s := range w {
			write(Tstring, S.key, S.idx, []byte(s))
		}
	}()

	return w, r
}

func MakeInt(key string) (chan<- int, <-chan int) {
	start.Do(run)

	w, r := make(chan int), make(chan int)

	intl.Lock()
	var I wrint
	I.key = key
	I.idx = len(intm[key])
	I.r = r
	intm[key] = append(intm[key], I)
	intl.Unlock()

	go func() {
		for i := range w {
			write(Tint, I.key, I.idx, int2bytes(i))
		}
	}()

	return w, r
}

func MakeInt8(key string) (chan<- int8, <-chan int8) {
	start.Do(run)

	w, r := make(chan int8), make(chan int8)

	int8l.Lock()
	var I wrint8
	I.key = key
	I.idx = len(int8m[key])
	I.r = r
	int8m[key] = append(int8m[key], I)
	int8l.Unlock()

	go func() {
		for i := range w {
			write(Tint8, I.key, I.idx, int82bytes(i))
		}
	}()

	return w, r
}

func MakeInt16(key string) (chan<- int16, <-chan int16) {
	start.Do(run)

	w, r := make(chan int16), make(chan int16)

	int16l.Lock()
	var I wrint16
	I.key = key
	I.idx = len(int16m[key])
	I.r = r
	int16m[key] = append(int16m[key], I)
	int16l.Unlock()

	go func() {
		for i := range w {
			write(Tint16, I.key, I.idx, int162bytes(i))
		}
	}()

	return w, r
}

func MakeInt32(key string) (chan<- int32, <-chan int32) {
	start.Do(run)

	w, r := make(chan int32), make(chan int32)

	int32l.Lock()
	var I wrint32
	I.key = key
	I.idx = len(int32m[key])
	I.r = r
	int32m[key] = append(int32m[key], I)
	int32l.Unlock()

	go func() {
		for i := range w {
			write(Tint32, I.key, I.idx, int322bytes(i))
		}
	}()

	return w, r
}

func MakeInt64(key string) (chan<- int64, <-chan int64) {
	start.Do(run)

	w, r := make(chan int64), make(chan int64)

	int64l.Lock()
	var I wrint64
	I.key = key
	I.idx = len(int64m[key])
	I.r = r
	int64m[key] = append(int64m[key], I)
	int64l.Unlock()

	go func() {
		for i := range w {
			write(Tint64, I.key, I.idx, int642bytes(i))
		}
	}()

	return w, r
}

func MakeUint(key string) (chan<- uint, <-chan uint) {
	start.Do(run)

	w, r := make(chan uint), make(chan uint)

	uintl.Lock()
	var U wruint
	U.key = key
	U.idx = len(uintm[key])
	U.r = r
	uintm[key] = append(uintm[key], U)
	uintl.Unlock()

	go func() {
		for u := range w {
			write(Tuint, U.key, U.idx, uint2bytes(u))
		}
	}()

	return w, r
}

func MakeUint8(key string) (chan<- uint8, <-chan uint8) {
	start.Do(run)

	w, r := make(chan uint8), make(chan uint8)

	uint8l.Lock()
	var U wruint8
	U.key = key
	U.idx = len(uint8m[key])
	U.r = r
	uint8m[key] = append(uint8m[key], U)
	uint8l.Unlock()

	go func() {
		for u := range w {
			write(Tuint8, U.key, U.idx, uint82bytes(u))
		}
	}()

	return w, r
}

func MakeUint16(key string) (chan<- uint16, <-chan uint16) {
	start.Do(run)

	w, r := make(chan uint16), make(chan uint16)

	uint16l.Lock()
	var U wruint16
	U.key = key
	U.idx = len(uint16m[key])
	U.r = r
	uint16m[key] = append(uint16m[key], U)
	uint16l.Unlock()

	go func() {
		for u := range w {
			write(Tuint16, U.key, U.idx, uint162bytes(u))
		}
	}()

	return w, r
}

func MakeUint32(key string) (chan<- uint32, <-chan uint32) {
	start.Do(run)

	w, r := make(chan uint32), make(chan uint32)

	uint32l.Lock()
	var U wruint32
	U.key = key
	U.idx = len(uint32m[key])
	U.r = r
	uint32m[key] = append(uint32m[key], U)
	uint32l.Unlock()

	go func() {
		for u := range w {
			write(Tuint32, U.key, U.idx, uint322bytes(u))
		}
	}()

	return w, r
}

func MakeUint64(key string) (chan<- uint64, <-chan uint64) {
	start.Do(run)

	w, r := make(chan uint64), make(chan uint64)

	uint64l.Lock()
	var U wruint64
	U.key = key
	U.idx = len(uint64m[key])
	U.r = r
	uint64m[key] = append(uint64m[key], U)
	uint64l.Unlock()

	go func() {
		for u := range w {
			write(Tuint64, U.key, U.idx, uint642bytes(u))
		}
	}()

	return w, r
}

func MakeByte(key string) (chan<- byte, <-chan byte) {
	start.Do(run)

	w, r := make(chan byte), make(chan byte)

	bytel.Lock()
	var B wrbyte
	B.key = key
	B.idx = len(bytem[key])
	B.r = r
	bytem[key] = append(bytem[key], B)
	bytel.Unlock()

	go func() {
		for b := range w {
			write(Tbyte, B.key, B.idx, []byte{b})
		}
	}()

	return w, r
}

func MakeBytes(key string) (chan<- []byte, <-chan []byte) {
	start.Do(run)

	w, r := make(chan []byte), make(chan []byte)

	bytesl.Lock()
	var B wrbytes
	B.key = key
	B.idx = len(bytesm[key])
	B.r = r
	bytesm[key] = append(bytesm[key], B)
	bytesl.Unlock()

	go func() {
		for b := range w {
			write(Tbytes, B.key, B.idx, b)
		}
	}()

	return w, r
}

func MakeRune(key string) (chan<- rune, <-chan rune) {
	start.Do(run)

	w, r := make(chan rune), make(chan rune)

	runel.Lock()
	var R wrrune
	R.key = key
	R.idx = len(runem[key])
	R.r = r
	runem[key] = append(runem[key], R)
	runel.Unlock()

	go func() {
		for r := range w {
			write(Trune, R.key, R.idx, rune2bytes(r))
		}
	}()

	return w, r
}

func MakeFloat32(key string) (chan<- float32, <-chan float32) {
	start.Do(run)

	w, r := make(chan float32), make(chan float32)

	float32l.Lock()
	var F wrfloat32
	F.key = key
	F.idx = len(float32m[key])
	F.r = r
	float32m[key] = append(float32m[key], F)
	float32l.Unlock()

	go func() {
		for f := range w {
			write(Tfloat32, F.key, F.idx, float322bytes(f))
		}
	}()

	return w, r
}

func MakeFloat64(key string) (chan<- float64, <-chan float64) {
	start.Do(run)

	w, r := make(chan float64), make(chan float64)

	float64l.Lock()
	var F wrfloat64
	F.key = key
	F.idx = len(float64m[key])
	F.r = r
	float64m[key] = append(float64m[key], F)
	float64l.Unlock()

	go func() {
		for f := range w {
			write(Tfloat64, F.key, F.idx, float642bytes(f))
		}
	}()

	return w, r
}

func MakeError(key string) (chan<- error, <-chan error) {
	start.Do(run)

	w, r := make(chan error), make(chan error)

	errorl.Lock()
	var E wrerror
	E.key = key
	E.idx = len(errorm[key])
	E.r = r
	errorm[key] = append(errorm[key], E)
	errorl.Unlock()

	go func() {
		for e := range w {
			write(Terror, E.key, E.idx, error2bytes(e))
		}
	}()

	return w, r
}
