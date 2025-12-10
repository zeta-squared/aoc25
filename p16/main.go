package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Circuit struct {
	count int
}

type JunctionBox struct {
	position string
	circuit *Circuit
}

type Pair struct {
	junctionBoxA *JunctionBox
	junctionBoxB *JunctionBox
	distance int
}

func (pair Pair) l2Norm() int {
	p1 := strings.Split(pair.junctionBoxA.position, ",")
	p2 := strings.Split(pair.junctionBoxB.position, ",")
	sum := 0
	for i := range len(p1) {
		x1, _ := strconv.Atoi(p1[i])
		x2, _ := strconv.Atoi(p2[i])
		sum += (x2-x1)*(x2-x1)
	}

	return sum
}

func (pair Pair) prod() int {
	p1 := strings.Split(pair.junctionBoxA.position, ",")
	p2 := strings.Split(pair.junctionBoxB.position, ",")
	x1, _ := strconv.Atoi(p1[0])
	x2, _ := strconv.Atoi(p2[0])

	return x1*x2
}

func main() {
	f, _ := os.ReadFile("./advent8.txt")
	f = f[:len(f)-1]
	s := strings.Split(string(f), "\n")
	junctionBoxes := map[string]*JunctionBox{}
	for _, v := range s {
		junctionBoxes[v] = &JunctionBox{position: v}
	}

	pairs := make([]Pair, 0, len(s)*(len(s)-1))
	for _, v := range s {
		for _, u := range s {
			if u == v { continue }
			
			p := Pair{junctionBoxA: junctionBoxes[v], junctionBoxB: junctionBoxes[u]}
			p.distance = p.l2Norm()
			pairs = append(pairs, p)
		}
	}

	slices.SortFunc(pairs, func (a, b Pair) int {
		return a.distance-b.distance
	})

	circuits := make([]*Circuit, 0, 1000)
	circuitCount := len(s)
	var last Pair
	for _, v := range pairs {
		if v.junctionBoxA.circuit == nil && v.junctionBoxB.circuit == nil {
			c := &Circuit{count: 2}
			circuits = append(circuits, c)
			circuitCount--
			v.junctionBoxA.circuit = c
			v.junctionBoxB.circuit = c
		} else if v.junctionBoxA.circuit != nil && v.junctionBoxB.circuit == nil {
			v.junctionBoxB.circuit = v.junctionBoxA.circuit
			v.junctionBoxA.circuit.count++
			circuitCount--
		} else if v.junctionBoxA.circuit == nil && v.junctionBoxB.circuit != nil {
			v.junctionBoxA.circuit = v.junctionBoxB.circuit
			v.junctionBoxA.circuit.count++
			circuitCount--
		} else if v.junctionBoxA.circuit != nil && v.junctionBoxB.circuit != nil {
			if v.junctionBoxA.circuit != v.junctionBoxB.circuit {
				ca := v.junctionBoxA.circuit
				cb := v.junctionBoxB.circuit

				ca.count += cb.count
				cb.count = 0
				circuitCount--

				for _, jb := range junctionBoxes {
					if jb.circuit == cb {
						jb.circuit = ca
					}
				}
			}
		}

		if circuitCount == 1 {
			last = v
			break
		}
	}

	fmt.Println(last.prod())
}
