package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	f, _ := os.Open("./advent7.txt")
	buff := bufio.NewReader(f)
	b, _ := buff.ReadBytes('\n')
	m := map[int]bool{}
	splits := 0
	for i := range len(b) {
		switch b[i] {
		case 'S':
			m[i] = true
		default:
			m[i] = false
		}
	}

	for {
		b, err := buff.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
		}

		for i := range len(b[:len(b)-1]) {
			if b[i] == '^' {
				if m[i] {
					switch i {
					case len(b)-1:
						m[i-1] = true
					case 0:
						m[i+1] = true
					default:
						m[i-1] = true
						m[i+1] = true
					}
					m[i] = false
					splits++
				}
			}
		}
	}

	fmt.Println(splits)
}
