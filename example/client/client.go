package main

import (
	"time"

	"github.com/mrmiguu/sock"
)

func main() {
	begin, _ := sock.MakeBool("start")

	println(`begin <-...`)
	begin <- true
	println(`        !!!`)

	sTest, rTest := sock.MakeBytes("test")

	time.Sleep(1 * time.Second)
	println(`3...`)
	time.Sleep(1 * time.Second)
	println(`.2..`)
	time.Sleep(1 * time.Second)
	println(`..1.`)
	time.Sleep(1 * time.Second)
	println(`...GO!`)

	start := time.Now()
	for range [100]int{} {
		<-rTest
	}
	println(int(float64(time.Since(start).Nanoseconds())/1000000/100), "ms (recv)")

	start = time.Now()
	for range [100]int{} {
		sTest <- nil
	}
	println(int(float64(time.Since(start).Nanoseconds())/1000000/100), "ms (send)")

	select {}
}
