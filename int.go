package sock

func MakeInt(name string, buf ...int) (chan<- int, <-chan int) {
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

	intDict.Lock()
	if intDict.m == nil {
		intDict.m = map[string][]*tint{}
	}
	I := &tint{
		name: name,
		len:  buflen,
		idx:  len(intDict.m[name]),
		w:    make(chan []byte, buflen),
		r:    make(chan []byte, buflen),
		cw:   make(chan int, buflen),
		cr:   make(chan int, buflen),
	}
	if !IsClient {
		I.n = make(chan int)
	}
	intDict.m[I.name] = append(intDict.m[I.name], I)
	intDict.Unlock()

	go wIfClient(I.w, Tint, I.name, I.idx)
	go rIfClient(I.r, Tint, I.name, I.idx)
	go I.selsend()
	go I.selrecv()

	return I.cw, I.cr
}

func (I *tint) selsend() {
	for {
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
