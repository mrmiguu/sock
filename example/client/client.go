package main

import (
	"errors"

	"github.com/mrmiguu/sock"
)

func main() {
	sckerr := sock.MakeError("err")
	sckerr <- errors.New("something went wrong")
	close(sckerr)
	select {}
}
