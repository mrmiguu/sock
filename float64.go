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
		w:    make(chan []byte, buflen),
		r:    make(chan []byte, buflen),
		cw:   make(chan float64, buflen),
		cr:   make(chan float64, buflen),
	}
	if !IsClient {
		F.n = make(chan int)
	}
	float64Dict.m[F.name] = append(float64Dict.m[F.name], F)
	float64Dict.Unlock()

	go wIfClient(F.w, Tfloat64, F.name, F.idx)
	go rIfClient(F.r, Tfloat64, F.name, F.idx)
	go F.selsend()
	go F.selrecv()

	return F.cw, F.cr
}

func (F *tfloat64) selsend() {
	for {
		b := float642bytes(<-F.cw)
		for ok := true; ok; ok = (len(F.n) > 0) {
			if !IsClient {
				<-F.n
			}
			F.w <- b
		}
	}
}

func (F *tfloat64) selrecv() {
	for {
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

func (F *tfloat64) getfloat64(b []byte) {
	F.r <- b
}

func (F *tfloat64) setfloat64() []byte {
	F.n <- 1
	return <-F.w
}
