package main

import (
	"github.com/mrmiguu/sock"
)

func main() {
	Name := sock.Wstring()

	name := "YourNameHere"
	Found := sock.Rbool(name)

	Name <- name

	println(<-Found)

	select {}
}
