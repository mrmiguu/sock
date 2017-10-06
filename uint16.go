package sock

func MakeUint16(name string, buf ...int) (chan<- uint16, <-chan uint16) {
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

	uint16Dict.Lock()
	if uint16Dict.m == nil {
		uint16Dict.m = map[string][]*tuint16{}
	}
	U := &tuint16{
		name: name,
		len:  buflen,
		idx:  len(uint16Dict.m[name]),
		selw: make(chan []byte, buflen),
		selr: make(chan []byte, buflen),
		w:    make(chan []byte, buflen),
		r:    make(chan []byte, buflen),
		cw:   make(chan uint16, buflen),
		cr:   make(chan uint16, buflen),
	}
	if !IsClient {
		U.seln = make(chan int)
		U.n = make(chan int)
	}
	uint16Dict.m[U.name] = append(uint16Dict.m[U.name], U)
	uint16Dict.Unlock()

	go wIfClient(U.selw, U.w, Tuint16, U.name, U.idx)
	go rIfClient(U.selr, U.r, Tuint16, U.name, U.idx)
	go U.selsend()
	go U.selrecv()

	return U.cw, U.cr
}

func (U *tuint16) selsend() {
	for {
		for ok := true; ok; ok = (len(U.seln) > 0) {
			if !IsClient {
				<-U.seln
			}
			U.selw <- nil
		}

		b := uint162bytes(<-U.cw)
		for ok := true; ok; ok = (len(U.n) > 0) {
			if !IsClient {
				<-U.n
			}
			U.w <- b
		}
	}
}

func (U *tuint16) selrecv() {
	for {
		<-U.selr
		U.cr <- bytes2uint16(<-U.r)
	}
}

func finduint16(name string, idx int) (*tuint16, bool) {
	uint16Dict.RLock()
	defer uint16Dict.RUnlock()

	Ui, found := uint16Dict.m[name]
	if !found || idx > len(Ui)-1 {
		return nil, false
	}
	return Ui[idx], true
}

func (U *tuint16) getuint16(sel byte, b []byte) {
	if sel == 1 {
		U.selr <- nil
	} else {
		U.r <- b
	}
}

func (U *tuint16) setuint16(sel byte) []byte {
	if sel == 1 {
		U.seln <- 1
		return <-U.selw
	}
	U.n <- 1
	return <-U.w
}
