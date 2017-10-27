package main

import "github.com/mrmiguu/sock"

func main() {
	Name := sock.Rstring()

	for name := range Name {
		SOCKName := "name=" + name
		Found := sock.Wbool(SOCKName)
		Found <- true
	}
}
