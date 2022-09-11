package kit

import (
	"testing"

	kit "github.com/xiiapp/kit/sctucture/type"
)

// cmd: go test -v str_slice_test.go str_slice.go -test.run TestStrSlice
func TestIsStrSlice(t *testing.T) {
	type v []string
	v1 := &v{"a", "b"}
	// if !IsStrSlice(v1) {
	// 	t.Error("IsStrSlice error")
	// }

	if kit.IsSlice(v1) {
		t.Error("kit.IsSlice error")
	}

	v2 := &StrSlice{"a", "b"}
	if !IsStrSlice(v2) {
		t.Error("IsStrSlice error")
	}

}
