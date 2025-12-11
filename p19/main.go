package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.ReadFile("./a.txt")
	s := strings.Split(string(f), "\n")
	s = s[:len(s)-1]
	for i := range 1 {
		u := strings.Split(s[i], " ")
		l := indicatorLights(u[0][1:len(u[0])-1])
		u = u[1:len(u)-1]
		buttons := make([][]int, len(u))
		for j := range u {
			buttons[j] = make([]int, len(l))
			buttonToBinary(buttons[j], u[j][1:len(u[j])-1])
		}

		for j := range buttons {
			
		}
	}
}

func indicatorLights(u string) []int {
	lights := make([]int, len(u))
	for i := range u {
		switch u[i] {
		case '.':
			lights[i] = 0
		case '#':
			lights[i] = 1
		default:
		}
	}

	return lights
}

func buttonToBinary(b []int, u string) {
	for i := range u {
		if u[i] == ',' {
			continue
		}
		j, _ := strconv.Atoi(string(u[i]))
		b[j] = 1
	}
}

func l1Norm(a, b []int) int {
	norm := 0

	return norm
}
