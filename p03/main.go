package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.ReadFile("./advent2.txt");
	if err != nil {
		fmt.Println(err);
	}

	rec := make([]int, 0, 10)
	ranges := strings.SplitSeq(string(f[:len(f)-1]), ",")
	for v := range ranges {
		split := strings.Split(v, "-")
		m, _ := strconv.Atoi(split[0])
		M, _ := strconv.Atoi(split[1])
		for i := m; i <= M; i++ {
			strNum := strconv.Itoa(i)
			if len(strNum) % 2 != 0 {
				continue
			}
			n1, n2 := strNum[:len(strNum)/2], strNum[len(strNum)/2:]
			if compareStrNum(n1, n2) {
				rec = append(rec, i)
			}
		}
	}

	sum := 0
	for _, v := range rec {
		sum += v
	}
	
	fmt.Println(sum)
}

func compareStrNum(n1, n2 string) bool {
	for i := range len(n1) {
		if n1[i] != n2[i] {
			return false
		}
	}

	return true
}
