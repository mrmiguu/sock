package sock

type Error chan error

type sel struct {
	w, r chan interface{}
	n    chan int
}

type terror struct {
	name string
	len  int
	idx  int
	sel  sel
	c    Error
	n    chan int
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
