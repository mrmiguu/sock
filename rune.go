package sock

func MakeRune(name string, buf ...int) (chan<- rune, <-chan rune) {
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

	runeDict.Lock()
	if runeDict.m == nil {
		runeDict.m = map[string][]*trune{}
	}
	R := &trune{
		name: name,
		len:  buflen,
		idx:  len(runeDict.m[name]),
		selw: make(chan []byte, buflen),
		selr: make(chan []byte, buflen),
		w:    make(chan []byte, buflen),
		r:    make(chan []byte, buflen),
		cw:   make(chan rune, buflen),
		cr:   make(chan rune, buflen),
	}
	if !IsClient {
		R.seln = make(chan int)
		R.n = make(chan int)
	}
	runeDict.m[R.name] = append(runeDict.m[R.name], R)
	runeDict.Unlock()

	go wIfClient(R.selw, R.w, Trune, R.name, R.idx)
	go rIfClient(R.selr, R.r, Trune, R.name, R.idx)
	go R.selsend()
	go R.selrecv()

	return R.cw, R.cr
}

func (R *trune) selsend() {
	// defer func() { recover() }()
	for {
		for ok := true; ok; ok = (len(R.seln) > 0) {
			if !IsClient {
				<-R.seln
			}
			R.selw <- nil
		}

		b := rune2bytes(<-R.cw)
		for ok := true; ok; ok = (len(R.n) > 0) {
			if !IsClient {
				<-R.n
			}
			R.w <- b
		}
	}
}

func (R *trune) selrecv() {
	// defer func() {
	// 	done := load.New("closing " + R.name + "#" + strconv.Itoa(R.idx))
	// 	recover()

	// 	runeDict.Lock()
	// 	Ri := runeDict.m[R.name]
	// 	if len(Ri) == 1 {
	// 		delete(runeDict.m, R.name)
	// 	} else {
	// 		runeDict.m[R.name] = append(Ri[:R.idx], Ri[R.idx+1:]...)
	// 	}
	// 	R := Ri[R.idx]
	// 	close(R.selw)
	// 	close(R.selr)
	// 	close(R.w)
	// 	close(R.r)
	// 	if !IsClient {
	// 		close(R.seln)
	// 		close(R.n)
	// 	}
	// 	runeDict.Unlock()

	// 	done <- true
	// }()

	for {
		<-R.selr
		R.cr <- bytes2rune(<-R.r)
	}
}

func findrune(name string, idx int) (*trune, bool) {
	runeDict.RLock()
	defer runeDict.RUnlock()

	Ri, found := runeDict.m[name]
	if !found || idx > len(Ri)-1 {
		return nil, false
	}
	return Ri[idx], true
}

func (R *trune) getrune(sel byte, b []byte) {
	if sel == 1 {
		R.selr <- nil
	} else {
		R.r <- b
	}
}

func (R *trune) setrune(sel byte) []byte {
	if sel == 1 {
		R.seln <- 1
		return <-R.selw
	}
	R.n <- 1
	return <-R.w
}
