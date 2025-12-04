package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("./advent3.txt");
	if err != nil {
		fmt.Println(err);
	}

	defer f.Close()
	buff := bufio.NewReader(f)
	joltages := make([]int, 0, 10)
	for {
		b, err := buff.ReadBytes('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				fmt.Println("End of file reached")
			} else {
				fmt.Println(err)
			}
			break
		}

		joltage := 0
		maxIndex := 0
		for i := 1; i >= 0; i-- {
			n, _ := strconv.Atoi(string(b[maxIndex]))
			for j := maxIndex+1; j < len(b)-i-1; j++ {
				c, _ := strconv.Atoi(string(b[j]))
				if c > n {
					n = c
					maxIndex = j
				}
			}
			joltage += n * int(math.Pow10(i))
			maxIndex++
		}

		joltages = append(joltages, joltage)
	}

	total := 0
	for _, v := range joltages {
		total += v
	}

	fmt.Println(total)
}
