package sock

func MakeBytes(name string, buf ...int) (chan<- []byte, <-chan []byte) {
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

	go wAndOrRIfServer()

	bytesDict.Lock()
	if bytesDict.m == nil {
		bytesDict.m = map[string][]*tbytes{}
	}
	B := &tbytes{
		name: name,
		len:  buflen,
		idx:  len(bytesDict.m[name]),
		selw: make(chan []byte, buflen),
		selr: make(chan []byte, buflen),
		w:    make(chan []byte, buflen),
		r:    make(chan []byte, buflen),
		cw:   make(chan []byte, buflen),
		cr:   make(chan []byte, buflen),
	}
	if !IsClient {
		B.seln = make(chan int)
		B.n = make(chan int)
	}
	bytesDict.m[B.name] = append(bytesDict.m[B.name], B)
	bytesDict.Unlock()

	go wIfClient(B.selw, Tbytes, B.name, B.idx, 1)
	go rIfClient(B.selr, Tbytes, B.name, B.idx, 1)
	go wIfClient(B.w, Tbytes, B.name, B.idx, 0)
	go rIfClient(B.r, Tbytes, B.name, B.idx, 0)
	go B.selsend()
	go B.selrecv()

	return B.cw, B.cr
}

func (B *tbytes) selsend() {
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
			B.w <- <-B.cw
		}
	}
}

func (B *tbytes) selrecv() {
	for {
		<-B.selr
		B.cr <- <-B.r
	}
}

func findbytes(name string, idx int) (*tbytes, bool) {
	bytesDict.RLock()
	defer bytesDict.RUnlock()

	Bi, found := bytesDict.m[name]
	if !found || idx > len(Bi)-1 {
		return nil, false
	}
	return Bi[idx], true
}

func (B *tbytes) getbytes(sel byte, b []byte) {
	if sel == 1 {
		B.selr <- []byte{}
	} else {
		B.r <- b
	}
}

func (B *tbytes) setbytes(sel byte) []byte {
	if sel == 1 {
		B.seln <- 1
		<-B.selw
		return []byte{}
	}
	B.n <- 1
	return <-B.w
}
