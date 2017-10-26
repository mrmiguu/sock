package main

import (
	"github.com/mrmiguu/jsutil"
	"github.com/mrmiguu/sock"
)

func main() {
	F := sock.Wstring()
	R := sock.Rstring()
	E := sock.Rerror()

	for {
		F <- jsutil.Prompt("ENTER FUNCTION:")
		select {
		case res := <-R:
			jsutil.Alert(res)
		case err := <-E:
			jsutil.Alert("ERROR: " + err.Error())
		}
	}
}
