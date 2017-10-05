package main

import (
	"time"

	"github.com/mrmiguu/sock"
)

func main() {
	w, r := sock.MakeError("")

	w <- nil

	start := time.Now()
	for range [100]int{} {
		w <- nil
	}
	println(int(float64(time.Since(start).Nanoseconds())/100000000), "ms (w <- nil)")

	start = time.Now()
	for range [100]int{} {
		<-r
	}
	println(int(float64(time.Since(start).Nanoseconds())/100000000), "ms (<-r)")

	select {}
}
