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
	w, r := sock.MakeBytes("")

	X := x{Home: "This is Home"}
	X.V.Z = 420.69
	b, _ := json.Marshal(X)

	w <- b

	start := time.Now()
	for range [100]int{} {
		w <- nil
	}
	println(int(float64(time.Since(start).Nanoseconds())/100000000), "ms (w <-)")

	start = time.Now()
	for range [100]int{} {
		<-r
	}
	println(int(float64(time.Since(start).Nanoseconds())/100000000), "ms (<-r)")

	select {}
}
