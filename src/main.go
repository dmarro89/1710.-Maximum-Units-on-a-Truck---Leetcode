package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maximumUnits([][]int{{1, 3}, {2, 2}, {3, 1}}, 4))
}

func maximumUnits(boxTypes [][]int, truckSize int) int {
	sort.SliceStable(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] > boxTypes[j][1]
	})

	sum := 0
	counter := truckSize
	for _, box := range boxTypes {
		if counter == 0 {
			break
		}

		avBox := counter
		if counter-box[0] > 0 {
			avBox = box[0]
		}

		sum += avBox * box[1]
		counter = counter - avBox
	}

	return sum
}
