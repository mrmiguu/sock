package sock

func MakeInt8(name string, buf ...int) (chan<- int8, <-chan int8) {
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

	int8Dict.Lock()
	if int8Dict.m == nil {
		int8Dict.m = map[string][]*tint8{}
	}
	I := &tint8{
		name: name,
		len:  buflen,
		idx:  len(int8Dict.m[name]),
		selw: make(chan []byte, buflen),
		selr: make(chan []byte, buflen),
		w:    make(chan []byte, buflen),
		r:    make(chan []byte, buflen),
		cw:   make(chan int8, buflen),
		cr:   make(chan int8, buflen),
	}
	if !IsClient {
		I.seln = make(chan int)
		I.n = make(chan int)
	}
	int8Dict.m[I.name] = append(int8Dict.m[I.name], I)
	int8Dict.Unlock()

	go wIfClient(I.selw, Tint8, I.name, I.idx, 1)
	go rIfClient(I.selr, Tint8, I.name, I.idx, 1)
	go wIfClient(I.w, Tint8, I.name, I.idx, 0)
	go rIfClient(I.r, Tint8, I.name, I.idx, 0)
	go I.selsend()
	go I.selrecv()

	return I.cw, I.cr
}

func (I *tint8) selsend() {
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
			I.w <- int82bytes(<-I.cw)
		}
	}
}

func (I *tint8) selrecv() {
	for {
		<-I.selr
		I.cr <- bytes2int8(<-I.r)
	}
}

func findint8(name string, idx int) (*tint8, bool) {
	int8Dict.RLock()
	defer int8Dict.RUnlock()

	Ii, found := int8Dict.m[name]
	if !found || idx > len(Ii)-1 {
		return nil, false
	}
	return Ii[idx], true
}

func (I *tint8) getint8(sel byte, b []byte) {
	if sel == 1 {
		I.selr <- []byte{}
	} else {
		I.r <- b
	}
}

func (I *tint8) setint8(sel byte) []byte {
	if sel == 1 {
		I.seln <- 1
		<-I.selw
		return []byte{}
	}
	I.n <- 1
	return <-I.w
}
