package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	f, _ := os.ReadFile("./a.txt")
	s := strings.Split(string(f), "\n")
	devices := map[string][]string{}
	for i := range len(s)-1 {
		machines := strings.Split(s[i], " ")
		devices[machines[0][:3]] = machines[1:]
	}

	stack := make([]string, 0, 100)
	stack = append(stack, devices["you"]...)
	paths := 0
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if len(devices[top]) == 1 && devices[top][0] == "out" {
			paths++
			continue
		}

		stack = append(stack, devices[top]...)
	}

	fmt.Println(paths)
}
