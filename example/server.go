package main

import (
	"fmt"
	"time"

	"github.com/mrmiguu/sock"
)

func main() {
	_, begin := sock.MakeBool("start")
	println(`start`, <-begin)

	sTest, rTest := sock.MakeByte("test")

	time.Sleep(1 * time.Second)
	println(`3...`)
	time.Sleep(1 * time.Second)
	println(`.2..`)
	time.Sleep(1 * time.Second)
	println(`..1.`)
	time.Sleep(1 * time.Second)
	println(`...GO!`)

	start := time.Now()
	for i := range [128]int{} {
		sTest <- byte(i)
	}
	println(int(float64(time.Since(start).Nanoseconds())/1000000/128), "ms (send)")

	var bytes []byte

	start = time.Now()
	for range [128]int{} {
		bytes = append(bytes, <-rTest)
	}
	fmt.Println(bytes)
	println(int(float64(time.Since(start).Nanoseconds())/1000000/128), "ms (recv)")

	select {}
}
