package main

import (
	"fmt"
	"os"
	"strings"
)

type Path struct {
	current string
	visited map[string]bool
}

func main() {
	f, _ := os.ReadFile("./b.txt")
	s := strings.Split(string(f), "\n")
	devices := map[string][]string{}
	for i := range len(s)-1 {
		machines := strings.Split(s[i], " ")
		devices[machines[0][:3]] = machines[1:]
	}

	stack := make([]Path, 0, 100)
	for i := range devices["svr"] {
		stack = append(stack, Path{current: devices["svr"][i], visited: map[string]bool{devices["svr"][i]: true}})
	}
	paths := 0
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if len(devices[top.current]) == 1 && devices[top.current][0] == "out" {
			if top.visited["dac"] && top.visited["fft"] {
				paths++
				fmt.Println(paths)
			}
			continue
		}

		for i := range devices[top.current] {
			if top.visited[devices[top.current][i]] {
				continue
			}

			m := make(map[string]bool, len(top.visited)+1)
			for k, v := range top.visited {
				m[k] = v
			}
			m[devices[top.current][i]] = true
			stack = append(stack, Path{current: devices[top.current][i], visited: m})
		}
	}

	fmt.Println(paths)
}
