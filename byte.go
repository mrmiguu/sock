package sock

func Byte(key string) (chan<- byte, <-chan byte) {
	start.Do(run)

	w, r := make(chan byte), make(chan byte)

	bytel.Lock()
	B := wrbyte{
		key: key,
		idx: len(bytem[key]),
		r:   r,
	}
	bytem[key] = append(bytem[key], B)
	bytel.Unlock()

	go func() {
		for b := range w {
			write(Tbyte, B.key, B.idx, []byte{b})
		}
	}()

	return w, r
}

type wrbyte struct {
	key string
	idx int
	r   chan byte
}

func findByte(key string, idx int) *wrbyte {
	bytel.RLock()
	defer bytel.RUnlock()
	v, found := bytem[key]
	if !found {
		return nil
	}
	if idx >= len(v) || idx < 0 {
		return nil
	}
	return &v[idx]
}
