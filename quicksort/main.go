package main

import (
	"fmt"
	"math/rand"
)

func quickSort(a []int, iteration *int) []int {
	if len(a) < 2 {
		return a
	}
	left := 0
	right := len(a) - 1
	i := rand.Intn(len(a))
	pivot := a[i]
	// importante, si uso random, muevo el pivot al final
	a[i], a[right] = a[right], a[i]
	fmt.Println("pivot ", pivot)
	for pos, v := range a {
		*iteration++
		if v < pivot {
			a[left], a[pos] = a[pos], a[left]
			left++
		}
	}
	// cuando termino de recorrer el array, meto el pivot en donde quedó left
	a[left], a[right] = a[right], a[left]
	// el pivot quedó en su posición, ahora divido el problema en 2 subproblemas
	if left > 0 {
		quickSort(a[0:left], iteration)
	}
	if left < right {
		quickSort(a[left+1:], iteration)
	}
	return a
}

func main() {
	a1 := []int{7,5,1,8,3,2,5,9,22,23,20}
	a2 := []int{7,1,8,3,2,5,9}

	it1 := 0
	fmt.Println(quickSort(a1, &it1))
	fmt.Println(it1, " iterations with n = ", len(a1))
	fmt.Println("------ 2nd example ------")
	it2 := 0
	fmt.Println(quickSort(a2, &it2))
	fmt.Println(it2, " iterations with n = ", len(a2))
}
