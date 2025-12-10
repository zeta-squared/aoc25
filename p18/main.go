package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.ReadFile("./advent9.txt")
	f = f[:len(f)-1]
	s := strings.Split(string(f), "\n")
	g := fillGrid(s)
	area := 0
	
	for _, u := range s {
		for _, v := range s {
			if v == u {
				continue
			}
			if rectangleInGrid(u, v, g) {
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

	l := int(math.Abs(float64(x2-x1+1)))
	w := int(math.Abs(float64(y2-y1+1)))

	return float64(l*w)
}

func fillGrid(s []string) [][]bool {
	xMax := 98411
	yMax := 98411
	g := make([][]bool, yMax)
	for i := range g {
		g[i] = make([]bool, xMax)
	}

	prev := s[0]
	s = append(s[1:], prev)
	for len(s) > 0 {
		next := s[0]
		s = s[1:]

		prevSplit := strings.Split(prev, ",")
		nextSplit := strings.Split(next, ",")

		xPrev, _ := strconv.Atoi(prevSplit[0])
		yPrev, _ := strconv.Atoi(prevSplit[1])
		xNext, _ := strconv.Atoi(nextSplit[0])
		yNext, _ := strconv.Atoi(nextSplit[1])

		if xPrev == xNext {
			// Move vertically
			yMin := int(math.Min(float64(yPrev), float64(yNext)))
			yMax := int(math.Max(float64(yPrev), float64(yNext)))
			for i := yMin; i <= yMax; i++ {
				g[i][xPrev] = true
			}
		} else {
			// Move horizontally
			xMin := int(math.Min(float64(xPrev), float64(xNext)))
			xMax := int(math.Max(float64(xPrev), float64(xNext)))
			for i := xMin; i <= xMax; i++ {
				g[yPrev][i] = true
			}
		}

		prev = next
	}

	for i := range g {
		l := -1
		r := -1
		for j := range g[i] {
			if l == -1 && g[i][j] {
				l = j
			}

			if r == -1 && g[i][len(g[i])-1-j] {
				r = len(g[i])-1-j
			}

			if l != -1 && r != -1 {
				break
			}
		}

		if l != -1 && r != -1 {
			for j := l; j <= r; j++ {
				g[i][j] = true
			}
		}
	}

	return g
}

func rectangleInGrid(a, b string, g [][]bool) bool {
	s := strings.Split(a, ",")
	x1, _ := strconv.Atoi(s[0])
	y1, _ := strconv.Atoi(s[1])

	s = strings.Split(b, ",")
	x2, _ := strconv.Atoi(s[0])
	y2, _ := strconv.Atoi(s[1])
	
	xMin := int(math.Min(float64(x1), float64(x2)))
	xMax := int(math.Max(float64(x1), float64(x2)))
	yMin := int(math.Min(float64(y1), float64(y2)))
	yMax := int(math.Max(float64(y1), float64(y2)))

	xBoundary := map[int][]int{
		x1: {yMin, yMax},
		x2: {yMin, yMax},
	}
	yBoundary := map[int][]int{
		y1: {xMin, xMax},
		y2: {xMin, xMax},
	}

	for k, v := range xBoundary {
		for i := v[0]; i <= v[1]; i++ {
			if !g[i][k] {
				return false
			}
		}
	}

	for k, v := range yBoundary {
		for i := v[0]; i <= v[1]; i++ {
			if !g[k][i] {
				return false
			}
		}
	}

	return true
}
