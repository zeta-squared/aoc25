package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	f, _ := os.Open("./advent4.txt");
	defer f.Close()
	buff := bufio.NewReader(f)
	paperMatrix := make([][]byte, 0, 10)
	for {
		b, err  := buff.ReadBytes('\n')
		if err != nil {
			break
		}
		paperMatrix = append(paperMatrix, b)
	}

	rollCount := 0
	for {
		iterCount := 0
		for i := range len(paperMatrix) {
			for j := range len(paperMatrix[i]) {
				if paperMatrix[i][j] != '@' {
					continue
				}

				colMin := int(math.Max(0, float64(j-1)))
				colMax := int(math.Min(float64(len(paperMatrix[i])-1), float64(j+1)))+1
				rowMin := int(math.Max(0, float64(i-1)))
				rowMax := int(math.Min(float64(i+1), float64(len(paperMatrix)-1)))+1
				count := 0
				for m := rowMin; m < rowMax; m++ {
					for n := colMin; n < colMax; n++ {
						if m == i && n == j {
							continue
						}

						if paperMatrix[m][n] == '@' || paperMatrix[m][n] == 'M' {
							count++
						}
					}
				}

				if count < 4 {
					paperMatrix[i][j] = 'M'
					iterCount++
				}
			}
		}

		if iterCount > 0 {
			rollCount += iterCount
			for i := range len(paperMatrix) {
				for j := range len(paperMatrix[i]) {
					if paperMatrix[i][j] == 'M' {
						paperMatrix[i][j] = '.'
					}
				}
			}
		} else {
			break
		}
	}

	fmt.Println(rollCount)
}
