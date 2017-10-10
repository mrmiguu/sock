package sock

type keyAndIdx struct {
	key string
	idx int
}

type wrbool struct {
	keyAndIdx
	r chan bool
}

type wrstring struct {
	keyAndIdx
	r chan string
}

type wrint struct {
	keyAndIdx
	r chan int
}

type wrint8 struct {
	keyAndIdx
	r chan int8
}

type wrint16 struct {
	keyAndIdx
	r chan int16
}

type wrint32 struct {
	keyAndIdx
	r chan int32
}

type wrint64 struct {
	keyAndIdx
	r chan int64
}

type wruint struct {
	keyAndIdx
	r chan uint
}

type wruint8 struct {
	keyAndIdx
	r chan uint8
}

type wruint16 struct {
	keyAndIdx
	r chan uint16
}

type wruint32 struct {
	keyAndIdx
	r chan uint32
}

type wruint64 struct {
	keyAndIdx
	r chan uint64
}

type wrbyte struct {
	keyAndIdx
	r chan byte
}

type wrbytes struct {
	keyAndIdx
	r chan []byte
}

type wrrune struct {
	keyAndIdx
	r chan rune
}

type wrfloat32 struct {
	keyAndIdx
	r chan float32
}

type wrfloat64 struct {
	keyAndIdx
	r chan float64
}

type wrerror struct {
	keyAndIdx
	r chan error
}
