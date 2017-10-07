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

	go started.Do(wAndOrRIfServer)

	bytesDict.Lock()
	if bytesDict.m == nil {
		bytesDict.m = map[string][]*tbytes{}
	}
	B := &tbytes{
		name: name,
		len:  buflen,
		idx:  len(bytesDict.m[name]),
		w:    make(chan []byte, buflen),
		r:    make(chan []byte, buflen),
		cw:   make(chan []byte, buflen),
		cr:   make(chan []byte, buflen),
	}
	if !IsClient {
		B.n = make(chan int)
	}
	bytesDict.m[B.name] = append(bytesDict.m[B.name], B)
	bytesDict.Unlock()

	go wIfClient(B.w, Tbytes, B.name, B.idx)
	go rIfClient(B.r, Tbytes, B.name, B.idx)
	go B.selsend()
	go B.selrecv()

	return B.cw, B.cr
}

func (B *tbytes) selsend() {
		b := <-B.cw
		for ok := true; ok; ok = (len(B.n) > 0) {
			if !IsClient {
				<-B.n
			}
			B.w <- b
		}
	}
}

func (B *tbytes) selrecv() {
	for {
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

func (B *tbytes) getbytes(b []byte) {
		B.r <- b
}

func (B *tbytes) setbytes() []byte {
	B.n <- 1
	return <-B.w
}
