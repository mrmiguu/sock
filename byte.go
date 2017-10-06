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
		selw: make(chan []byte, buflen),
		selr: make(chan []byte, buflen),
		w:    make(chan []byte, buflen),
		r:    make(chan []byte, buflen),
		cw:   make(chan byte, buflen),
		cr:   make(chan byte, buflen),
	}
	if !IsClient {
		B.seln = make(chan int)
		B.n = make(chan int)
	}
	byteDict.m[B.name] = append(byteDict.m[B.name], B)
	byteDict.Unlock()

	go wIfClient(B.selw, Tbyte, B.name, B.idx, 1)
	go rIfClient(B.selr, Tbyte, B.name, B.idx, 1)
	go wIfClient(B.w, Tbyte, B.name, B.idx, 0)
	go rIfClient(B.r, Tbyte, B.name, B.idx, 0)
	go B.selsend()
	go B.selrecv()

	return B.cw, B.cr
}

func (B *tbyte) selsend() {
	for {
		for ok := true; ok; ok = (len(B.seln) > 0) {
			if !IsClient {
				<-B.seln
			}
			B.selw <- nil
		}

		for ok := true; ok; ok = (len(B.n) > 0) {
			if !IsClient {
				<-B.n
			}
			B.w <- []byte{<-B.cw}
		}
	}
}

func (B *tbyte) selrecv() {
	for {
		<-B.selr
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

func (B *tbyte) getbyte(sel byte, b []byte) {
	if sel == 1 {
		B.selr <- nil
	} else {
		B.r <- b
	}
}

func (B *tbyte) setbyte(sel byte) []byte {
	if sel == 1 {
		B.seln <- 1
		return <-B.selw
	}
	B.n <- 1
	return <-B.w
}
