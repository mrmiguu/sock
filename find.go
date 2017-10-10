package sock

func findBool(key string, idx int) *wrbool {
	booll.RLock()
	defer booll.RUnlock()
	v, found := boolm[key]
	if !found {
		return nil
	}
	if idx >= len(v) || idx < 0 {
		return nil
	}
	return &v[idx]
}

func findString(key string, idx int) *wrstring {
	stringl.RLock()
	defer stringl.RUnlock()
	v, found := stringm[key]
	if !found {
		return nil
	}
	if idx >= len(v) || idx < 0 {
		return nil
	}
	return &v[idx]
}

func findInt(key string, idx int) *wrint {
	intl.RLock()
	defer intl.RUnlock()
	v, found := intm[key]
	if !found {
		return nil
	}
	if idx >= len(v) || idx < 0 {
		return nil
	}
	return &v[idx]
}

func findInt8(key string, idx int) *wrint8 {
	int8l.RLock()
	defer int8l.RUnlock()
	v, found := int8m[key]
	if !found {
		return nil
	}
	if idx >= len(v) || idx < 0 {
		return nil
	}
	return &v[idx]
}

func findInt16(key string, idx int) *wrint16 {
	int16l.RLock()
	defer int16l.RUnlock()
	v, found := int16m[key]
	if !found {
		return nil
	}
	if idx >= len(v) || idx < 0 {
		return nil
	}
	return &v[idx]
}

func findInt32(key string, idx int) *wrint32 {
	int32l.RLock()
	defer int32l.RUnlock()
	v, found := int32m[key]
	if !found {
		return nil
	}
	if idx >= len(v) || idx < 0 {
		return nil
	}
	return &v[idx]
}

func findInt64(key string, idx int) *wrint64 {
	int64l.RLock()
	defer int64l.RUnlock()
	v, found := int64m[key]
	if !found {
		return nil
	}
	if idx >= len(v) || idx < 0 {
		return nil
	}
	return &v[idx]
}

func findUint(key string, idx int) *wruint {
	uintl.RLock()
	defer uintl.RUnlock()
	v, found := uintm[key]
	if !found {
		return nil
	}
	if idx >= len(v) || idx < 0 {
		return nil
	}
	return &v[idx]
}

func findUint8(key string, idx int) *wruint8 {
	uint8l.RLock()
	defer uint8l.RUnlock()
	v, found := uint8m[key]
	if !found {
		return nil
	}
	if idx >= len(v) || idx < 0 {
		return nil
	}
	return &v[idx]
}

func findUint16(key string, idx int) *wruint16 {
	uint16l.RLock()
	defer uint16l.RUnlock()
	v, found := uint16m[key]
	if !found {
		return nil
	}
	if idx >= len(v) || idx < 0 {
		return nil
	}
	return &v[idx]
}

func findUint32(key string, idx int) *wruint32 {
	uint32l.RLock()
	defer uint32l.RUnlock()
	v, found := uint32m[key]
	if !found {
		return nil
	}
	if idx >= len(v) || idx < 0 {
		return nil
	}
	return &v[idx]
}

func findUint64(key string, idx int) *wruint64 {
	uint64l.RLock()
	defer uint64l.RUnlock()
	v, found := uint64m[key]
	if !found {
		return nil
	}
	if idx >= len(v) || idx < 0 {
		return nil
	}
	return &v[idx]
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

func findBytes(key string, idx int) *wrbytes {
	bytesl.RLock()
	defer bytesl.RUnlock()
	v, found := bytesm[key]
	if !found {
		return nil
	}
	if idx >= len(v) || idx < 0 {
		return nil
	}
	return &v[idx]
}

func findRune(key string, idx int) *wrrune {
	runel.RLock()
	defer runel.RUnlock()
	v, found := runem[key]
	if !found {
		return nil
	}
	if idx >= len(v) || idx < 0 {
		return nil
	}
	return &v[idx]
}

func findFloat32(key string, idx int) *wrfloat32 {
	float32l.RLock()
	defer float32l.RUnlock()
	v, found := float32m[key]
	if !found {
		return nil
	}
	if idx >= len(v) || idx < 0 {
		return nil
	}
	return &v[idx]
}

func findFloat64(key string, idx int) *wrfloat64 {
	float64l.RLock()
	defer float64l.RUnlock()
	v, found := float64m[key]
	if !found {
		return nil
	}
	if idx >= len(v) || idx < 0 {
		return nil
	}
	return &v[idx]
}

func findError(key string, idx int) *wrerror {
	errorl.RLock()
	defer errorl.RUnlock()
	v, found := errorm[key]
	if !found {
		return nil
	}
	if idx >= len(v) || idx < 0 {
		return nil
	}
	return &v[idx]
}
