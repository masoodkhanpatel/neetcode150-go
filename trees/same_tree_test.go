package trees

import (
	"testing"
)

func TestIsSameTree(t *testing.T) {
	tests := []struct {
		p    []int
		q    []int
		want bool
	}{
		{[]int{1, 2, 3}, []int{1, 2, 3}, true},
		{[]int{1, 2}, []int{1, -1, 2}, false},
		{[]int{1, 2, 1}, []int{1, 1, 2}, false},
	}
	for _, tt := range tests {
		p := SliceToTree(tt.p)
		q := SliceToTree(tt.q)
		if got := isSameTree(p, q); got != tt.want {
			t.Errorf("isSameTree() = %v, want %v", got, tt.want)
		}
	}
}
