package sock

func MakeUint32(name string, buf ...int) (chan<- uint32, <-chan uint32) {
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

	uint32Dict.Lock()
	if uint32Dict.m == nil {
		uint32Dict.m = map[string][]*tuint32{}
	}
	U := &tuint32{
		name: name,
		len:  buflen,
		idx:  len(uint32Dict.m[name]),
		w:    make(chan []byte, buflen),
		r:    make(chan []byte, buflen),
		cw:   make(chan uint32, buflen),
		cr:   make(chan uint32, buflen),
	}
	if !IsClient {
		U.n = make(chan int)
	}
	uint32Dict.m[U.name] = append(uint32Dict.m[U.name], U)
	uint32Dict.Unlock()

	go wIfClient(U.w, Tuint32, U.name, U.idx)
	go rIfClient(U.r, Tuint32, U.name, U.idx)
	go U.selsend()
	go U.selrecv()

	return U.cw, U.cr
}

func (U *tuint32) selsend() {
	for {
		b := uint322bytes(<-U.cw)
		for ok := true; ok; ok = (len(U.n) > 0) {
			if !IsClient {
				<-U.n
			}
			U.w <- b
		}
	}
}

func (U *tuint32) selrecv() {
	for {
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
