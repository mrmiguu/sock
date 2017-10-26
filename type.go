package sock

type keyAndIdx struct {
	key string
	idx int
}

type wbool struct {
	keyAndIdx
	w chan bool
}
type rbool struct {
	keyAndIdx
	r chan bool
}

type wstring struct {
	keyAndIdx
	w chan string
}
type rstring struct {
	keyAndIdx
	r chan string
}

type wint struct {
	keyAndIdx
	w chan int
}
type rint struct {
	keyAndIdx
	r chan int
}

type wint8 struct {
	keyAndIdx
	w chan int8
}
type rint8 struct {
	keyAndIdx
	r chan int8
}

type wint16 struct {
	keyAndIdx
	w chan int16
}
type rint16 struct {
	keyAndIdx
	r chan int16
}

type wint32 struct {
	keyAndIdx
	w chan int32
}
type rint32 struct {
	keyAndIdx
	r chan int32
}

type wint64 struct {
	keyAndIdx
	w chan int64
}
type rint64 struct {
	keyAndIdx
	r chan int64
}

type wuint struct {
	keyAndIdx
	w chan uint
}
type ruint struct {
	keyAndIdx
	r chan uint
}

type wuint8 struct {
	keyAndIdx
	w chan uint8
}
type ruint8 struct {
	keyAndIdx
	r chan uint8
}

type wuint16 struct {
	keyAndIdx
	w chan uint16
}
type ruint16 struct {
	keyAndIdx
	r chan uint16
}

type wuint32 struct {
	keyAndIdx
	w chan uint32
}
type ruint32 struct {
	keyAndIdx
	r chan uint32
}

type wuint64 struct {
	keyAndIdx
	w chan uint64
}
type ruint64 struct {
	keyAndIdx
	r chan uint64
}

type wbyte struct {
	keyAndIdx
	w chan byte
}
type rbyte struct {
	keyAndIdx
	r chan byte
}

type wbytes struct {
	keyAndIdx
	w chan []byte
}
type rbytes struct {
	keyAndIdx
	r chan []byte
}

type wrune struct {
	keyAndIdx
	w chan rune
}
type rrune struct {
	keyAndIdx
	r chan rune
}

type wfloat32 struct {
	keyAndIdx
	w chan float32
}
type rfloat32 struct {
	keyAndIdx
	r chan float32
}

type wfloat64 struct {
	keyAndIdx
	w chan float64
}
type rfloat64 struct {
	keyAndIdx
	r chan float64
}

type werror struct {
	keyAndIdx
	w chan error
}
type rerror struct {
	keyAndIdx
	r chan error
}
