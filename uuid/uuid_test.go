package uuid

import (
	"testing"
)

// === RUN TestUUID
// --- PASS: TestUUID (0.00s)
// 	uuid_test.go:12: 0e48f4e8-9fc8-4d7e-5e0c-12a7f5972997
// PASS
// ok  	github.com/xigang/goutils/uuid	0.005s
func TestUUID(t *testing.T) {
	uuid, err := UUID()
	if err != nil {
		t.Error(err)
	}
	t.Log(uuid)
}
