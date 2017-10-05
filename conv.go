package sock

import (
	"encoding/binary"
	"errors"
	"strconv"
)

func int2bytes(i int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(i))
	return b
}

func bytes2int(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}

func bool2bytes(b bool) []byte {
	return strconv.AppendBool(nil, b)
}

func bytes2bool(b []byte) bool {
	tf, _ := strconv.ParseBool(string(b))
	return tf
}

func error2bytes(e error) []byte {
	if e == nil {
		return []byte{}
	}
	return []byte(e.Error())
}

func bytes2error(b []byte) error {
	if len(b) == 0 {
		return nil
	}
	return errors.New(string(b))
}
