package main

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"os"
)

func main() {
	var (
		n   int
		err error
		seq bytes.Buffer
	)
	buf := make([]byte, 1024)

	for {
		n, err = os.Stdin.Read(buf)
		for _, ch := range buf[:n] {
			switch ch {
			case 'A':
				seq.WriteByte('T')
			case 'C':
				seq.WriteByte('G')
			case 'G':
				seq.WriteByte('C')
			case 'T':
				seq.WriteByte('A')
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
		log.Fatalf("Encountered an error: %s", err)
	}

	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	b := seq.Bytes()
	for i := 0; i < len(b); i++ {
		out.WriteByte(b[len(b)-1-i])
	}
}
