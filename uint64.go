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

	go started.Do(wAndOrRIfServer)

	uint64Dict.Lock()
	if uint64Dict.m == nil {
		uint64Dict.m = map[string][]*tuint64{}
	}
	U := &tuint64{
		name: name,
		len:  buflen,
		idx:  len(uint64Dict.m[name]),
		w:    make(chan []byte, buflen),
		r:    make(chan []byte, buflen),
		cw:   make(chan uint64, buflen),
		cr:   make(chan uint64, buflen),
	}
	if !IsClient {
		U.n = make(chan int)
	}
	uint64Dict.m[U.name] = append(uint64Dict.m[U.name], U)
	uint64Dict.Unlock()

	go wIfClient(U.w, Tuint64, U.name, U.idx)
	go rIfClient(U.r, Tuint64, U.name, U.idx)
	go U.selsend()
	go U.selrecv()

	return U.cw, U.cr
}

func (U *tuint64) selsend() {
	for {
		b := uint642bytes(<-U.cw)
		for ok := true; ok; ok = (len(U.n) > 0) {
			if !IsClient {
				<-U.n
			}
			U.w <- b
		}
	}
}

func (U *tuint64) selrecv() {
	for {
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

func (U *tuint64) getuint64(b []byte) {
		U.r <- b
}

func (U *tuint64) setuint64() []byte {
	U.n <- 1
	return <-U.w
}
