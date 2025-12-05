package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("./advent5.txt");
	defer f.Close()
	buff := bufio.NewReader(f)
	ranges := make([][]byte, 0, 10)
	onAvailable := false
	count := 0
	for {
		b, err := buff.ReadBytes('\n')
		if err != nil {
			fmt.Println(err)
			break
		}

		if !onAvailable {
			if b[0] == '\n' {
				onAvailable = true
				continue
			} else {
				ranges = append(ranges, b[:len(b)-1])
			}
		}

		if onAvailable {
			currId, _ := strconv.Atoi(string(b[:len(b)-1]))
			for i := range len(ranges) {
				s := strings.Split(string(ranges[i]), "-")
				m, _ := strconv.Atoi(s[0])
				M, _ := strconv.Atoi(s[1])
				if m <= currId && currId <= M {
					count++
					break
				}
			}
		}
	}

	fmt.Println(count)
}
