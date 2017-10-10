package main

import (
	load "github.com/mrmiguu/Loading"
	"github.com/mrmiguu/sock/sock"
)

func main() {
	_, r := sock.Byte("")

	done := load.New("<-r")
	println(<-r)
	done <- true

	// _, begin := sock.MakeBool("start")

	// println(`<-begin...`)
	// <-begin
	// println(`       !!!`)

	// sTest, rTest := sock.MakeBytes("test")

	// time.Sleep(1 * time.Second)
	// println(`3...`)
	// time.Sleep(1 * time.Second)
	// println(`.2..`)
	// time.Sleep(1 * time.Second)
	// println(`..1.`)
	// time.Sleep(1 * time.Second)
	// println(`...GO!`)

	// start := time.Now()
	// for range [100]int{} {
	// 	sTest <- nil
	// }
	// println(int(float64(time.Since(start).Nanoseconds())/1000000/100), "ms (send)")

	// start = time.Now()
	// for range [100]int{} {
	// 	<-rTest
	// }
	// println(int(float64(time.Since(start).Nanoseconds())/1000000/100), "ms (recv)")

	select {}
}
