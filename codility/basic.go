package main

import (
	"fmt"
	"math"
)

var log = fmt.Println

func SmallestNotPresentPositiveInt(A []int) int {
	result := 1
	m := make(map[int]bool, len(A))
	for _, v := range A {
		m[v] = true
	}
	for i := 1; i <= len(A)+1; i++ {
		if _, ok := m[i]; !ok {
			result = i
			break
		}
	}
	return result
}

func MaxGap(N int) int {
	remainder := N
	var positions []int
	for remainder > 0 {
		pos := int(math.Log2(float64(remainder)))
		positions = append(positions, pos)
		remainder = remainder - int(math.Pow(float64(2), float64(pos)))
	}
	if len(positions) < 2 {
		return 0
	}
	max := 0
	for i := 0; i < len(positions)-1; i++ {
		if (positions[i] - positions[i+1] - 1) > max {
			max = positions[i] - positions[i+1] - 1
		}
	}
	return max
}

func Test_MaxGap() {
	N := 1041
	fmt.Printf("%b", N)
	log(MaxGap(N))
}
