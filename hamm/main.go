package main

import (
	"flag"
	"fmt"
	"log"
)

var s = flag.String("s", "", "string s")
var t = flag.String("t", "", "string t")

func main() {
	flag.Parse()
	bs := []byte(*s)
	bt := []byte(*t)
	if len(bs) != len(bt) {
		log.Fatal("s and t must be the same length")
	}

	dist := 0
	for i := 0; i < len(bs); i++ {
		if bs[i] != bt[i] {
			dist++
		}
	}
	fmt.Println(dist)
}
