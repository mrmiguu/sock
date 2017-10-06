package main

import (
	"encoding/json"
	"fmt"
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
	_, a := sock.MakeBytes("a")
	w, r := sock.MakeBool("a")

	var X x
	json.Unmarshal(<-a, &X)
	fmt.Println(X)

	start := time.Now()
	for range [2]int{} {
		<-r
	}
	println(int(float64(time.Since(start).Nanoseconds())/2000000), "ms (<-r)")

	start = time.Now()
	for range [2]int{} {
		w <- true
	}
	println(int(float64(time.Since(start).Nanoseconds())/2000000), "ms (w <-)")

	select {}
}
