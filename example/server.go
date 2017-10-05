package main

import (
	"fmt"

	"github.com/mrmiguu/sock"
)

func main() {
	sckerr := sock.MakeError("err")
	fmt.Println(<-sckerr)
	select {}
}
