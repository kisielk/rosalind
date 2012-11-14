package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	var (
		err       error
		seq       []byte
		name      []byte
		bestName  []byte
		bestRatio float64
	)

	in := bufio.NewReader(os.Stdin)
	for {
		var total, gc int = 0, 0
		if name, err = in.ReadBytes('\n'); err != nil {
			break
		}
		seq, err = in.ReadBytes('>')
		for _, ch := range seq {
			switch ch {
			case 'G', 'C':
				gc++
				total++
			case 'T', 'A':
				total++
			case ' ', '\n':
				continue
			case '>':
				break
			default:
				log.Fatalf("Unknown character: %s", ch)
			}
		}
		ratio := float64(gc) / float64(total)
		if ratio > bestRatio {
			bestName = name[:len(name)-1]
			bestRatio = ratio
		}
		if err != nil {
			break
		}
	}
	if err != io.EOF {
		log.Fatalf("Encountered an error: %s", err)
	}
	fmt.Printf("%s\n", bestName)
	fmt.Printf("%f%%\n", bestRatio*100)
}
