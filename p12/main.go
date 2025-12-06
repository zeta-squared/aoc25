package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	b, _ := os.ReadFile("./advent6.txt")
	f := strings.Split(string(b), "\n")
	ops := strings.Fields(string(f[len(f)-2]))
	f = f[:len(f)-2]
	// Transpose the text
	t := make([][]byte, len(f[0]))
	for i := range len(t) {
		t[i] = make([]byte, len(f))
	}

	for i := range len(f) {
		for j := range len(f[i]) {
			t[j][i] = f[i][j]
		}
	}

	results := make([]int, len(ops))
	for i := range len(ops) {
		switch ops[i] {
		case "*":
			results[i] = 1
		default:
		}
	}

	opIdx := 0
	for i := range len(t) {
		s := strings.Fields(string(t[i]))
		if len(s) == 0 {
			opIdx++
			continue
		}

		n, _ := strconv.Atoi(s[0])
		switch ops[opIdx] {
		case "*":
			results[opIdx] *= n
		default:
			results[opIdx] += n
		}
	}

	sum := 0
	for _, res := range results {
		sum += res
	}

	fmt.Println(sum)
}
