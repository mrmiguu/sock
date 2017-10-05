package sock

type terror struct {
	name             string
	len              int
	idx              int
	selw, w, selr, r chan []byte
	cw, cr           chan error
	seln, n          chan int
}

type tbytes struct {
	name             string
	len              int
	idx              int
	selw, w, selr, r chan []byte
	cw, cr           chan []byte
	seln, n          chan int
}

type tstring struct {
	name             string
	len              int
	idx              int
	selw, w, selr, r chan []byte
	cw, cr           chan string
	seln, n          chan int
}

type tint struct {
	name             string
	len              int
	idx              int
	selw, w, selr, r chan []byte
	cw, cr           chan int
	seln, n          chan int
}

type trune struct {
	name             string
	len              int
	idx              int
	selw, w, selr, r chan []byte
	cw, cr           chan rune
	seln, n          chan int
}

type tbool struct {
	name             string
	len              int
	idx              int
	selw, w, selr, r chan []byte
	cw, cr           chan bool
	seln, n          chan int
}

type tfloat64 struct {
	name             string
	len              int
	idx              int
	selw, w, selr, r chan []byte
	cw, cr           chan float64
	seln, n          chan int
}

type tint8 struct {
	name             string
	len              int
	idx              int
	selw, w, selr, r chan []byte
	cw, cr           chan int8
	seln, n          chan int
}

type tint16 struct {
	name             string
	len              int
	idx              int
	selw, w, selr, r chan []byte
	cw, cr           chan int16
	seln, n          chan int
}

type tint32 struct {
	name             string
	len              int
	idx              int
	selw, w, selr, r chan []byte
	cw, cr           chan int32
	seln, n          chan int
}

type tint64 struct {
	name             string
	len              int
	idx              int
	selw, w, selr, r chan []byte
	cw, cr           chan int64
	seln, n          chan int
}

type tuint struct {
	name             string
	len              int
	idx              int
	selw, w, selr, r chan []byte
	cw, cr           chan uint
	seln, n          chan int
}

type tuint8 struct {
	name             string
	len              int
	idx              int
	selw, w, selr, r chan []byte
	cw, cr           chan uint8
	seln, n          chan int
}

type tuint16 struct {
	name             string
	len              int
	idx              int
	selw, w, selr, r chan []byte
	cw, cr           chan uint16
	seln, n          chan int
}

type tuint32 struct {
	name             string
	len              int
	idx              int
	selw, w, selr, r chan []byte
	cw, cr           chan uint32
	seln, n          chan int
}

type tuint64 struct {
	name             string
	len              int
	idx              int
	selw, w, selr, r chan []byte
	cw, cr           chan uint64
	seln, n          chan int
}

type tuintptr struct {
	name             string
	len              int
	idx              int
	selw, w, selr, r chan []byte
	cw, cr           chan uintptr
	seln, n          chan int
}

type tbyte struct {
	name             string
	len              int
	idx              int
	selw, w, selr, r chan []byte
	cw, cr           chan byte
	seln, n          chan int
}

type tfloat32 struct {
	name             string
	len              int
	idx              int
	selw, w, selr, r chan []byte
	cw, cr           chan float32
	seln, n          chan int
}

type tcomplex64 struct {
	name             string
	len              int
	idx              int
	selw, w, selr, r chan []byte
	cw, cr           chan complex64
	seln, n          chan int
}

type tcomplex128 struct {
	name             string
	len              int
	idx              int
	selw, w, selr, r chan []byte
	cw, cr           chan complex128
	seln, n          chan int
}
