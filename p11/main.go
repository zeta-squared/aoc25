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
	results := make([]int, len(ops))
	for i, v := range strings.Fields(f[0]) {
		n, _ := strconv.Atoi(v)
		results[i] = n
	}

	f = f[1:len(f)-2]
	for i := range len(f) {
		s := strings.Fields(f[i])
		for j := range len(s) {
			n, _ := strconv.Atoi(s[j])
			switch ops[j] {
			case "*":
				results[j] *= n
			default:
				results[j] += n
			}
		}
	}

	sum := 0
	for _, res := range results {
		sum += res
	}

	fmt.Println(sum)
}
