package main

import (
	"time"

	"github.com/mrmiguu/sock"
)

func main() {
	_, begin := sock.Byte("start")

	println(`<-begin...`)
	<-begin
	println(`       !!!`)

	sTest, rTest := sock.Byte("test")

	time.Sleep(1 * time.Second)
	println(`3...`)
	time.Sleep(1 * time.Second)
	println(`.2..`)
	time.Sleep(1 * time.Second)
	println(`..1.`)
	time.Sleep(1 * time.Second)
	println(`...GO!`)

	start := time.Now()
	for i := range [256]int{} {
		sTest <- byte(i)
	}
	println(int(float64(time.Since(start).Nanoseconds())/1000000/256), "ms (send)")

	start = time.Now()
	for range [256]int{} {
		println(<-rTest)
	}
	println(int(float64(time.Since(start).Nanoseconds())/1000000/256), "ms (recv)")

	select {}
}
