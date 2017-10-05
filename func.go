package sock

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func init() {
	if IsClient {
		Addr = DefaultClientAddr
	} else {
		Addr = DefaultServerAddr
	}
}

func wAndOrRIfServer() {
	if IsClient {
		return
	}

	// consider commenting out? idk
	http.Handle("/", http.FileServer(http.Dir("client")))

	http.HandleFunc("/"+POST, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")

		b, err := ioutil.ReadAll(r.Body)
		r.Body.Close()
		if err != nil {
			delayedError(w, http.StatusBadRequest)
			return
		}
		parts := bytes.Split(b, v)
		t, name, idx, sel, body := parts[0][0], string(parts[1]), bytes2int(parts[2]), parts[3][0], parts[4]

		// defer func() { recover() }()

		switch t {
		case Terror:
			E, ok := finderror(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			E.geterror(sel, body)

		case Tstring:
			S, ok := findstring(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			S.getstring(sel, body)

		case Tint:
			I, ok := findint(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			I.getint(sel, body)

		case Tbool:
			B, ok := findbool(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			B.getbool(sel, body)

		case Tbytes:
			B, ok := findbytes(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			B.getbytes(sel, body)

		case Tfloat64:
			F, ok := findfloat64(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			F.getfloat64(sel, body)

		case Trune:
			R, ok := findrune(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			R.getrune(sel, body)

		case Tint8:
			I, ok := findint8(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			I.getint8(sel, body)

		case Tint16:
			I, ok := findint16(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			I.getint16(sel, body)

		case Tint32:
			I, ok := findint32(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			I.getint32(sel, body)

		case Tint64:
			I, ok := findint64(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			I.getint64(sel, body)

		case Tuint:
			U, ok := finduint(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			U.getuint(sel, body)

		case Tuint8:
			U, ok := finduint8(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			U.getuint8(sel, body)

		case Tuint16:
			U, ok := finduint16(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			U.getuint16(sel, body)

		case Tuint32:
			U, ok := finduint32(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			U.getuint32(sel, body)

		case Tuint64:
			U, ok := finduint64(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			U.getuint64(sel, body)

		case Tbyte:
			B, ok := findbyte(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			B.getbyte(sel, body)

		case Tfloat32:
			F, ok := findfloat32(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			F.getfloat32(sel, body)

		case Tcomplex64:
			C, ok := findcomplex64(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			C.getcomplex64(sel, body)

		case Tcomplex128:
			C, ok := findcomplex128(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			C.getcomplex128(sel, body)

		default:
			delayedError(w, http.StatusBadRequest)
		}
	})

	http.HandleFunc("/"+GET, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")

		b, err := ioutil.ReadAll(r.Body)
		r.Body.Close()
		if err != nil {
			delayedError(w, http.StatusBadRequest)
			return
		}
		parts := bytes.Split(b, v)
		t, name, idx, sel := parts[0][0], string(parts[1]), bytes2int(parts[2]), parts[3][0]

		// defer func() { recover() }()

		switch t {
		case Terror:
			E, ok := finderror(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			b = E.seterror(sel)

		case Tstring:
			S, ok := findstring(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			b = S.setstring(sel)

		case Tint:
			I, ok := findint(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			b = I.setint(sel)

		case Tbool:
			B, ok := findbool(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			b = B.setbool(sel)

		case Tbytes:
			B, ok := findbytes(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			b = B.setbytes(sel)

		case Tfloat64:
			F, ok := findfloat64(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			b = F.setfloat64(sel)

		case Trune:
			R, ok := findrune(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			b = R.setrune(sel)

		case Tint8:
			I, ok := findint8(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			b = I.setint8(sel)

		case Tint16:
			I, ok := findint16(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			b = I.setint16(sel)

		case Tint32:
			I, ok := findint32(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			b = I.setint32(sel)

		case Tint64:
			I, ok := findint64(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			b = I.setint64(sel)

		case Tuint:
			U, ok := finduint(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			b = U.setuint(sel)

		case Tuint8:
			U, ok := finduint8(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			b = U.setuint8(sel)

		case Tuint16:
			U, ok := finduint16(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			b = U.setuint16(sel)

		case Tuint32:
			U, ok := finduint32(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			b = U.setuint32(sel)

		case Tuint64:
			U, ok := finduint64(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			b = U.setuint64(sel)

		case Tbyte:
			B, ok := findbyte(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			b = B.setbyte(sel)

		case Tfloat32:
			F, ok := findfloat32(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			b = F.setfloat32(sel)

		case Tcomplex64:
			C, ok := findcomplex64(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			b = C.setcomplex64(sel)

		case Tcomplex128:
			C, ok := findcomplex128(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			b = C.setcomplex128(sel)

		default:
			delayedError(w, http.StatusBadRequest)
			return
		}

		w.Write(b)
	})

	log.Fatal(http.ListenAndServe(Addr, nil))
}

func wIfClient(w chan []byte, t byte, name string, idx int, sel byte) {
	if !IsClient {
		return
	}
	if len(Addr) == 0 || Addr[len(Addr)-1] != '/' {
		Addr += "/"
	}
	for {
		pkt := bytes.Join([][]byte{[]byte{t}, []byte(name), int2bytes(idx), []byte{sel}, <-w}, v)
		for {
			resp, err := http.Post(Addr+POST, "text/plain", bytes.NewReader(pkt))
			if err == nil && resp.StatusCode < 300 {
				break
			}
		}
	}
}

func rIfClient(r chan []byte, t byte, name string, idx int, sel byte) {
	if !IsClient {
		return
	}
	if len(Addr) == 0 || Addr[len(Addr)-1] != '/' {
		Addr += "/"
	}
	// defer func() { recover() }()
	for {
		pkt := bytes.Join([][]byte{[]byte{t}, []byte(name), int2bytes(idx), []byte{sel}}, v)
		for {
			resp, err := http.Post(Addr+GET, "text/plain", bytes.NewReader(pkt))
			if err != nil || resp.StatusCode > 299 {
				continue
			}
			b, err := ioutil.ReadAll(resp.Body)
			resp.Body.Close()
			if err == nil {
				r <- b
				break
			}
		}
	}
}

func delayedError(w http.ResponseWriter, code int) {
	time.Sleep(ErrorDelay)
	http.Error(w, "", code)
}
