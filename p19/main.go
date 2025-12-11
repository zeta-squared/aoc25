package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.ReadFile("./a.txt")
	s := strings.Split(string(f), "\n")
	s = s[:len(s)-1]
	for i := range len(s) {
		u := strings.Split(s[i], " ")
		l := indicatorLights(u[0][1 : len(u[0])-1])
		u = u[1 : len(u)-1]
		buttons := make([][]int, len(u))
		for j := range u {
			buttons[j] = buttonToBinary(len(l), u[j][1:len(u[j])-1])
		}

		queues := createQueues(buttons)
		minCount := math.MaxInt
		for i := range queues {
			count := 1

		outer:
			for len(queues[i]) > 0 {
				q := queues[i][0]
				queues[i] = queues[i][1:]

				// A single cycle involves iterating over all buttons
				count++
				for j := range buttons {
					n := make([]int, len(buttons[j]))
					// Add all elements of the button to the queue element (mod 2)
					for k := range buttons[j] {
						n[k] = (q[k] + buttons[j][k]) % 2
					}

					// If this modified state is the target state
					if slices.Equal(l, n) {
						break outer
					}

					// Push new element to the queue
					queues[i] = append(queues[i], n)
				}
			}

			minCount = int(math.Min(float64(count), float64(minCount)))
		}

		fmt.Println(minCount)
	}
}

func indicatorLights(light string) []int {
	lights := make([]int, len(light))
	for i := range light {
		switch light[i] {
		case '.':
			lights[i] = 0
		case '#':
			lights[i] = 1
		default:
		}
	}

	return lights
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

func createQueues(buttons [][]int) [][][]int {
	queues := make([][][]int, len(buttons))
	for i := range buttons {
		queues[i] = make([][]int, 1, 100)
		queues[i][0] = buttons[i]
	}

	return queues
}
