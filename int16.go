package sock

func MakeInt16(name string, buf ...int) (chan<- int16, <-chan int16) {
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

	int16Dict.Lock()
	if int16Dict.m == nil {
		int16Dict.m = map[string][]*tint16{}
	}
	I := &tint16{
		name: name,
		len:  buflen,
		idx:  len(int16Dict.m[name]),
		w:    make(chan []byte, buflen),
		r:    make(chan []byte, buflen),
		cw:   make(chan int16, buflen),
		cr:   make(chan int16, buflen),
	}
	if !IsClient {
		I.n = make(chan int)
	}
	int16Dict.m[I.name] = append(int16Dict.m[I.name], I)
	int16Dict.Unlock()

	go wIfClient(I.w, Tint16, I.name, I.idx)
	go rIfClient(I.r, Tint16, I.name, I.idx)
	go I.selsend()
	go I.selrecv()

	return I.cw, I.cr
}

func (I *tint16) selsend() {
	for {
		b := int162bytes(<-I.cw)
		for ok := true; ok; ok = (len(I.n) > 0) {
			if !IsClient {
				<-I.n
			}
			I.w <- b
		}
	}
}

func (I *tint16) selrecv() {
	for {
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
