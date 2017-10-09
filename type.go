package sock

type tbyte struct {
	name   string
	len    int
	idx    int
	w, r   chan []byte
	cw, cr chan byte
	n      chan int
}
