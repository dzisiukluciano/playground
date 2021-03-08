package main

import (
	"fmt"
)

var p = fmt.Println

func mergesort(a []int) []int {
	if len(a) <= 1 {
		return a
	}
	mid := int(len(a) / 2)
	p("mid", mid)
	left := mergesort(a[:mid])
	right := mergesort(a[mid:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	a := make([]int, len(left)+len(right))
	lidx := 0
	ridx := 0
	for i := 0; i < len(a); i++ {
		if lidx >= len(left) {
			a[i] = right[ridx]
			ridx++
		} else if ridx >= len(right) {
			a[i] = left[lidx]
			lidx++
		} else {
			if left[lidx] <= right[ridx] {
				a[i] = left[lidx]
				lidx++
			} else {
				a[i] = right[ridx]
				ridx++
			}
		}
	}
	return a
}

func main() {
	a1 := []int{5, 7, 3, 8, 9, 1, 23, 10}
	fmt.Println(mergesort(a1))
}
