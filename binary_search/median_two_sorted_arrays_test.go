package binary_search

import "testing"

func TestFindMedianSortedArrays(t *testing.T) {
	tests := []struct {
		nums1 []int
		nums2 []int
		want  float64
	}{
		{[]int{1, 3}, []int{2}, 2.0},
		{[]int{1, 2}, []int{3, 4}, 2.5},
	}

	for _, tt := range tests {
		got := findMedianSortedArrays(tt.nums1, tt.nums2)
		if got != tt.want {
			t.Errorf("findMedianSortedArrays(%v, %v) = %f; want %f", tt.nums1, tt.nums2, got, tt.want)
		}
	}
}
