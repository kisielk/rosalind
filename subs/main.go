package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var (
		sep, in   string
		sepb, inb []byte
		result    []int
		err       error
	)
	r := bufio.NewReader(os.Stdin)
	if inb, err = r.ReadBytes('\n'); err != nil {
		log.Fatalf("Couldn't read input: %s", err)
	} else {
		in = string(inb[:len(inb)-1])
	}
	if sepb, err = r.ReadBytes('\n'); err != nil {
		log.Fatalf("Couldn't read substring: %s", err)
	} else {
		sep = string(sepb[:len(sepb)-1])
	}
	n := len(sep)
	if n == 0 {
		fmt.Println(result)
		return
	}
	c := sep[0]
	for i := 0; i+n <= len(in); i++ {
		if in[i] == c && in[i:i+n] == sep {
			result = append(result, i+1)
		}
	}
	fmt.Print(result)
}
