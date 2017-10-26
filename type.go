package sock

type keyAndIdx struct {
	key string
	idx int
}

type wboolx struct {
	i  int
	sl []wbool
}
type rboolx struct {
	i  int
	sl []rbool
}
type wstringx struct {
	i  int
	sl []wstring
}
type rstringx struct {
	i  int
	sl []rstring
}
type wintx struct {
	i  int
	sl []wint
}
type rintx struct {
	i  int
	sl []rint
}
type wint8x struct {
	i  int
	sl []wint8
}
type rint8x struct {
	i  int
	sl []rint8
}
type wint16x struct {
	i  int
	sl []wint16
}
type rint16x struct {
	i  int
	sl []rint16
}
type wint32x struct {
	i  int
	sl []wint32
}
type rint32x struct {
	i  int
	sl []rint32
}
type wint64x struct {
	i  int
	sl []wint64
}
type rint64x struct {
	i  int
	sl []rint64
}
type wuintx struct {
	i  int
	sl []wuint
}
type ruintx struct {
	i  int
	sl []ruint
}
type wuint8x struct {
	i  int
	sl []wuint8
}
type ruint8x struct {
	i  int
	sl []ruint8
}
type wuint16x struct {
	i  int
	sl []wuint16
}
type ruint16x struct {
	i  int
	sl []ruint16
}
type wuint32x struct {
	i  int
	sl []wuint32
}
type ruint32x struct {
	i  int
	sl []ruint32
}
type wuint64x struct {
	i  int
	sl []wuint64
}
type ruint64x struct {
	i  int
	sl []ruint64
}
type wbytex struct {
	i  int
	sl []wbyte
}
type rbytex struct {
	i  int
	sl []rbyte
}
type wbytesx struct {
	i  int
	sl []wbytes
}
type rbytesx struct {
	i  int
	sl []rbytes
}
type wrunex struct {
	i  int
	sl []wrune
}
type rrunex struct {
	i  int
	sl []rrune
}
type wfloat32x struct {
	i  int
	sl []wfloat32
}
type rfloat32x struct {
	i  int
	sl []rfloat32
}
type wfloat64x struct {
	i  int
	sl []wfloat64
}
type rfloat64x struct {
	i  int
	sl []rfloat64
}
type werrorx struct {
	i  int
	sl []werror
}
type rerrorx struct {
	i  int
	sl []rerror
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
