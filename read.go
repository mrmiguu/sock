package sock

import (
	"bytes"
	"errors"
)

func read(pkt []byte) error {
	parts := bytes.Split(pkt, v)
	if len(parts) != 4 || len(parts[0]) < 1 {
		return errors.New("invalid packet")
	}
	t, key, idx, body := parts[0][0], string(parts[1]), bytes2int(parts[2]), parts[3]

	switch t {
	case Tbyte:
		B := findByte(key, idx)
		if B != nil {
			B.r <- body[0]
		}
	}

	return nil
}
