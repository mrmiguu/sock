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

	go started.Do(wAndOrRIfServer)

	int32Dict.Lock()
	if int32Dict.m == nil {
		int32Dict.m = map[string][]*tint32{}
	}
	I := &tint32{
		name: name,
		len:  buflen,
		idx:  len(int32Dict.m[name]),
		w:    make(chan []byte, buflen),
		r:    make(chan []byte, buflen),
		cw:   make(chan int32, buflen),
		cr:   make(chan int32, buflen),
	}
	if !IsClient {
		I.n = make(chan int)
	}
	int32Dict.m[I.name] = append(int32Dict.m[I.name], I)
	int32Dict.Unlock()

	go wIfClient(I.w, Tint32, I.name, I.idx)
	go rIfClient(I.r, Tint32, I.name, I.idx)
	go I.selsend()
	go I.selrecv()

	return I.cw, I.cr
}

func (I *tint32) selsend() {
	for {
		b := int322bytes(<-I.cw)
		for ok := true; ok; ok = (len(I.n) > 0) {
			if !IsClient {
				<-I.n
			}
			I.w <- b
		}
	}
}

func (I *tint32) selrecv() {
	for {
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

func (I *tint32) getint32(b []byte) {
		I.r <- b
}

func (I *tint32) setint32() []byte {
	I.n <- 1
	return <-I.w
}
