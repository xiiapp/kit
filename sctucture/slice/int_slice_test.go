package kit

import (
	"testing"

	kit "github.com/xiiapp/kit/sctucture/type"
)

// cmd: go test -v int_slice_test.go int_slice.go -test.run TestIsIntSlice
func TestIsIntSlice(t *testing.T) {
	type v []int
	v1 := &v{1, 2, 3}
	if kit.IsSlice(v1) {
		t.Error("kit.IsSlice error")
	}

	v2 := &IntSlice{1, 2, 3, 4}
	if !IsIntSlice(v2) {
		t.Error("IsIntSlice error")
	}

	if IsIntSlice(&[]int{1, 2, 3}) {
		t.Error("IsIntSlice error", &[]int{1, 2, 3})
	}

}
