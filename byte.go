package sock

func MakeByte(name string, buf ...int) (chan<- byte, <-chan byte) {
	if len(buf) > 1 {
		panic("too many arguments")
	}
	buflen := 0
	if len(buf) > 0 {
		if buf[0] < 0 {
			panic("negative buffer argument")
		}
		buflen = buf[0]
	}

	go started.Do(wAndOrRIfServer)

	byteDict.Lock()
	if byteDict.m == nil {
		byteDict.m = map[string][]*tbyte{}
	}
	B := &tbyte{
		name: name,
		len:  buflen,
		idx:  len(byteDict.m[name]),
		w:    make(chan []byte, buflen),
		r:    make(chan []byte, buflen),
		cw:   make(chan byte, buflen),
		cr:   make(chan byte, buflen),
	}
	if !IsClient {
		B.n = make(chan int)
	}
	byteDict.m[B.name] = append(byteDict.m[B.name], B)
	byteDict.Unlock()

	go wIfClient(B.w, Tbyte, B.name, B.idx)
	go rIfClient(B.r, Tbyte, B.name, B.idx)
	go B.selsend()
	go B.selrecv()

	return B.cw, B.cr
}

func (B *tbyte) selsend() {
		b := []byte{<-B.cw}
		for ok := true; ok; ok = (len(B.n) > 0) {
			if !IsClient {
				<-B.n
			}
			B.w <- b
		}
	}
}

func (B *tbyte) selrecv() {
	for {
		B.cr <- (<-B.r)[0]
	}
}

func findbyte(name string, idx int) (*tbyte, bool) {
	byteDict.RLock()
	defer byteDict.RUnlock()

	Bi, found := byteDict.m[name]
	if !found || idx > len(Bi)-1 {
		return nil, false
	}
	return Bi[idx], true
}

func (B *tbyte) getbyte(b []byte) {
		B.r <- b
}

func (B *tbyte) setbyte() []byte {
	B.n <- 1
	return <-B.w
}
