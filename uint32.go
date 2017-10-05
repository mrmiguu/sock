package sock

func MakeUint32(name string, buf ...int) (chan<- uint32, <-chan uint32) {
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

	uint32Dict.Lock()
	if uint32Dict.m == nil {
		uint32Dict.m = map[string][]*tuint32{}
	}
	U := &tuint32{
		name: name,
		len:  buflen,
		idx:  len(uint32Dict.m[name]),
		selw: make(chan []byte, buflen),
		selr: make(chan []byte, buflen),
		w:    make(chan []byte, buflen),
		r:    make(chan []byte, buflen),
		cw:   make(chan uint32, buflen),
		cr:   make(chan uint32, buflen),
	}
	if !IsClient {
		U.seln = make(chan int)
		U.n = make(chan int)
	}
	uint32Dict.m[U.name] = append(uint32Dict.m[U.name], U)
	uint32Dict.Unlock()

	go wIfClient(U.selw, Tuint32, U.name, U.idx, 1)
	go rIfClient(U.selr, Tuint32, U.name, U.idx, 1)
	go wIfClient(U.w, Tuint32, U.name, U.idx, 0)
	go rIfClient(U.r, Tuint32, U.name, U.idx, 0)
	go U.selsend()
	go U.selrecv()

	return U.cw, U.cr
}

func (U *tuint32) selsend() {
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
			U.w <- uint322bytes(<-U.cw)
		}
	}
}

func (U *tuint32) selrecv() {
	for {
		<-U.selr
		U.cr <- bytes2uint32(<-U.r)
	}
}

func finduint32(name string, idx int) (*tuint32, bool) {
	uint32Dict.RLock()
	defer uint32Dict.RUnlock()

	Ui, found := uint32Dict.m[name]
	if !found || idx > len(Ui)-1 {
		return nil, false
	}
	return Ui[idx], true
}

func (U *tuint32) getuint32(sel byte, b []byte) {
	if sel == 1 {
		U.selr <- []byte{}
	} else {
		U.r <- b
	}
}

func (U *tuint32) setuint32(sel byte) []byte {
	if sel == 1 {
		U.seln <- 1
		<-U.selw
		return []byte{}
	}
	U.n <- 1
	return <-U.w
}
