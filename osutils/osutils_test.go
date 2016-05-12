package osutils

import (
	"testing"
)

// ➜  osutils git:(master) ✗ go test -v
// === RUN TestIsDirExists
// --- PASS: TestIsDirExists (0.00s)
// 	osutils_test.go:11: this path exist.
// PASS
// ok  	github.com/xigang/goutils/osutils	0.004s
func TestIsDirExists(t *testing.T) {
	path := "/usr/local"

	if IsDirExists(path) {
		t.Log("this path exist.")
	} else {
		t.Log("this path is not exist.")
	}
}
