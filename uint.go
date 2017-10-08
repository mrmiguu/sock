package sock

func MakeUint(name string, buf ...int) (chan<- uint, <-chan uint) {
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

	uintDict.Lock()
	if uintDict.m == nil {
		uintDict.m = map[string][]*tuint{}
	}
	U := &tuint{
		name: name,
		len:  buflen,
		idx:  len(uintDict.m[name]),
		w:    make(chan []byte, buflen),
		r:    make(chan []byte, buflen),
		cw:   make(chan uint, buflen),
		cr:   make(chan uint, buflen),
	}
	if !IsClient {
		U.n = make(chan int)
	}
	uintDict.m[U.name] = append(uintDict.m[U.name], U)
	uintDict.Unlock()

	go wIfClient(U.w, Tuint, U.name, U.idx)
	go rIfClient(U.r, Tuint, U.name, U.idx)
	go U.selsend()
	go U.selrecv()

	return U.cw, U.cr
}

func (U *tuint) selsend() {
	for {
		b := uint2bytes(<-U.cw)
		for ok := true; ok; ok = (len(U.n) > 0) {
			if !IsClient {
				<-U.n
			}
			U.w <- b
		}
	}
}

func (U *tuint) selrecv() {
	for {
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
