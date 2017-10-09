package sock

import (
	"bytes"
	"errors"
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
			err := get(E.r, body)
			if err != nil {
				http.Error(w, "", http.StatusRequestTimeout)
				return
			}

		case Tstring:
			S, ok := findstring(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			err := get(S.r, body)
			if err != nil {
				http.Error(w, "", http.StatusRequestTimeout)
				return
			}

		case Tint:
			I, ok := findint(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			err := get(I.r, body)
			if err != nil {
				http.Error(w, "", http.StatusRequestTimeout)
				return
			}

		case Tbool:
			B, ok := findbool(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			err := get(B.r, body)
			if err != nil {
				http.Error(w, "", http.StatusRequestTimeout)
				return
			}

		case Tbytes:
			B, ok := findbytes(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			err := get(B.r, body)
			if err != nil {
				http.Error(w, "", http.StatusRequestTimeout)
				return
			}

		case Tfloat64:
			F, ok := findfloat64(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			err := get(F.r, body)
			if err != nil {
				http.Error(w, "", http.StatusRequestTimeout)
				return
			}

		case Trune:
			R, ok := findrune(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			err := get(R.r, body)
			if err != nil {
				http.Error(w, "", http.StatusRequestTimeout)
				return
			}

		case Tint8:
			I, ok := findint8(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			err := get(I.r, body)
			if err != nil {
				http.Error(w, "", http.StatusRequestTimeout)
				return
			}

		case Tint16:
			I, ok := findint16(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			err := get(I.r, body)
			if err != nil {
				http.Error(w, "", http.StatusRequestTimeout)
				return
			}

		case Tint32:
			I, ok := findint32(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			err := get(I.r, body)
			if err != nil {
				http.Error(w, "", http.StatusRequestTimeout)
				return
			}

		case Tint64:
			I, ok := findint64(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			err := get(I.r, body)
			if err != nil {
				http.Error(w, "", http.StatusRequestTimeout)
				return
			}

		case Tuint:
			U, ok := finduint(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			err := get(U.r, body)
			if err != nil {
				http.Error(w, "", http.StatusRequestTimeout)
				return
			}

		case Tuint8:
			U, ok := finduint8(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			err := get(U.r, body)
			if err != nil {
				http.Error(w, "", http.StatusRequestTimeout)
				return
			}

		case Tuint16:
			U, ok := finduint16(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			err := get(U.r, body)
			if err != nil {
				http.Error(w, "", http.StatusRequestTimeout)
				return
			}

		case Tuint32:
			U, ok := finduint32(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			err := get(U.r, body)
			if err != nil {
				http.Error(w, "", http.StatusRequestTimeout)
				return
			}

		case Tuint64:
			U, ok := finduint64(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			err := get(U.r, body)
			if err != nil {
				http.Error(w, "", http.StatusRequestTimeout)
				return
			}

		case Tbyte:
			B, ok := findbyte(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			err := get(B.r, body)
			if err != nil {
				http.Error(w, "", http.StatusRequestTimeout)
				return
			}

		case Tfloat32:
			F, ok := findfloat32(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			err := get(F.r, body)
			if err != nil {
				http.Error(w, "", http.StatusRequestTimeout)
				return
			}

		default:
			delayedError(w, http.StatusBadRequest)
			return
		}
	})

	go http.HandleFunc("/"+GET, func(w http.ResponseWriter, r *http.Request) {
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
			b, err = set(E.n, E.w)
			if err != nil {
				http.Error(w, "", http.StatusRequestTimeout)
				return
			}

		case Tstring:
			S, ok := findstring(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			b, err = set(S.n, S.w)
			if err != nil {
				http.Error(w, "", http.StatusRequestTimeout)
				return
			}

		case Tint:
			I, ok := findint(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			b, err = set(I.n, I.w)
			if err != nil {
				http.Error(w, "", http.StatusRequestTimeout)
				return
			}

		case Tbool:
			B, ok := findbool(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			b, err = set(B.n, B.w)
			if err != nil {
				http.Error(w, "", http.StatusRequestTimeout)
				return
			}

		case Tbytes:
			B, ok := findbytes(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			b, err = set(B.n, B.w)
			if err != nil {
				http.Error(w, "", http.StatusRequestTimeout)
				return
			}

		case Tfloat64:
			F, ok := findfloat64(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			b, err = set(F.n, F.w)
			if err != nil {
				http.Error(w, "", http.StatusRequestTimeout)
				return
			}

		case Trune:
			R, ok := findrune(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			b, err = set(R.n, R.w)
			if err != nil {
				http.Error(w, "", http.StatusRequestTimeout)
				return
			}

		case Tint8:
			I, ok := findint8(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			b, err = set(I.n, I.w)
			if err != nil {
				http.Error(w, "", http.StatusRequestTimeout)
				return
			}

		case Tint16:
			I, ok := findint16(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			b, err = set(I.n, I.w)
			if err != nil {
				http.Error(w, "", http.StatusRequestTimeout)
				return
			}

		case Tint32:
			I, ok := findint32(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			b, err = set(I.n, I.w)
			if err != nil {
				http.Error(w, "", http.StatusRequestTimeout)
				return
			}

		case Tint64:
			I, ok := findint64(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			b, err = set(I.n, I.w)
			if err != nil {
				http.Error(w, "", http.StatusRequestTimeout)
				return
			}

		case Tuint:
			U, ok := finduint(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			b, err = set(U.n, U.w)
			if err != nil {
				http.Error(w, "", http.StatusRequestTimeout)
				return
			}

		case Tuint8:
			U, ok := finduint8(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			b, err = set(U.n, U.w)
			if err != nil {
				http.Error(w, "", http.StatusRequestTimeout)
				return
			}

		case Tuint16:
			U, ok := finduint16(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			b, err = set(U.n, U.w)
			if err != nil {
				http.Error(w, "", http.StatusRequestTimeout)
				return
			}

		case Tuint32:
			U, ok := finduint32(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			b, err = set(U.n, U.w)
			if err != nil {
				http.Error(w, "", http.StatusRequestTimeout)
				return
			}

		case Tuint64:
			U, ok := finduint64(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			b, err = set(U.n, U.w)
			if err != nil {
				http.Error(w, "", http.StatusRequestTimeout)
				return
			}

		case Tbyte:
			B, ok := findbyte(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			b, err = set(B.n, B.w)
			if err != nil {
				http.Error(w, "", http.StatusRequestTimeout)
				return
			}

		case Tfloat32:
			F, ok := findfloat32(name, idx)
			if !ok {
				delayedError(w, http.StatusNotFound)
				return
			}
			b, err = set(F.n, F.w)
			if err != nil {
				http.Error(w, "", http.StatusRequestTimeout)
				return
			}

		default:
			delayedError(w, http.StatusBadRequest)
			return
		}

		w.Write(b)
	})

	log.Fatal(http.ListenAndServe(Addr, nil))
}

func get(r chan []byte, b []byte) error {
	timeout := time.NewTimer(Timeout)
	defer timeout.Stop()
	select {
	case <-timeout.C:
		return errors.New("timeout; retry")
	case r <- b:
		return nil
	}
}

func set(n chan int, w chan []byte) ([]byte, error) {
	timeout := time.NewTimer(Timeout)
	defer timeout.Stop()
	select {
	case <-timeout.C:
		return nil, errors.New("timeout; retry")
	case n <- 1:
		select {
		case <-timeout.C:
			return nil, errors.New("timeout; retry")
		case b := <-w:
			return b, nil
		}
	}
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
	time.Sleep(Timeout)
	http.Error(w, "", code)
}
