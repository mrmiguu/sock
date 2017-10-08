package sock

func MakeUint16(name string, buf ...int) (chan<- uint16, <-chan uint16) {
	if len(buf) > 1 {
		panic("too many arguments")
	}
	buflen := 1
	if len(buf) > 0 {
		if buf[0] < 1 {
			panic("buffer argument less than one")
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
		w:    make(chan []byte, buflen),
		r:    make(chan []byte, buflen),
		cw:   make(chan uint16, buflen),
		cr:   make(chan uint16, buflen),
	}
	if !IsClient {
		U.n = make(chan int)
	}
	uint16Dict.m[U.name] = append(uint16Dict.m[U.name], U)
	uint16Dict.Unlock()

	go wIfClient(U.w, Tuint16, U.name, U.idx)
	go rIfClient(U.r, Tuint16, U.name, U.idx)
	go U.selsend()
	go U.selrecv()

	return U.cw, U.cr
}

func (U *tuint16) selsend() {
	for {
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
