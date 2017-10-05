package sock

func MakeUint(name string, buf ...int) (chan<- uint, <-chan uint) {
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

	uintDict.Lock()
	if uintDict.m == nil {
		uintDict.m = map[string][]*tuint{}
	}
	U := &tuint{
		name: name,
		len:  buflen,
		idx:  len(uintDict.m[name]),
		selw: make(chan []byte, buflen),
		selr: make(chan []byte, buflen),
		w:    make(chan []byte, buflen),
		r:    make(chan []byte, buflen),
		cw:   make(chan uint, buflen),
		cr:   make(chan uint, buflen),
	}
	if !IsClient {
		U.seln = make(chan int)
		U.n = make(chan int)
	}
	uintDict.m[U.name] = append(uintDict.m[U.name], U)
	uintDict.Unlock()

	go wIfClient(U.selw, Tuint, U.name, U.idx, 1)
	go rIfClient(U.selr, Tuint, U.name, U.idx, 1)
	go wIfClient(U.w, Tuint, U.name, U.idx, 0)
	go rIfClient(U.r, Tuint, U.name, U.idx, 0)
	go U.selsend()
	go U.selrecv()

	return U.cw, U.cr
}

func (U *tuint) selsend() {
	for {
		for ok := true; ok; ok = (len(U.seln) > 0) {
			if !IsClient {
				<-U.seln
			}
			U.selw <- nil
		}

		for ok := true; ok; ok = (len(U.n) > 0) {
			if !IsClient {
				<-U.n
			}
			U.w <- uint2bytes(<-U.cw)
		}
	}
}

func (U *tuint) selrecv() {
	for {
		<-U.selr
		U.cr <- bytes2uint(<-U.r)
	}
}

func finduint(name string, idx int) (*tuint, bool) {
	uintDict.RLock()
	defer uintDict.RUnlock()

	Ui, found := uintDict.m[name]
	if !found || idx > len(Ui)-1 {
		return nil, false
	}
	return Ui[idx], true
}

func (U *tuint) getuint(sel byte, b []byte) {
	if sel == 1 {
		U.selr <- []byte{}
	} else {
		U.r <- b
	}
}

func (U *tuint) setuint(sel byte) []byte {
	if sel == 1 {
		U.seln <- 1
		<-U.selw
		return []byte{}
	}
	U.n <- 1
	return <-U.w
}