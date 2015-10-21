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
		{{1, 3, 5}, {2, 4, 6}, {1, 2, 3, 4, 5, 6}},
		{{}, {}, {}},
		{{9}, {3}, {3, 9}},
		{{1}, {}, {1}},
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
	var cases = []struct{ input, expected}[]int}{
		{{}, {}},
		{{1}, {1}},
		{{1, 2}, {1, 2}},
		{{2, 1}, {1, 2}},
		{{2, 1, 2}, {1, 2, 2}},
		{{1, 5, 2, 6, 3, 8}, {1, 2, 3, 5, 6, 8}},
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
func BenchmarkSort_2(b *testing.B) {
	 for i:= 0; i< b.N; i++ {
		sort_2(a)
	}
}
