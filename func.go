package sock

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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

	if _, err := os.Stat(ClientFolder); len(ClientFolder) > 0 && os.IsNotExist(err) {
		panic("client folder not found")
	}
	http.Handle("/", http.FileServer(http.Dir(ClientFolder)))

	http.HandleFunc("/"+POST, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")

		b, err := ioutil.ReadAll(r.Body)
		r.Body.Close()
		if err != nil {
			delayedError(w, http.StatusBadRequest)
			return
		}
		parts := bytes.Split(b, v)
		t, name, idx, body := parts[0][0], string(parts[1]), bytes2int(parts[2]), parts[3]

		switch t {
		case Terror:
			E, ok := finderror(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			E.geterror(body)

		case Tstring:
			S, ok := findstring(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			S.getstring(body)

		case Tint:
			I, ok := findint(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			I.getint(body)

		case Tbool:
			B, ok := findbool(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			B.getbool(body)

		case Tbytes:
			B, ok := findbytes(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			B.getbytes(body)

		case Tfloat64:
			F, ok := findfloat64(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			F.getfloat64(body)

		case Trune:
			R, ok := findrune(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			R.getrune(body)

		case Tint8:
			I, ok := findint8(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			I.getint8(body)

		case Tint16:
			I, ok := findint16(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			I.getint16(body)

		case Tint32:
			I, ok := findint32(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			I.getint32(body)

		case Tint64:
			I, ok := findint64(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			I.getint64(body)

		case Tuint:
			U, ok := finduint(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			U.getuint(body)

		case Tuint8:
			U, ok := finduint8(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			U.getuint8(body)

		case Tuint16:
			U, ok := finduint16(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			U.getuint16(body)

		case Tuint32:
			U, ok := finduint32(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			U.getuint32(body)

		case Tuint64:
			U, ok := finduint64(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			U.getuint64(body)

		case Tbyte:
			B, ok := findbyte(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			B.getbyte(body)

		case Tfloat32:
			F, ok := findfloat32(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			F.getfloat32(body)

		default:
			delayedError(w, http.StatusBadRequest)
			return
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
		t, name, idx := parts[0][0], string(parts[1]), bytes2int(parts[2])

		switch t {
		case Terror:
			E, ok := finderror(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			b = E.seterror()

		case Tstring:
			S, ok := findstring(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			b = S.setstring()

		case Tint:
			I, ok := findint(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			b = I.setint()

		case Tbool:
			B, ok := findbool(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			b = B.setbool()

		case Tbytes:
			B, ok := findbytes(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			b = B.setbytes()

		case Tfloat64:
			F, ok := findfloat64(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			b = F.setfloat64()

		case Trune:
			R, ok := findrune(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			b = R.setrune()

		case Tint8:
			I, ok := findint8(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			b = I.setint8()

		case Tint16:
			I, ok := findint16(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			b = I.setint16()

		case Tint32:
			I, ok := findint32(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			b = I.setint32()

		case Tint64:
			I, ok := findint64(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			b = I.setint64()

		case Tuint:
			U, ok := finduint(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			b = U.setuint()

		case Tuint8:
			U, ok := finduint8(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			b = U.setuint8()

		case Tuint16:
			U, ok := finduint16(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			b = U.setuint16()

		case Tuint32:
			U, ok := finduint32(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			b = U.setuint32()

		case Tuint64:
			U, ok := finduint64(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			b = U.setuint64()

		case Tbyte:
			B, ok := findbyte(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			b = B.setbyte()

		case Tfloat32:
			F, ok := findfloat32(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			b = F.setfloat32()

		default:
			delayedError(w, http.StatusBadRequest)
			return
		}

		w.Write(b)
	})

	log.Fatal(http.ListenAndServe(Addr, nil))
}

func wIfClient(w chan []byte, t byte, name string, idx int) {
	if !IsClient {
		return
	}
	if len(Addr) == 0 || Addr[len(Addr)-1] != '/' {
		Addr += "/"
	}
	for {
		pkt = bytes.Join([][]byte{[]byte{t}, []byte(name), int2bytes(idx), <-w}, v)
		for {
			resp, err := http.Post(Addr+POST, "text/plain", bytes.NewReader(pkt))
			if err == nil && resp.StatusCode < 300 {
				break
			}
		}
	}
}

func rIfClient(r chan []byte, t byte, name string, idx int) {
	if !IsClient {
		return
	}
	if len(Addr) == 0 || Addr[len(Addr)-1] != '/' {
		Addr += "/"
	}
	for {
		pkt = bytes.Join([][]byte{[]byte{t}, []byte(name), int2bytes(idx)}, v)
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
	time.Sleep(ErrorStatusDelay)
	http.Error(w, "", code)
}
