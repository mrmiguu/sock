package main

import (
	"github.com/mrmiguu/sock"
)

func main() {
	sckerr := sock.MakeError("err")
	sckerr <- nil
	select {}
}
