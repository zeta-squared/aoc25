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
	maxPairs := 2*1000
	for i, v := range pairs {
		if i == maxPairs {
			break
		}
		if v.junctionBoxA.circuit == nil && v.junctionBoxB.circuit == nil {
			c := &Circuit{count: 2}
			circuits = append(circuits, c)
			v.junctionBoxA.circuit = c
			v.junctionBoxB.circuit = c
		} else if v.junctionBoxA.circuit != nil && v.junctionBoxB.circuit == nil {
			v.junctionBoxB.circuit = v.junctionBoxA.circuit
			v.junctionBoxA.circuit.count++
		} else if v.junctionBoxA.circuit == nil && v.junctionBoxB.circuit != nil {
			v.junctionBoxA.circuit = v.junctionBoxB.circuit
			v.junctionBoxA.circuit.count++
		} else if v.junctionBoxA.circuit != nil && v.junctionBoxB.circuit != nil {
			if v.junctionBoxA.circuit != v.junctionBoxB.circuit {
				ca := v.junctionBoxA.circuit
				cb := v.junctionBoxB.circuit

				ca.count += cb.count
				cb.count = 0

				for _, jb := range junctionBoxes {
					if jb.circuit == cb {
						jb.circuit = ca
					}
				}
			}
		}
	}

	slices.SortFunc(circuits, func (a, b *Circuit) int {
		return b.count-a.count
	})

	prod := 1
	for i := range 3 {
		prod *= circuits[i].count
	}

	fmt.Println(prod)
}
