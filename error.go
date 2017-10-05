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

	go wAndOrRIfServer()

	errorDict.Lock()
	if errorDict.m == nil {
		errorDict.m = map[string][]*terror{}
	}
	E := &terror{
		name: name,
		len:  buflen,
		idx:  len(errorDict.m[name]),
		selw: make(chan []byte, buflen),
		selr: make(chan []byte, buflen),
		w:    make(chan []byte, buflen),
		r:    make(chan []byte, buflen),
		cw:   make(chan error, buflen),
		cr:   make(chan error, buflen),
	}
	if !IsClient {
		E.seln = make(chan int)
		E.n = make(chan int)
	}
	errorDict.m[E.name] = append(errorDict.m[E.name], E)
	errorDict.Unlock()

	go wIfClient(E.selw, Terror, E.name, E.idx, 1)
	go rIfClient(E.selr, Terror, E.name, E.idx, 1)
	go wIfClient(E.w, Terror, E.name, E.idx, 0)
	go rIfClient(E.r, Terror, E.name, E.idx, 0)
	go E.selsend()
	go E.selrecv()

	return E.cw, E.cr
}

func (E *terror) selsend() {
	// defer func() { recover() }()
	for {
		for ok := true; ok; ok = (len(E.seln) > 0) {
			if !IsClient {
				<-E.seln
			}
			E.selw <- nil
		}

		for ok := true; ok; ok = (len(E.n) > 0) {
			if !IsClient {
				<-E.n
			}
			E.w <- error2bytes(<-E.cw)
		}
	}
}

func (E *terror) selrecv() {
	// defer func() {
	// 	done := load.New("closing " + E.name + "#" + strconv.Itoa(E.idx))
	// 	recover()

	// 	errorDict.Lock()
	// 	Ei := errorDict.m[E.name]
	// 	if len(Ei) == 1 {
	// 		delete(errorDict.m, E.name)
	// 	} else {
	// 		errorDict.m[E.name] = append(Ei[:E.idx], Ei[E.idx+1:]...)
	// 	}
	// 	E := Ei[E.idx]
	// 	close(E.selw)
	// 	close(E.selr)
	// 	close(E.w)
	// 	close(E.r)
	// 	if !IsClient {
	// 		close(E.seln)
	// 		close(E.n)
	// 	}
	// 	errorDict.Unlock()

	// 	done <- true
	// }()

	for {
		<-E.selr
		E.cr <- bytes2error(<-E.r)
	}
}
