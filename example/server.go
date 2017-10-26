package main

import (
	"errors"
	"strconv"
	"strings"

	"github.com/mrmiguu/sock"
)

func main() {
	F := sock.Rstring()
	R := sock.Wstring()
	E := sock.Werror()

	for f := range F {
		args := strings.Split(f, "][")
		switch args[0] {
		case "add":
			a, _ := strconv.Atoi(args[1])
			b, _ := strconv.Atoi(args[2])
			R <- strconv.Itoa(a + b)
		default:
			E <- errors.New("bad function format")
		}
	}
}
