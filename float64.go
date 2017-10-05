package sock

func MakeFloat64(name string, buf ...int) (chan<- float64, <-chan float64) {
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

	float64Dict.Lock()
	if float64Dict.m == nil {
		float64Dict.m = map[string][]*tfloat64{}
	}
	F := &tfloat64{
		name: name,
		len:  buflen,
		idx:  len(float64Dict.m[name]),
		selw: make(chan []byte, buflen),
		selr: make(chan []byte, buflen),
		w:    make(chan []byte, buflen),
		r:    make(chan []byte, buflen),
		cw:   make(chan float64, buflen),
		cr:   make(chan float64, buflen),
	}
	if !IsClient {
		F.seln = make(chan int)
		F.n = make(chan int)
	}
	float64Dict.m[F.name] = append(float64Dict.m[F.name], F)
	float64Dict.Unlock()

	go wIfClient(F.selw, Tfloat64, F.name, F.idx, 1)
	go rIfClient(F.selr, Tfloat64, F.name, F.idx, 1)
	go wIfClient(F.w, Tfloat64, F.name, F.idx, 0)
	go rIfClient(F.r, Tfloat64, F.name, F.idx, 0)
	go F.selsend()
	go F.selrecv()

	return F.cw, F.cr
}

func (F *tfloat64) selsend() {
	for {
		for ok := true; ok; ok = (len(F.seln) > 0) {
			if !IsClient {
				<-F.seln
			}
			F.selw <- nil
		}

		for ok := true; ok; ok = (len(F.n) > 0) {
			if !IsClient {
				<-F.n
			}
			F.w <- float642bytes(<-F.cw)
		}
	}
}

func (F *tfloat64) selrecv() {
	for {
		<-F.selr
		F.cr <- bytes2float64(<-F.r)
	}
}

func findfloat64(name string, idx int) (*tfloat64, bool) {
	float64Dict.RLock()
	defer float64Dict.RUnlock()

	Fi, found := float64Dict.m[name]
	if !found || idx > len(Fi)-1 {
		return nil, false
	}
	return Fi[idx], true
}

func (F *tfloat64) getfloat64(sel byte, b []byte) {
	if sel == 1 {
		F.selr <- []byte{}
	} else {
		F.r <- b
	}
}

func (F *tfloat64) setfloat64(sel byte) []byte {
	if sel == 1 {
		F.seln <- 1
		<-F.selw
		return []byte{}
	}
	F.n <- 1
	return <-F.w
}
