package sock

func MakeInt16(name string, buf ...int) (chan<- int16, <-chan int16) {
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

	int16Dict.Lock()
	if int16Dict.m == nil {
		int16Dict.m = map[string][]*tint16{}
	}
	I := &tint16{
		name: name,
		len:  buflen,
		idx:  len(int16Dict.m[name]),
		selw: make(chan []byte, buflen),
		selr: make(chan []byte, buflen),
		w:    make(chan []byte, buflen),
		r:    make(chan []byte, buflen),
		cw:   make(chan int16, buflen),
		cr:   make(chan int16, buflen),
	}
	if !IsClient {
		I.seln = make(chan int)
		I.n = make(chan int)
	}
	int16Dict.m[I.name] = append(int16Dict.m[I.name], I)
	int16Dict.Unlock()

	go wIfClient(I.selw, Tint16, I.name, I.idx, 1)
	go rIfClient(I.selr, Tint16, I.name, I.idx, 1)
	go wIfClient(I.w, Tint16, I.name, I.idx, 0)
	go rIfClient(I.r, Tint16, I.name, I.idx, 0)
	go I.selsend()
	go I.selrecv()

	return I.cw, I.cr
}

func (I *tint16) selsend() {
	for {
		for ok := true; ok; ok = (len(I.seln) > 0) {
			if !IsClient {
				<-I.seln
			}
			I.selw <- nil
		}

		for ok := true; ok; ok = (len(I.n) > 0) {
			if !IsClient {
				<-I.n
			}
			I.w <- int162bytes(<-I.cw)
		}
	}
}

func (I *tint16) selrecv() {
	for {
		<-I.selr
		I.cr <- bytes2int16(<-I.r)
	}
}

func findint16(name string, idx int) (*tint16, bool) {
	int16Dict.RLock()
	defer int16Dict.RUnlock()

	Ii, found := int16Dict.m[name]
	if !found || idx > len(Ii)-1 {
		return nil, false
	}
	return Ii[idx], true
}

func (I *tint16) getint16(sel byte, b []byte) {
	if sel == 1 {
		I.selr <- []byte{}
	} else {
		I.r <- b
	}
}

func (I *tint16) setint16(sel byte) []byte {
	if sel == 1 {
		I.seln <- 1
		<-I.selw
		return []byte{}
	}
	I.n <- 1
	return <-I.w
}
