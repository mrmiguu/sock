package sock

type terror struct {
	name   string
	len    int
	idx    int
	w, r   chan []byte
	cw, cr chan error
	n      chan int
}

type tbytes struct {
	name   string
	len    int
	idx    int
	w, r   chan []byte
	cw, cr chan []byte
	n      chan int
}

type tstring struct {
	name   string
	len    int
	idx    int
	w, r   chan []byte
	cw, cr chan string
	n      chan int
}

type tint struct {
	name   string
	len    int
	idx    int
	w, r   chan []byte
	cw, cr chan int
	n      chan int
}

type trune struct {
	name   string
	len    int
	idx    int
	w, r   chan []byte
	cw, cr chan rune
	n      chan int
}

type tbool struct {
	name   string
	len    int
	idx    int
	w, r   chan []byte
	cw, cr chan bool
	n      chan int
}

type tfloat64 struct {
	name   string
	len    int
	idx    int
	w, r   chan []byte
	cw, cr chan float64
	n      chan int
}

type tint8 struct {
	name   string
	len    int
	idx    int
	w, r   chan []byte
	cw, cr chan int8
	n      chan int
}

type tint16 struct {
	name   string
	len    int
	idx    int
	w, r   chan []byte
	cw, cr chan int16
	n      chan int
}

type tint32 struct {
	name   string
	len    int
	idx    int
	w, r   chan []byte
	cw, cr chan int32
	n      chan int
}

type tint64 struct {
	name   string
	len    int
	idx    int
	w, r   chan []byte
	cw, cr chan int64
	n      chan int
}

type tuint struct {
	name   string
	len    int
	idx    int
	w, r   chan []byte
	cw, cr chan uint
	n      chan int
}

type tuint8 struct {
	name   string
	len    int
	idx    int
	w, r   chan []byte
	cw, cr chan uint8
	n      chan int
}

type tuint16 struct {
	name   string
	len    int
	idx    int
	w, r   chan []byte
	cw, cr chan uint16
	n      chan int
}

type tuint32 struct {
	name   string
	len    int
	idx    int
	w, r   chan []byte
	cw, cr chan uint32
	n      chan int
}

type tuint64 struct {
	name   string
	len    int
	idx    int
	w, r   chan []byte
	cw, cr chan uint64
	n      chan int
}

type tuintptr struct {
	name   string
	len    int
	idx    int
	w, r   chan []byte
	cw, cr chan uintptr
	n      chan int
}

type tbyte struct {
	name   string
	len    int
	idx    int
	w, r   chan []byte
	cw, cr chan byte
	n      chan int
}

type tfloat32 struct {
	name   string
	len    int
	idx    int
	w, r   chan []byte
	cw, cr chan float32
	n      chan int
}

type tcomplex64 struct {
	name   string
	len    int
	idx    int
	w, r   chan []byte
	cw, cr chan complex64
	n      chan int
}

type tcomplex128 struct {
	name   string
	len    int
	idx    int
	w, r   chan []byte
	cw, cr chan complex128
	n      chan int
}
