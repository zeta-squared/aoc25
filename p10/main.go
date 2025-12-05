package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("./advent5.txt");
	defer f.Close()
	buff := bufio.NewReader(f)
	ranges := make([][]int, 0, 10)
	count := 0
	for {
		b, _ := buff.ReadBytes('\n')
		if b[0] == '\n' {
			break
		} else {
			r := string(b[:len(b)-1])
			s := strings.Split(r, "-")
			m, _ := strconv.Atoi(s[0])
			M, _ := strconv.Atoi(s[1])
			ranges = append(ranges, []int{m, M})
			count += M-m+1
			stack := make([][]int, 0, 3)
			stack = append(stack, []int{m, M})
			for {
				if len(stack) == 0 {
					break
				}
				
				curr := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				for i := 0; i < len(ranges)-1; i++ {
					if ranges[i][0] <= curr[0] && curr[0] <= ranges[i][1] && ranges[i][1] < curr[1] {
						count -= ranges[i][1]-curr[0]+1
						stack = append(stack, []int{ranges[i][1]+1, curr[1]})
						break
					} else if curr[0] < ranges[i][0] && ranges[i][0] <= curr[1] && curr[1] <= ranges[i][1] {
						count -= curr[1]-ranges[i][0]+1
						stack = append(stack, []int{curr[0], ranges[i][0]-1})
						break
					} else if ranges[i][0] <= curr[0] && curr[1] <= ranges[i][1] {
						count -= curr[1]-curr[0]+1
						break
					} else if curr[0] < ranges[i][0] && ranges[i][1] < curr[1] {
						count -= ranges[i][1]-ranges[i][0]+1
						stack = append(stack, []int{curr[0], ranges[i][0]-1}, []int{ranges[i][1]+1, curr[1]})
						break
					}
				}
			}
		}
	}

	fmt.Println(count)
}
