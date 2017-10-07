package sock

func MakeBool(name string, buf ...int) (chan<- bool, <-chan bool) {
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

	boolDict.Lock()
	if boolDict.m == nil {
		boolDict.m = map[string][]*tbool{}
	}
	B := &tbool{
		name: name,
		len:  buflen,
		idx:  len(boolDict.m[name]),
		w:    make(chan []byte, buflen),
		r:    make(chan []byte, buflen),
		cw:   make(chan bool, buflen),
		cr:   make(chan bool, buflen),
	}
	if !IsClient {
		B.n = make(chan int)
	}
	boolDict.m[B.name] = append(boolDict.m[B.name], B)
	boolDict.Unlock()

	go wIfClient(B.w, Tbool, B.name, B.idx)
	go rIfClient(B.r, Tbool, B.name, B.idx)
	go B.selsend()
	go B.selrecv()

	return B.cw, B.cr
}

func (B *tbool) selsend() {
	for {
		b := bool2bytes(<-B.cw)
		for ok := true; ok; ok = (len(B.n) > 0) {
			if !IsClient {
				<-B.n
			}
			B.w <- b
		}
	}
}

func (B *tbool) selrecv() {
	for {
		B.cr <- bytes2bool(<-B.r)
	}
}

func findbool(name string, idx int) (*tbool, bool) {
	boolDict.RLock()
	defer boolDict.RUnlock()

	Bi, found := boolDict.m[name]
	if !found || idx > len(Bi)-1 {
		return nil, false
	}
	return Bi[idx], true
}

func (B *tbool) getbool(b []byte) {
	B.r <- b
}

func (B *tbool) setbool() []byte {
	B.n <- 1
	return <-B.w
}
