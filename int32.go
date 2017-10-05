package sock

func MakeInt32(name string, buf ...int) (chan<- int32, <-chan int32) {
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

	int32Dict.Lock()
	if int32Dict.m == nil {
		int32Dict.m = map[string][]*tint32{}
	}
	I := &tint32{
		name: name,
		len:  buflen,
		idx:  len(int32Dict.m[name]),
		selw: make(chan []byte, buflen),
		selr: make(chan []byte, buflen),
		w:    make(chan []byte, buflen),
		r:    make(chan []byte, buflen),
		cw:   make(chan int32, buflen),
		cr:   make(chan int32, buflen),
	}
	if !IsClient {
		I.seln = make(chan int)
		I.n = make(chan int)
	}
	int32Dict.m[I.name] = append(int32Dict.m[I.name], I)
	int32Dict.Unlock()

	go wIfClient(I.selw, Tint32, I.name, I.idx, 1)
	go rIfClient(I.selr, Tint32, I.name, I.idx, 1)
	go wIfClient(I.w, Tint32, I.name, I.idx, 0)
	go rIfClient(I.r, Tint32, I.name, I.idx, 0)
	go I.selsend()
	go I.selrecv()

	return I.cw, I.cr
}

func (I *tint32) selsend() {
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
			I.w <- int322bytes(<-I.cw)
		}
	}
}

func (I *tint32) selrecv() {
	for {
		<-I.selr
		I.cr <- bytes2int32(<-I.r)
	}
}

func findint32(name string, idx int) (*tint32, bool) {
	int32Dict.RLock()
	defer int32Dict.RUnlock()

	Ii, found := int32Dict.m[name]
	if !found || idx > len(Ii)-1 {
		return nil, false
	}
	return Ii[idx], true
}

func (I *tint32) getint32(sel byte, b []byte) {
	if sel == 1 {
		I.selr <- []byte{}
	} else {
		I.r <- b
	}
}

func (I *tint32) setint32(sel byte) []byte {
	if sel == 1 {
		I.seln <- 1
		<-I.selw
		return []byte{}
	}
	I.n <- 1
	return <-I.w
}