package sock

func MakeError(name string, buf ...int) chan error {
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
		c:    make(chan error, buflen),
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

	return E.c
}

func (E *terror) selsend() {
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
			E.w <- error2bytes(<-E.c)
		}
	}
}

func (E *terror) selrecv() {
	for {
		<-E.selr
		E.c <- bytes2error(<-E.r)
	}
}

// func (E *terror) makeW() {
// 	go postIfClient(E.p.w.c, Terror, E.Name)
// }

// func (E *terror) makeR() {
// 	go getIfClient(E.p.r.c, Terror, E.Name)
// }

// func (E *terror) to(e error) {
// 	if IsClient {
// 		E.p.w.c <- []byte(e.Error())
// 		return
// 	}
// 	for {
// 		<-E.p.n.c
// 		E.p.w.c <- []byte(e.Error())
// 		if len(E.p.n.c) == 0 {
// 			break
// 		}
// 	}
// }

// func (E *terror) from() error {
// 	return errors.New(string(<-E.p.r.c))
// }

// func (E *terror) S() chan<- error {
// 	c := make(chan error, E.Len)
// 	go started.Do(getAndOrPostIfServer)
// 	E.add()
// 	E.p.w.Do(E.makeW)
// 	E.p.n.Do(E.makeNIfServer)
// 	go func() {
// 		E.to(errors.New(""))
// 		i := <-c
// 		close(c)
// 		E.to(i)
// 	}()
// 	return c
// }

// func (E *terror) R() <-chan error {
// 	c := make(chan error, E.Len)
// 	go started.Do(getAndOrPostIfServer)
// 	E.add()
// 	E.p.r.Do(E.makeR)
// 	go func() {
// 		E.from()
// 		c <- E.from()
// 		close(c)
// 	}()
// 	return c
// }
