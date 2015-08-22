package main

import "fmt"

func sort(s []int) []int {
	if (len(s) <= 1) {
		return s
	}
	middle := len(s)/2
	a, b := s[:middle], s[middle:]
	sa, sb := sort(a), sort(b)
	return merge(sa,sb)
}

func merge(a, b []int) []int {
	res := make([]int, len(a)+ len(b))
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

func sort_2(s []int) {
	if (len(s) <= 1) {
		return
	}
	middle := len(s)/2
	a, b := s[:middle], s[middle:]
	sort_2(a)
	sort_2(b)
	merge_2(s,a,b)
}

func merge_2(dst, a, b []int){
	var i, j, k int
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			dst[k] = a[i]
			i++
		} else {
			dst[k] = b[j]
			j++
		}
		k++
	}
	for i < len(a) {
		dst[k] = a[i]
		i++
		k++
	}
	for j < len(b) {
		dst[k] = b[j]
		j++
		k++
	}
	return
}


func main() {
	a:= []int{1,9,2,8,3,7,4,6,5}
	fmt.Println(sort(a))
	fmt.Println(a)
	sort_2(a)
	fmt.Println(a)
}
