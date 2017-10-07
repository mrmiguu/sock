package sock

func MakeError(name string, buf ...int) (chan<- error, <-chan error) {
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

	errorDict.Lock()
	if errorDict.m == nil {
		errorDict.m = map[string][]*terror{}
	}
	E := &terror{
		name: name,
		len:  buflen,
		idx:  len(errorDict.m[name]),
		w:    make(chan []byte, buflen),
		r:    make(chan []byte, buflen),
		cw:   make(chan error, buflen),
		cr:   make(chan error, buflen),
	}
	if !IsClient {
		E.n = make(chan int)
	}
	errorDict.m[E.name] = append(errorDict.m[E.name], E)
	errorDict.Unlock()

	go wIfClient(E.w, Terror, E.name, E.idx)
	go rIfClient(E.r, Terror, E.name, E.idx)
	go E.selsend()
	go E.selrecv()

	return E.cw, E.cr
}

func (E *terror) selsend() {
	for {
		b := error2bytes(<-E.cw)
		for ok := true; ok; ok = (len(E.n) > 0) {
			if !IsClient {
				<-E.n
			}
			E.w <- b
		}
	}
}

func (E *terror) selrecv() {
	for {
		E.cr <- bytes2error(<-E.r)
	}
}

func finderror(name string, idx int) (*terror, bool) {
	errorDict.RLock()
	defer errorDict.RUnlock()

	Ei, found := errorDict.m[name]
	if !found || idx > len(Ei)-1 {
		return nil, false
	}
	return Ei[idx], true
}

func (E *terror) geterror(b []byte) {
	E.r <- b
}

func (E *terror) seterror() []byte {
	E.n <- 1
	return <-E.w
}
