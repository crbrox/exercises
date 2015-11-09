package main

import "fmt"

func sort(s []int) []int {
	if len(s) <= 1 {
		return s
	}
	middle := len(s) / 2
	a, b := s[:middle], s[middle:]
	sa, sb := sort(a), sort(b)
	return merge(sa, sb)
}

func merge(a, b []int) []int {
	res := make([]int, len(a)+len(b))
	var i, j, k int
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			res[k] = a[i]
			i++
		} else {
			res[k] = b[j]
			j++
		}
		k++
	}
	for i < len(a) {
		res[k] = a[i]
		i++
		k++
	}
	for j < len(b) {
		res[k] = b[j]
		j++
		k++
	}
	return res
}

func main() {
	a := []int{1, 9, 2, 8, 3, 7, 4, 6, 5}
	fmt.Println("version original", a)
	fmt.Println("version ordenada", sort(a))

}
