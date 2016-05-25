package join_bench

import (
	"strings"
	"testing"
)

// Cuidado con las optimizaciones, el compilador
// puede eliminar código que no vea que tiene un
// uso posterior

const maxLen = 1e3

var strs = make([]string, 0, maxLen)

func TestMain(m *testing.M) {
	for i := 0; i < maxLen; i++ {
		strs = append(strs, strings.Repeat("x", i))
	}
	m.Run()
}

func BenchmarkConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var res, sep string
		for _, s := range strs {
			res += sep + s
			sep = " "
		}
		_ = res
	}
}

func BenchmarkJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		res := strings.Join(strs, " ")
		_ = res
	}
}
