package sock

func MakeError(name string, buf ...int) Error {
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

	errorDict.Lock()
	defer errorDict.Unlock()

	if errorDict.m == nil {
		errorDict.m = map[string][]*terror{}
	}
	E := &terror{
		name: name,
		len:  buflen,
		idx:  len(errorDict.m[name]),
		sel: sel{
			w: make(chan interface{}, buflen),
			r: make(chan interface{}, buflen),
		},
		c: make(chan error, buflen),
	}
	if !IsClient {
		E.sel.n = make(chan int)
		E.n = make(chan int)
	}
	errorDict.m[E.name] = append(errorDict.m[E.name], E)

	return E.c
}

// func (E *terror) postIfClient(t byte, name string) {
// 	if !IsClient {
// 		return
// 	}
// 	if len(Addr) == 0 || Addr[len(Addr)-1] != '/' {
// 		Addr += "/"
// 	}
// 	for {
// 		pkt := bytes.Join([][]byte{[]byte{t}, []byte(name), <-w}, v)
// 		for {
// 			resp, err := http.Post(Addr+POST, "text/plain", bytes.NewReader(pkt))
// 			if err == nil && resp.StatusCode < 300 {
// 				break
// 			}
// 		}
// 	}
// }

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
