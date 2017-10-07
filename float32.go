package sock

func MakeFloat32(name string, buf ...int) (chan<- float32, <-chan float32) {
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

	float32Dict.Lock()
	if float32Dict.m == nil {
		float32Dict.m = map[string][]*tfloat32{}
	}
	F := &tfloat32{
		name: name,
		len:  buflen,
		idx:  len(float32Dict.m[name]),
		w:    make(chan []byte, buflen),
		r:    make(chan []byte, buflen),
		cw:   make(chan float32, buflen),
		cr:   make(chan float32, buflen),
	}
	if !IsClient {
		F.n = make(chan int)
	}
	float32Dict.m[F.name] = append(float32Dict.m[F.name], F)
	float32Dict.Unlock()

	go wIfClient(F.w, Tfloat32, F.name, F.idx)
	go rIfClient(F.r, Tfloat32, F.name, F.idx)
	go F.selsend()
	go F.selrecv()

	return F.cw, F.cr
}

func (F *tfloat32) selsend() {
	for {
		b := float322bytes(<-F.cw)
		for ok := true; ok; ok = (len(F.n) > 0) {
			if !IsClient {
				<-F.n
			}
			F.w <- b
		}
	}
}

func (F *tfloat32) selrecv() {
	for {
		F.cr <- bytes2float32(<-F.r)
	}
}

func findfloat32(name string, idx int) (*tfloat32, bool) {
	float32Dict.RLock()
	defer float32Dict.RUnlock()

	Fi, found := float32Dict.m[name]
	if !found || idx > len(Fi)-1 {
		return nil, false
	}
	return Fi[idx], true
}

func (F *tfloat32) getfloat32(b []byte) {
	F.r <- b
}

func (F *tfloat32) setfloat32() []byte {
	F.n <- 1
	return <-F.w
}
