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

		password += dis / 100
		dis %= 100
		switch dir {
		case 'L':
			if pos != 0 && dis >= pos {
				password++
			}
			pos = (pos-dis+100) % 100
		default:
			if pos != 0 && dis >= 100-pos {
				password++
			}
			pos = (pos+dis) % 100
		}

	}

	fmt.Println(password)
}
