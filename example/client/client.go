package main

import (
	"encoding/json"
	"time"

	"github.com/mrmiguu/sock"
)

type x struct {
	Home string
	V    struct {
		Z float64
	}
}

func main() {
	a, _ := sock.MakeBytes("a")
	w, r := sock.MakeBool("a")

	X := x{Home: "This is Home"}
	X.V.Z = 420.69
	b, _ := json.Marshal(X)

	a <- b

	start := time.Now()
	for range [2]int{} {
		w <- true
	}
	println(int(float64(time.Since(start).Nanoseconds())/2000000), "ms (w <-)")

	start = time.Now()
	for range [2]int{} {
		<-r
	}
	println(int(float64(time.Since(start).Nanoseconds())/2000000), "ms (<-r)")

	select {}
}
