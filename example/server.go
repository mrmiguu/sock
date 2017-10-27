package main

import "github.com/mrmiguu/sock"

func main() {
	Name := sock.Rstring()

	for name := range Name {
		Found := sock.Wbool(name)
		Found <- true
		sock.Close(name)
	}
}
