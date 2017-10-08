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
			get(E.r, body)

		case Tstring:
			S, ok := findstring(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			get(S.r, body)

		case Tint:
			I, ok := findint(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			get(I.r, body)

		case Tbool:
			B, ok := findbool(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			get(B.r, body)

		case Tbytes:
			B, ok := findbytes(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			get(B.r, body)

		case Tfloat64:
			F, ok := findfloat64(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			get(F.r, body)

		case Trune:
			R, ok := findrune(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			get(R.r, body)

		case Tint8:
			I, ok := findint8(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			get(I.r, body)

		case Tint16:
			I, ok := findint16(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			get(I.r, body)

		case Tint32:
			I, ok := findint32(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			get(I.r, body)

		case Tint64:
			I, ok := findint64(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			get(I.r, body)

		case Tuint:
			U, ok := finduint(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			get(U.r, body)

		case Tuint8:
			U, ok := finduint8(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			get(U.r, body)

		case Tuint16:
			U, ok := finduint16(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			get(U.r, body)

		case Tuint32:
			U, ok := finduint32(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			get(U.r, body)

		case Tuint64:
			U, ok := finduint64(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			get(U.r, body)

		case Tbyte:
			B, ok := findbyte(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			get(B.r, body)

		case Tfloat32:
			F, ok := findfloat32(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			get(F.r, body)

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
			b = set(E.n, E.w)

		case Tstring:
			S, ok := findstring(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			b = set(S.n, S.w)

		case Tint:
			I, ok := findint(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			b = set(I.n, I.w)

		case Tbool:
			B, ok := findbool(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			b = set(B.n, B.w)

		case Tbytes:
			B, ok := findbytes(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			b = set(B.n, B.w)

		case Tfloat64:
			F, ok := findfloat64(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			b = set(F.n, F.w)

		case Trune:
			R, ok := findrune(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			b = set(R.n, R.w)

		case Tint8:
			I, ok := findint8(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			b = set(I.n, I.w)

		case Tint16:
			I, ok := findint16(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			b = set(I.n, I.w)

		case Tint32:
			I, ok := findint32(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			b = set(I.n, I.w)

		case Tint64:
			I, ok := findint64(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			b = set(I.n, I.w)

		case Tuint:
			U, ok := finduint(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			b = set(U.n, U.w)

		case Tuint8:
			U, ok := finduint8(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			b = set(U.n, U.w)

		case Tuint16:
			U, ok := finduint16(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			b = set(U.n, U.w)

		case Tuint32:
			U, ok := finduint32(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			b = set(U.n, U.w)

		case Tuint64:
			U, ok := finduint64(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			b = set(U.n, U.w)

		case Tbyte:
			B, ok := findbyte(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			b = set(B.n, B.w)

		case Tfloat32:
			F, ok := findfloat32(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			b = set(F.n, F.w)

		default:
			delayedError(w, http.StatusBadRequest)
			return
		}

		w.Write(b)
	})

	log.Fatal(http.ListenAndServe(Addr, nil))
}

func get(r chan []byte, b []byte) {
	r <- b
}

func set(n chan int, w chan []byte) []byte {
	n <- 1
	return <-w
}

func wIfClient(w chan []byte, t byte, name string, idx int) {
	if !IsClient {
		return
	}
	if len(Addr) == 0 || Addr[len(Addr)-1] != '/' {
		Addr += "/"
	}
	for {
		pkt := bytes.Join([][]byte{[]byte{t}, []byte(name), int2bytes(idx), <-w}, v)
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
		pkt := bytes.Join([][]byte{[]byte{t}, []byte(name), int2bytes(idx)}, v)
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
