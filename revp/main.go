package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var cmpl = map[byte]byte{'G': 'C', 'C': 'G', 'T': 'A', 'A': 'T'}

func isReversePalindrome(a []byte) bool {
	for i, ch := range a {
		if a[len(a)-i-1] != cmpl[ch] {
			return false
		}
	}
	return true
}

func main() {
	r := bufio.NewReader(os.Stdin)
	s, err := r.ReadBytes('\n')
	if err != nil {
		log.Fatalf("failed to read: %s", err)
	}
	a := s[:len(s)-1]
	for i := 0; i < len(a)-4; i++ {
		for n := 4; n <= 8 && i+n <= len(a); n++ {
			if isReversePalindrome(a[i : i+n]) {
				fmt.Println(i+1, n)
			}
		}
	}
}
