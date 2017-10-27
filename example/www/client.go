package main

import (
	"time"

	"github.com/mrmiguu/sock"
)

func main() {

	Name := sock.Wstring()
	name := "TheHowdyBuy"
	Name <- name

	time.Sleep(1 * time.Second)

	SOCKName := "name=" + name
	Found := sock.Rbool(SOCKName)
	println(<-Found)

	select {}
}
