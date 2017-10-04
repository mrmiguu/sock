package main

import "github.com/mrmiguu/sock"

func main() {
	sckerr := sock.MakeError("err")
	err := <-sckerr
	println(err.Error())
	close(sckerr)
	select {}
}
