package sock

import "time"

const (
	DefaultClientAddr = "/"
	DefaultServerAddr = ":80"
	SOCK              = "9b466094ec991a03cb95c489c19c4d75635f0ae5"
	V                 = "▼"
	Timeout           = 30 * time.Second

	Terror byte = iota
	Tbool
	Tstring
	Tint
	Tint8
	Tint16
	Tint32
	Tint64
	Tuint
	Tuint8
	Tuint16
	Tuint32
	Tuint64
	Tbyte
	Tbytes
	Trune
	Tfloat32
	Tfloat64
)
