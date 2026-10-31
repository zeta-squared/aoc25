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
	f, _ := os.ReadFile("./advent9.txt")
	f = f[:len(f)-1]
	s := strings.Split(string(f), "\n")
	closure := gridClosure(s)
	area := 0

	for _, u := range s {
		for _, v := range s {
			if v == u {
				continue
			}
			if rectangleInGrid(u, v, closure) {
				area = int(math.Max(rectangleArea(u, v), float64(area)))
			}
		}
	}

	fmt.Println(area)
}

func rectangleArea(a, b string) float64 {
	s := strings.Split(a, ",")
	x1, _ := strconv.Atoi(s[0])
	y1, _ := strconv.Atoi(s[1])

	s = strings.Split(b, ",")
	x2, _ := strconv.Atoi(s[0])
	y2, _ := strconv.Atoi(s[1])

	l := int(math.Abs(float64(x2-x1))) + 1
	w := int(math.Abs(float64(y2-y1))) + 1

	return float64(l * w)
}

func rectangleInGrid(a, b string, closure map[int][][]int) bool {
	s := strings.Split(a, ",")
	c1, _ := strconv.Atoi(s[0])
	r1, _ := strconv.Atoi(s[1])

	s = strings.Split(b, ",")
	c2, _ := strconv.Atoi(s[0])
	r2, _ := strconv.Atoi(s[1])

	colMin := int(math.Min(float64(c1), float64(c2)))
	colMax := int(math.Max(float64(c1), float64(c2)))
	rowMin := int(math.Min(float64(r1), float64(r2)))
	rowMax := int(math.Max(float64(r1), float64(r2)))

	contained := 0
	for i := rowMin; i <= rowMax; i++ {
		for j := range closure[i] {
			if closure[i][j][0] <= colMin && colMax <= closure[i][j][1] {
				contained++
				break
			}
		}
	}

	return contained == rowMax-rowMin+1
}

func gridHBoundaries(s []string) map[int]map[int]bool {
	hBoundary := map[int]map[int]bool{}

	prev := s[0]
	s = append(s[1:], prev)
	for len(s) > 0 {
		next := s[0]
		s = s[1:]

		prevSplit := strings.Split(prev, ",")
		nextSplit := strings.Split(next, ",")

		colPrev, _ := strconv.Atoi(prevSplit[0])
		rowPrev, _ := strconv.Atoi(prevSplit[1])
		colNext, _ := strconv.Atoi(nextSplit[0])
		rowNext, _ := strconv.Atoi(nextSplit[1])

		if colPrev == colNext {
			// Move vertically
			rowMin := int(math.Min(float64(rowPrev), float64(rowNext)))
			rowMax := int(math.Max(float64(rowPrev), float64(rowNext)))
			for i := rowMin; i <= rowMax; i++ {
				if _, ok := hBoundary[i]; !ok {
					hBoundary[i] = map[int]bool{colPrev: true}
					continue
				}

				hBoundary[i][colPrev] = true
			}
		}

		prev = next
	}

	return hBoundary
}

func gridClosure(s []string) map[int][][]int {
	hBoundary := gridHBoundaries(s)
	closure := map[int][][]int{}
	for k, v := range hBoundary {
		bounds := make([]int, 0, 10)
		for l := range v {
			bounds = append(bounds, l)
		}
		slices.Sort(bounds)

		if len(bounds)%2 == 0 {
			for len(bounds) > 0 {
				closure[k] = append(closure[k], []int{bounds[0], bounds[1]})
				bounds = bounds[2:]
			}
		} else {
			closure[k] = append(closure[k], []int{bounds[0], bounds[len(bounds)-1]})
		}
	}

	return closure
}
