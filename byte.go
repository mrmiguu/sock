package sock

func MakeByte(name string, buf ...int) (chan<- byte, <-chan byte) {
	if len(buf) > 1 {
		panic("too many arguments")
	}
	buflen := 1
	if len(buf) > 0 {
		if buf[0] < 1 {
			panic("buffer argument less than one")
		}
		buflen = buf[0]
	}

	go started.Do(wAndOrR)

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

	// go wIfClient(B.w, Tbyte, B.name, B.idx)
	// go rIfClient(B.r, Tbyte, B.name, B.idx)
	go B.selsend()
	go B.selrecv()

	return B.cw, B.cr
}

func (B *tbyte) selsend() {
	for {
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
