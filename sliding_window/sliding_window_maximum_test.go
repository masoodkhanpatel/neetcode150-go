package sliding_window

import (
	"reflect"
	"testing"
)

func TestMaxSlidingWindow(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want []int
	}{
		{[]int{1, 3, -1, -3, 5, 3, 6, 7}, 3, []int{3, 3, 5, 5, 6, 7}},
		{[]int{1}, 1, []int{1}},
	}

	for _, tt := range tests {
		got := maxSlidingWindow(tt.nums, tt.k)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("maxSlidingWindow(%v, %d) = %v; want %v", tt.nums, tt.k, got, tt.want)
		}
	}
}
