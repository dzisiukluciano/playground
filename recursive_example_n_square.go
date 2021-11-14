package main

import (
	"fmt"
)

// key characteristic of these functions: we iterate the entire array, and OUTSIDE the iteration (at the end)
// we re-process (recursion) n-1 elements
func fn(n int, it *int) {
	if n == 0 {
		return
	}
	for i:=0; i<n; i++ {
		*it++
	}
	fn(n-1, it)
}

// example of worst case scenario of quick sort, where we always choose the greatest element,
// therefore dividing in sub-arrays with n-1 elements
// technically, time complexity is O(n^2) because is n+(n-1)+(n-2)+...+1 = Sum(1 to n) of n = n*(n+1)/2 =~ (1/2)(n*n) =~ n^2
func main() {
	it := 0
	n := 10
	fn(n, &it)
	fmt.Println(it)
}
