package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.ReadFile("./advent2.txt");
	if err != nil {
		fmt.Println(err);
	}

	rec := make([]int, 0, 10)
	ranges := strings.SplitSeq(string(f[:len(f)-1]), ",")
	for v := range ranges {
		split := strings.Split(v, "-")
		m, _ := strconv.Atoi(split[0])
		M, _ := strconv.Atoi(split[1])
		for i := m; i <= M; i++ {
			strNum := strconv.Itoa(i)
			for j := 2; j <= len(strNum); j++ {
				if len(strNum) % j == 0 {
					pivot := len(strNum)/j
					nums := make([]string, 0, j)
					for k := 1; k <= j; k++ {
						nums = append(nums, strNum[(k-1)*pivot:k*pivot])
					}
					if compareStrNum(nums) {
						rec = append(rec, i)
						break
					}
				}
			}
		}
	}

	sum := 0
	for _, v := range rec {
		sum += v
	}
	
	fmt.Println(sum)
}

func compareStrNum(nums []string) bool {
	n0 := nums[0]
	for i := 1; i < len(nums); i++ {
		for j := range len(n0) {
			if n0[j] != nums[i][j] {
				return false
			}
		}
	}

	return true
}
