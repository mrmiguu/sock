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
	w, r := sock.MakeBytes("")

	var X x
	json.Unmarshal(<-r, &X)
	fmt.Println(X)

	start := time.Now()
	for range [100]int{} {
		<-r
	}
	println(int(float64(time.Since(start).Nanoseconds())/100000000), "ms (<-r)")

	start = time.Now()
	for range [100]int{} {
		w <- nil
	}
	println(int(float64(time.Since(start).Nanoseconds())/100000000), "ms (w <-)")

	select {}
}
