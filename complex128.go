package sock

func MakeComplex128(name string, buf ...int) (chan<- complex128, <-chan complex128) {
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

	complex128Dict.Lock()
	if complex128Dict.m == nil {
		complex128Dict.m = map[string][]*tcomplex128{}
	}
	C := &tcomplex128{
		name: name,
		len:  buflen,
		idx:  len(complex128Dict.m[name]),
		selw: make(chan []byte, buflen),
		selr: make(chan []byte, buflen),
		w:    make(chan []byte, buflen),
		r:    make(chan []byte, buflen),
		cw:   make(chan complex128, buflen),
		cr:   make(chan complex128, buflen),
	}
	if !IsClient {
		C.seln = make(chan int)
		C.n = make(chan int)
	}
	complex128Dict.m[C.name] = append(complex128Dict.m[C.name], C)
	complex128Dict.Unlock()

	go wIfClient(C.selw, Tcomplex128, C.name, C.idx, 1)
	go rIfClient(C.selr, Tcomplex128, C.name, C.idx, 1)
	go wIfClient(C.w, Tcomplex128, C.name, C.idx, 0)
	go rIfClient(C.r, Tcomplex128, C.name, C.idx, 0)
	go C.selsend()
	go C.selrecv()

	return C.cw, C.cr
}

func (C *tcomplex128) selsend() {
	for {
		for ok := true; ok; ok = (len(C.seln) > 0) {
			if !IsClient {
				<-C.seln
			}
			C.selw <- nil
		}

		for ok := true; ok; ok = (len(C.n) > 0) {
			if !IsClient {
				<-C.n
			}
			C.w <- complex1282bytes(<-C.cw)
		}
	}
}

func (C *tcomplex128) selrecv() {
	for {
		<-C.selr
		C.cr <- bytes2complex128(<-C.r)
	}
}

func findcomplex128(name string, idx int) (*tcomplex128, bool) {
	complex128Dict.RLock()
	defer complex128Dict.RUnlock()

	Ci, found := complex128Dict.m[name]
	if !found || idx > len(Ci)-1 {
		return nil, false
	}
	return Ci[idx], true
}

func (C *tcomplex128) getcomplex128(sel byte, b []byte) {
	if sel == 1 {
		C.selr <- []byte{}
	} else {
		C.r <- b
	}
}

func (C *tcomplex128) setcomplex128(sel byte) []byte {
	if sel == 1 {
		C.seln <- 1
		<-C.selw
		return []byte{}
	}
	C.n <- 1
	return <-C.w
}
