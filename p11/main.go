package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	b, _ := os.ReadFile("./a.txt")
	f := strings.Split(string(b), "\n")
	ops := strings.Fields(string(f[len(f)-2]))
	fmt.Println(ops)
	f = f[:len(f)-2]
	results := make([]int, len(ops))
	for i := range len(f) {
		s := strings.Split(f[i], " ")
		for j := range len(s) {
			switch ops[j] {
			case "*":
				n, _ := strconv.Atoi(s[j])
				results[j] += n
			default:
				n, _ := strconv.Atoi(string(s[j]))
				results[j] *= n
			}
		}
	}

	fmt.Println(results)
}
