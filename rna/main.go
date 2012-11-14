package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

func main() {
	var (
		n   int
		err error
	)
	buf := make([]byte, 1024)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	for err == nil {
		n, err = os.Stdin.Read(buf)
		for _, ch := range buf[:n] {
			switch ch {
			case 'A', 'C', 'G':
				out.WriteByte(ch)
			case 'T':
				out.WriteByte('U')
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
}
