package sock

func MakeComplex64(name string, buf ...int) (chan<- complex64, <-chan complex64) {
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

	complex64Dict.Lock()
	if complex64Dict.m == nil {
		complex64Dict.m = map[string][]*tcomplex64{}
	}
	C := &tcomplex64{
		name: name,
		len:  buflen,
		idx:  len(complex64Dict.m[name]),
		selw: make(chan []byte, buflen),
		selr: make(chan []byte, buflen),
		w:    make(chan []byte, buflen),
		r:    make(chan []byte, buflen),
		cw:   make(chan complex64, buflen),
		cr:   make(chan complex64, buflen),
	}
	if !IsClient {
		C.seln = make(chan int)
		C.n = make(chan int)
	}
	complex64Dict.m[C.name] = append(complex64Dict.m[C.name], C)
	complex64Dict.Unlock()

	go wIfClient(C.selw, Tcomplex64, C.name, C.idx, 1)
	go rIfClient(C.selr, Tcomplex64, C.name, C.idx, 1)
	go wIfClient(C.w, Tcomplex64, C.name, C.idx, 0)
	go rIfClient(C.r, Tcomplex64, C.name, C.idx, 0)
	go C.selsend()
	go C.selrecv()

	return C.cw, C.cr
}

func (C *tcomplex64) selsend() {
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
			C.w <- complex642bytes(<-C.cw)
		}
	}
}

func (C *tcomplex64) selrecv() {
	for {
		<-C.selr
		C.cr <- bytes2complex64(<-C.r)
	}
}

func findcomplex64(name string, idx int) (*tcomplex64, bool) {
	complex64Dict.RLock()
	defer complex64Dict.RUnlock()

	Ci, found := complex64Dict.m[name]
	if !found || idx > len(Ci)-1 {
		return nil, false
	}
	return Ci[idx], true
}

func (C *tcomplex64) getcomplex64(sel byte, b []byte) {
	if sel == 1 {
		C.selr <- []byte{}
	} else {
		C.r <- b
	}
}

func (C *tcomplex64) setcomplex64(sel byte) []byte {
	if sel == 1 {
		C.seln <- 1
		<-C.selw
		return []byte{}
	}
	C.n <- 1
	return <-C.w
}
