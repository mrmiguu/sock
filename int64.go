package sock

func MakeInt64(name string, buf ...int) (chan<- int64, <-chan int64) {
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

	int64Dict.Lock()
	if int64Dict.m == nil {
		int64Dict.m = map[string][]*tint64{}
	}
	I := &tint64{
		name: name,
		len:  buflen,
		idx:  len(int64Dict.m[name]),
		w:    make(chan []byte, buflen),
		r:    make(chan []byte, buflen),
		cw:   make(chan int64, buflen),
		cr:   make(chan int64, buflen),
	}
	if !IsClient {
		I.n = make(chan int)
	}
	int64Dict.m[I.name] = append(int64Dict.m[I.name], I)
	int64Dict.Unlock()

	go wIfClient(I.w, Tint64, I.name, I.idx)
	go rIfClient(I.r, Tint64, I.name, I.idx)
	go I.selsend()
	go I.selrecv()

	return I.cw, I.cr
}

func (I *tint64) selsend() {
	for {
		b := int642bytes(<-I.cw)
		for ok := true; ok; ok = (len(I.n) > 0) {
			if !IsClient {
				<-I.n
			}
			I.w <- b
		}
	}
}

func (I *tint64) selrecv() {
	for {
		I.cr <- bytes2int64(<-I.r)
	}
}

func findint64(name string, idx int) (*tint64, bool) {
	int64Dict.RLock()
	defer int64Dict.RUnlock()

	Ii, found := int64Dict.m[name]
	if !found || idx > len(Ii)-1 {
		return nil, false
	}
	return Ii[idx], true
}

func (I *tint64) getint64(b []byte) {
		I.r <- b
}

func (I *tint64) setint64() []byte {
	I.n <- 1
	return <-I.w
}
