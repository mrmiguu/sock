package sock

func MakeBool(name string, buf ...int) (chan<- bool, <-chan bool) {
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

	boolDict.Lock()
	if boolDict.m == nil {
		boolDict.m = map[string][]*tbool{}
	}
	B := &tbool{
		name: name,
		len:  buflen,
		idx:  len(boolDict.m[name]),
		selw: make(chan []byte, buflen),
		selr: make(chan []byte, buflen),
		w:    make(chan []byte, buflen),
		r:    make(chan []byte, buflen),
		cw:   make(chan bool, buflen),
		cr:   make(chan bool, buflen),
	}
	if !IsClient {
		B.seln = make(chan int)
		B.n = make(chan int)
	}
	boolDict.m[B.name] = append(boolDict.m[B.name], B)
	boolDict.Unlock()

	go wIfClient(B.selw, Tbool, B.name, B.idx, 1)
	go rIfClient(B.selr, Tbool, B.name, B.idx, 1)
	go wIfClient(B.w, Tbool, B.name, B.idx, 0)
	go rIfClient(B.r, Tbool, B.name, B.idx, 0)
	go B.selsend()
	go B.selrecv()

	return B.cw, B.cr
}

func (B *tbool) selsend() {
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
			B.w <- bool2bytes(<-B.cw)
		}
	}
}

func (B *tbool) selrecv() {
	for {
		<-B.selr
		B.cr <- bytes2bool(<-B.r)
	}
}

func findbool(name string, idx int) (*tbool, bool) {
	boolDict.RLock()
	defer boolDict.RUnlock()

	Bi, found := boolDict.m[name]
	if !found || idx > len(Bi)-1 {
		return nil, false
	}
	return Bi[idx], true
}

func (B *tbool) getbool(sel byte, b []byte) {
	if sel == 1 {
		B.selr <- nil
	} else {
		B.r <- b
	}
}

func (B *tbool) setbool(sel byte) []byte {
	if sel == 1 {
		B.seln <- 1
		return <-B.selw
	}
	B.n <- 1
	return <-B.w
}
