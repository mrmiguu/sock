package sock

func MakeString(name string, buf ...int) (chan<- string, <-chan string) {
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

	stringDict.Lock()
	if stringDict.m == nil {
		stringDict.m = map[string][]*tstring{}
	}
	S := &tstring{
		name: name,
		len:  buflen,
		idx:  len(stringDict.m[name]),
		w:    make(chan []byte, buflen),
		r:    make(chan []byte, buflen),
		cw:   make(chan string, buflen),
		cr:   make(chan string, buflen),
	}
	if !IsClient {
		S.n = make(chan int)
	}
	stringDict.m[S.name] = append(stringDict.m[S.name], S)
	stringDict.Unlock()

	go wIfClient(S.w, Tstring, S.name, S.idx)
	go rIfClient(S.r, Tstring, S.name, S.idx)
	go S.selsend()
	go S.selrecv()

	return S.cw, S.cr
}

func (S *tstring) selsend() {
	for {
		b := []byte(<-S.cw)
		for ok := true; ok; ok = (len(S.n) > 0) {
			if !IsClient {
				<-S.n
			}
			S.w <- b
		}
	}
}

func (S *tstring) selrecv() {
	for {
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
