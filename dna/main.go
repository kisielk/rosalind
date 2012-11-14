package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	var (
		a, c, g, t, n int
		err           error
	)
	buf := make([]byte, 1024)

	for err == nil {
		n, err = os.Stdin.Read(buf)
		for _, ch := range buf[:n] {
			switch ch {
			case 'A':
				a++
			case 'C':
				c++
			case 'G':
				g++
			case 'T':
				t++
			case ' ', '\n':
				continue
			default:
				log.Fatalf("Unknown character: %s", ch)
			}
		}
		if err != nil {
			break
		}
	}
	if err != io.EOF {
		log.Fatal(err)
	}
	fmt.Println(a, c, g, t)
}
