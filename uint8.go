package sock

func MakeUint8(name string, buf ...int) (chan<- uint8, <-chan uint8) {
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

	uint8Dict.Lock()
	if uint8Dict.m == nil {
		uint8Dict.m = map[string][]*tuint8{}
	}
	U := &tuint8{
		name: name,
		len:  buflen,
		idx:  len(uint8Dict.m[name]),
		w:    make(chan []byte, buflen),
		r:    make(chan []byte, buflen),
		cw:   make(chan uint8, buflen),
		cr:   make(chan uint8, buflen),
	}
	if !IsClient {
		U.n = make(chan int)
	}
	uint8Dict.m[U.name] = append(uint8Dict.m[U.name], U)
	uint8Dict.Unlock()

	go wIfClient(U.w, Tuint8, U.name, U.idx)
	go rIfClient(U.r, Tuint8, U.name, U.idx)
	go U.selsend()
	go U.selrecv()

	return U.cw, U.cr
}

func (U *tuint8) selsend() {
	for {
		b := uint82bytes(<-U.cw)
		for ok := true; ok; ok = (len(U.n) > 0) {
			if !IsClient {
				<-U.n
			}
			U.w <- b
		}
	}
}

func (U *tuint8) selrecv() {
	for {
		U.cr <- bytes2uint8(<-U.r)
	}
}

func finduint8(name string, idx int) (*tuint8, bool) {
	uint8Dict.RLock()
	defer uint8Dict.RUnlock()

	Ui, found := uint8Dict.m[name]
	if !found || idx > len(Ui)-1 {
		return nil, false
	}
	return Ui[idx], true
}
