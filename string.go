package sock

func MakeString(name string, buf ...int) (chan<- string, <-chan string) {
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

	stringDict.Lock()
	if stringDict.m == nil {
		stringDict.m = map[string][]*tstring{}
		println(`map[string][]*tstring // created!`)
	}
	S := &tstring{
		name: name,
		len:  buflen,
		idx:  len(stringDict.m[name]),
		selw: make(chan []byte, buflen),
		selr: make(chan []byte, buflen),
		w:    make(chan []byte, buflen),
		r:    make(chan []byte, buflen),
		cw:   make(chan string, buflen),
		cr:   make(chan string, buflen),
	}
	if !IsClient {
		S.seln = make(chan int)
		S.n = make(chan int)
	}
	stringDict.m[S.name] = append(stringDict.m[S.name], S)
	stringDict.Unlock()

	go wIfClient(S.selw, Tstring, S.name, S.idx, 1)
	go rIfClient(S.selr, Tstring, S.name, S.idx, 1)
	go wIfClient(S.w, Tstring, S.name, S.idx, 0)
	go rIfClient(S.r, Tstring, S.name, S.idx, 0)
	go S.selsend()
	go S.selrecv()

	return S.cw, S.cr
}

func (S *tstring) selsend() {
	for {
		for ok := true; ok; ok = (len(S.seln) > 0) {
			if !IsClient {
				<-S.seln
			}
			S.selw <- nil
		}

		for ok := true; ok; ok = (len(S.n) > 0) {
			if !IsClient {
				<-S.n
			}
			S.w <- []byte(<-S.cw)
		}
	}
}

func (S *tstring) selrecv() {
	for {
		<-S.selr
		S.cr <- string(<-S.r)
	}
}

func findstring(name string, idx int) (*tstring, bool) {
	stringDict.RLock()
	defer stringDict.RUnlock()

	Si, found := stringDict.m[name]
	if !found || idx > len(Si)-1 {
		return nil, false
	}
	return Si[idx], true
}

func (S *tstring) getstring(sel byte, b []byte) {
	if sel == 1 {
		S.selr <- nil
	} else {
		S.r <- b
	}
}

func (S *tstring) setstring(sel byte) []byte {
	if sel == 1 {
		S.seln <- 1
		return <-S.selw
	}
	S.n <- 1
	return <-S.w
}
