package math

import (
	"fmt"
	"testing"
)

//output:
//=== RUN TestToFixed
// --- PASS: TestToFixed (0.00s)
// 	common_func_test.go:20: 1
// 	common_func_test.go:20: 1.2
// 	common_func_test.go:20: 1.23
// 	common_func_test.go:20: 1.235
// PASS
// ok  	github.com/zakka-cn/goutils	0.015s
func TestToFixed(t *testing.T) {
	var numberList = []struct {
		number    float64
		precision int
	}{
		{1.2345678, 0},
		{1.2345678, 1},
		{1.2345678, 2},
		{1.2345678, 3},
	}

	for _, v := range numberList {
		num := ToFixed(v.number, v.precision)
		t.Log(num)
	}
}

//output:
//=== RUN TestVersionOrdinal
// 1.05.00.0156 > 1.0.221.9289
// 1 < 1.0.1
// 1.0.1 < 1.0.2
// 1.0.2 < 1.0.3
// 1.0.3 < 1.1
// 1.1 < 1.1.1
// 1.1.1 < 1.1.2
// 1.1.2 < 1.2
// --- PASS: TestVersionOrdinal (0.00s)
// PASS
// ok  	github.com/zakka-cn/goutils	0.007s

func TestVersionOrdinal(t *testing.T) {
	versions := []struct{ a, b string }{
		{"1.05.00.0156", "1.0.221.9289"},
		{"1", "1.0.1"},
		{"1.0.1", "1.0.2"},
		{"1.0.2", "1.0.3"},
		{"1.0.3", "1.1"},
		{"1.1", "1.1.1"},
		{"1.1.1", "1.1.2"},
		{"1.1.2", "1.2"},
	}

	for _, version := range versions {
		a, b := VersionOrdinal(version.a), VersionOrdinal(version.b)
		switch {
		case a > b:
			fmt.Println(version.a, ">", version.b)
		case a < b:
			fmt.Println(version.a, "<", version.b)
		case a == b:
			fmt.Println(version.a, "=", version.b)
		}
	}
}
