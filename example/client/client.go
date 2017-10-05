package main

import (
	"time"

	load "github.com/mrmiguu/Loading"
	"github.com/mrmiguu/sock"
)

func main() {
	c := sock.MakeError("")

	done := load.New("starting")
	c <- nil
	done <- true

	done = load.New("c <- nil")
	start := time.Now()
	for i := range [100]int{} {
		c <- nil
		println(i + 1)
	}
	done <- true
	println(int(float64(time.Since(start).Nanoseconds())/100000000), "ms (c <- nil)")

	done = load.New("<-c")
	start = time.Now()
	for range [100]int{} {
		<-c
	}
	done <- true
	println(int(float64(time.Since(start).Nanoseconds())/100000000), "ms (<-c)")

	select {}
}
