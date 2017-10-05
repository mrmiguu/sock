package sock

// type Error chan error

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

// type Bool struct {
// 	Name string
// 	Len  uint

// 	sel, reg private
// }

// type String struct {
// 	Name string
// 	Len  uint

// 	sel, reg private
// }

// type Int struct {
// 	Name string
// 	Len  uint

// 	sel, reg private
// }

// type Int8 struct {
// 	Name string
// 	Len  uint

// 	sel, reg private
// }

// type Int16 struct {
// 	Name string
// 	Len  uint

// 	sel, reg private
// }

// type Int32 struct {
// 	Name string
// 	Len  uint

// 	sel, reg private
// }

// type Int64 struct {
// 	Name string
// 	Len  uint

// 	sel, reg private
// }

// type Uint struct {
// 	Name string
// 	Len  uint

// 	sel, reg private
// }

// type Uint8 struct {
// 	Name string
// 	Len  uint

// 	sel, reg private
// }

// type Uint16 struct {
// 	Name string
// 	Len  uint

// 	sel, reg private
// }

// type Uint32 struct {
// 	Name string
// 	Len  uint

// 	sel, reg private
// }

// type Uint64 struct {
// 	Name string
// 	Len  uint

// 	sel, reg private
// }

// type Uintptr struct {
// 	Name string
// 	Len  uint

// 	sel, reg private
// }

// type Byte struct {
// 	Name string
// 	Len  uint

// 	sel, reg private
// }

// type Bytes struct {
// 	Name string
// 	Len  uint

// 	sel, reg private
// }

// type Rune struct {
// 	Name string
// 	Len  uint

// 	sel, reg private
// }

// type Float32 struct {
// 	Name string
// 	Len  uint

// 	sel, reg private
// }

// type Float64 struct {
// 	Name string
// 	Len  uint

// 	sel, reg private
// }

// type Complex64 struct {
// 	Name string
// 	Len  uint

// 	sel, reg private
// }

// type Complex128 struct {
// 	Name string
// 	Len  uint

// 	sel, reg private
// }
