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
	f, _ := os.ReadFile("./advent10.txt")
	s := strings.Split(string(f), "\n")
	s = s[:len(s)-1]
	total := 0
	for i := range len(s) {
		u := strings.Split(s[i], " ")
		l := indicatorLights(u[0][1 : len(u[0])-1])
		u = u[1 : len(u)-1]
		buttons := make([][]int, len(u))
		for j := range u {
			buttons[j] = buttonToBinary(len(l), u[j][1:len(u[j])-1])
		}

		minCount := math.MaxInt
		for i := range buttons {
			count := 1
			queue := make([][]int, 0, 100)
			if slices.Equal(l, buttons[i]) {
				minCount = count
				break
			}
			q := [][]int{buttons[i]}
		outer:
			for count <= minCount-1 {
				// A single cycle involves iterating over all buttons
				count++
				for j := range q {
					for k := range buttons {
						n := make([]int, len(buttons[k]))
						// Add all elements of the button to the queue element (mod 2)
						for h := range buttons[k] {
							n[h] = (q[j][h] + buttons[k][h]) % 2
						}

						// If this modified state is the target state
						if slices.Equal(l, n) {
							break outer
						}

						// Push new element to the queue
						queue = append(queue, n)
					}
				}

				if len(queue) == 0 {
					break
				}

				nextLevel := len(q)*len(buttons)
				q = queue[:nextLevel]
				queue = queue[nextLevel:]
			}

			minCount = int(math.Min(float64(count), float64(minCount)))
		}

		total += minCount
	}

	fmt.Println(total)
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
