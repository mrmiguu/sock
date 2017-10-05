package sock

func MakeUint64(name string, buf ...int) (chan<- uint64, <-chan uint64) {
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

	uint64Dict.Lock()
	if uint64Dict.m == nil {
		uint64Dict.m = map[string][]*tuint64{}
	}
	U := &tuint64{
		name: name,
		len:  buflen,
		idx:  len(uint64Dict.m[name]),
		selw: make(chan []byte, buflen),
		selr: make(chan []byte, buflen),
		w:    make(chan []byte, buflen),
		r:    make(chan []byte, buflen),
		cw:   make(chan uint64, buflen),
		cr:   make(chan uint64, buflen),
	}
	if !IsClient {
		U.seln = make(chan int)
		U.n = make(chan int)
	}
	uint64Dict.m[U.name] = append(uint64Dict.m[U.name], U)
	uint64Dict.Unlock()

	go wIfClient(U.selw, Tuint64, U.name, U.idx, 1)
	go rIfClient(U.selr, Tuint64, U.name, U.idx, 1)
	go wIfClient(U.w, Tuint64, U.name, U.idx, 0)
	go rIfClient(U.r, Tuint64, U.name, U.idx, 0)
	go U.selsend()
	go U.selrecv()

	return U.cw, U.cr
}

func (U *tuint64) selsend() {
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
			U.w <- uint642bytes(<-U.cw)
		}
	}
}

func (U *tuint64) selrecv() {
	for {
		<-U.selr
		U.cr <- bytes2uint64(<-U.r)
	}
}

func finduint64(name string, idx int) (*tuint64, bool) {
	uint64Dict.RLock()
	defer uint64Dict.RUnlock()

	Ui, found := uint64Dict.m[name]
	if !found || idx > len(Ui)-1 {
		return nil, false
	}
	return Ui[idx], true
}

func (U *tuint64) getuint64(sel byte, b []byte) {
	if sel == 1 {
		U.selr <- []byte{}
	} else {
		U.r <- b
	}
}

func (U *tuint64) setuint64(sel byte) []byte {
	if sel == 1 {
		U.seln <- 1
		<-U.selw
		return []byte{}
	}
	U.n <- 1
	return <-U.w
}
