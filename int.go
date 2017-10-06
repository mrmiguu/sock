package sock

func MakeInt(name string, buf ...int) (chan<- int, <-chan int) {
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

	intDict.Lock()
	if intDict.m == nil {
		intDict.m = map[string][]*tint{}
	}
	I := &tint{
		name: name,
		len:  buflen,
		idx:  len(intDict.m[name]),
		selw: make(chan []byte, buflen),
		selr: make(chan []byte, buflen),
		w:    make(chan []byte, buflen),
		r:    make(chan []byte, buflen),
		cw:   make(chan int, buflen),
		cr:   make(chan int, buflen),
	}
	if !IsClient {
		I.seln = make(chan int)
		I.n = make(chan int)
	}
	intDict.m[I.name] = append(intDict.m[I.name], I)
	intDict.Unlock()

	go wIfClient(I.selw, I.w, Tint, I.name, I.idx)
	go rIfClient(I.selr, I.r, Tint, I.name, I.idx)
	go I.selsend()
	go I.selrecv()

	return I.cw, I.cr
}

func (I *tint) selsend() {
	for {
		for ok := true; ok; ok = (len(I.seln) > 0) {
			if !IsClient {
				<-I.seln
			}
			I.selw <- nil
		}

		b := int2bytes(<-I.cw)
		for ok := true; ok; ok = (len(I.n) > 0) {
			if !IsClient {
				<-I.n
			}
			I.w <- b
		}
	}
}

func (I *tint) selrecv() {
	for {
		<-I.selr
		I.cr <- bytes2int(<-I.r)
	}
}

func findint(name string, idx int) (*tint, bool) {
	intDict.RLock()
	defer intDict.RUnlock()

	Ii, found := intDict.m[name]
	if !found || idx > len(Ii)-1 {
		return nil, false
	}
	return Ii[idx], true
}

func (I *tint) getint(sel byte, b []byte) {
	if sel == 1 {
		I.selr <- nil
	} else {
		I.r <- b
	}
}

func (I *tint) setint(sel byte) []byte {
	if sel == 1 {
		I.seln <- 1
		return <-I.selw
	}
	I.n <- 1
	return <-I.w
}
