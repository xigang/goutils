package timeutils

import (
	"fmt"
	"testing"
	"time"
)

// == RUN TestDiffDays
// 2016-04-10 10:00:00 +0000 UTC
// --- PASS: TestDiffDays (0.00s)
// 	common_func_test.go:80: diff days: 10
// PASS
// ok  	github.com/zakka-cn/goutils	0.007s
func TestDiffDays(t *testing.T) {
	strt := "2016-04-10 10:00:00"
	t1, _ := time.Parse("2006-01-02 15:04:05", strt)
	fmt.Println(t1)
	t2 := time.Now()
	diffDays := DiffDays(t1, t2)
	t.Log("diff days:", diffDays)
}
