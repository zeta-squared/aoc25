package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("./advent1.txt");
	if err != nil {
		fmt.Println(err);
	}

	defer f.Close()
	buff := bufio.NewReader(f)
	pos := 50
	password := 0
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

		dir := b[0]
		dis, err := strconv.Atoi(string(b[1:len(b)-1]))
		if err != nil {
			fmt.Println(err)
			break
		}

		switch dir {
		case 'L':
			pos = (pos-dis+100) % 100
		default:
			pos = (pos+dis) % 100
		}

		if pos == 0 {
			password++
		}
	}

	fmt.Println(password)
}
