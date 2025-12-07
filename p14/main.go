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
	m := map[int]int{}
	for i := range len(b) {
		switch b[i] {
		case 'S':
			m[i] = 1
		default:
			m[i] = 0
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
				if m[i] > 0 {
					switch i {
					case len(b)-1:
						m[i-1] += m[i]
					case 0:
						m[i+1] += m[i]
					default:
						m[i-1] += m[i]
						m[i+1] += m[i]
					}
					m[i] = 0
				}
			}
		}
	}

	journeys := 0
	for _, v := range m {
		journeys += v
	}
	fmt.Println(journeys)
}
