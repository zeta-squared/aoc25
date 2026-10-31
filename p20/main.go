package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Frame struct {
	state []int
	key string
	depth int
}

func main() {
	f, _ := os.ReadFile("./advent10.txt")
	s := strings.Split(string(f), "\n")
	s = s[:len(s)-1]
	total := 0
	for i := range len(s) {
		fmt.Printf("%d of %d\n", i+1, len(s))
		u := strings.Split(s[i], " ")
		jReq := joltageRequirements(u[len(u)-1][1 : len(u[len(u)-1])-1])
		u = u[1 : len(u)-1]
		buttons := make([][]int, len(u))
		for j := range u {
			buttons[j] = buttonToBinary(len(jReq), u[j][1:len(u[j])-1])
		}

		minCount := math.MaxInt
		start := make([]int, len(jReq))
		startKey := stateKey(start)
		queue := []Frame{{state: start, key: startKey, depth: 0}}
		visited := map[string]bool{startKey: true}
		for len(queue) > 0 {
			top := queue[0]
			queue = queue[1:]

			if top.depth >= minCount {
				continue
			}

			if slices.Equal(jReq, top.state) {
				minCount = top.depth
				continue
			}

		buttonsLoop:
			for j := range buttons {
				next := make([]int, len(top.state))
				for k := range top.state {
					next[k] = top.state[k] + buttons[j][k]
					if next[k] > jReq[k] {
						continue buttonsLoop
					}
				}

				key := stateKey(next)
				if visited[key] {
					continue
				}

				visited[key] = true
				queue = append(queue, Frame{state: next, key: key, depth: top.depth + 1})
			}
		}

		total += minCount
	}

	fmt.Println(total)
}

func joltageRequirements(target string) []int {
	tSplit := strings.Split(target, ",")
	joltages := make([]int, len(tSplit))
	for i := range tSplit {
		j, _ := strconv.Atoi(tSplit[i])
		joltages[i] = j
	}

	return joltages
}

func buttonToBinary(length int, button string) []int {
	binaryButton := make([]int, length)
	for i := range button {
		if button[i] == ',' {
			continue
		}
		j, _ := strconv.Atoi(string(button[i]))
		binaryButton[j] = 1
	}

	return binaryButton
}

func stateKey(s []int) string {
	b := make([]byte, 0, len(s)*3)
	for i := range s {
		b = strconv.AppendInt(b, int64(s[i]), 10)
		b = append(b, ',')
	}

	return string(b)
}
