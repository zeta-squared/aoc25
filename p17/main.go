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
	area := 0
	for _, u := range s {
		for _, v := range s {
			area = int(math.Max(rectangleArea(u, v), float64(area)))
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
