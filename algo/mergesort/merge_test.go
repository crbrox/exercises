package main

import ("testing"; "reflect"; "math/rand")

var a []int

func TestMain(t *testing.T) {
	     const long = 1024*1024*10
		//a very big array for testing
		a = make([]int, long)
		for i := range a {
			a[i] = rand.Intn(100)
		}

}
func TestMerge(t *testing.T) {
	var cases = []struct{a,b,expected []int}{
		{[]int{1, 3, 5}, []int{2, 4, 6}, []int{1, 2, 3, 4, 5, 6}},
		{[]int{}, []int{}, []int{}},
		{[]int{9}, []int{3}, []int{3, 9}},
		{[]int{1}, []int{}, []int{1}},
	}
	for i, c := range cases {
		actual := merge(c.a, c.b)
		if !reflect.DeepEqual(actual, c.expected) {
			t.Fatalf("case %d: should %#v, actual %#v\n", i, c.expected, actual)
		}
		// And it should be commutative
		actual = merge(c.b, c.a)
		if !reflect.DeepEqual(actual, c.expected) {
			t.Fatalf("case %d (reverse): should %#v, actual %#v\n", i, c.expected, actual)
		}
	}
}
func TestSort(t *testing.T, ) {
	var cases = []struct{ input, expected []int}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 2}, []int{1, 2}},
		{[]int{2, 1}, []int{1, 2}},
		{[]int{2, 1, 2}, []int{1, 2, 2}},
		{[]int{1, 5, 2, 6, 3, 8}, []int{1, 2, 3, 5, 6, 8}},
	}
	for i, c := range cases {
		actual := sort(c.input)
		if !reflect.DeepEqual(actual, c.expected) {
			t.Fatalf("case %d: should %#v, actual %#v\n", i, c.expected, actual)
		}
	}
}

func BenchmarkSort(b *testing.B) {
	 for i:= 0; i< b.N; i++ {
		sort(a)
	}
}

