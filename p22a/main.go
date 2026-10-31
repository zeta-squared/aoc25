package main

import (
	"fmt"
	"os"
	"slices"
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
	for i := range len(s) - 1 {
		machines := strings.Split(s[i], " ")
		devices[machines[0][:3]] = machines[1:]
	}

	dacFirstPath := 1
	dacFirstPath *= paths(devices, "svr", "dac")
	if dacFirstPath > 0 {
		dacFirstPath *= paths(devices, "dac", "fft")
	}

	if dacFirstPath > 0 {
		dacFirstPath *= paths(devices, "fft", "out")
	}

	fftFirstPath := 1
	fftFirstPath *= paths(devices, "svr", "fft")
	if fftFirstPath > 0 {
		fftFirstPath *= paths(devices, "fft", "dac")
	}

	if fftFirstPath > 0 {
		fftFirstPath *= paths(devices, "dac", "out")
	}

	fmt.Println(dacFirstPath + fftFirstPath)
}

func paths(devices map[string][]string, from string, to string) int {
	stack := make([]Path, 0, 100)
	for i := range devices[from] {
		stack = append(stack, Path{current: devices[from][i], visited: map[string]bool{devices[from][i]: true}})
	}

	paths := 0
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		_, ok := slices.BinarySearch(devices[top.current], to)
		if ok {
			paths++
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

	return paths
}
