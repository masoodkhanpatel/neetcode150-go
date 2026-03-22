package binary_search

import "math"

// Median of Two Sorted Arrays
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	A, B := nums1, nums2
	total := len(nums1) + len(nums2)
	half := total / 2

	if len(B) < len(A) {
		A, B = B, A
	}

	l, r := 0, len(A)-1
	for {
		i := int(math.Floor(float64(l+r) / 2)) // A
		j := half - i - 2                      // B

		var AL, AR, BL, BR float64

		if i >= 0 {
			AL = float64(A[i])
		} else {
			AL = math.Inf(-1)
		}

		if (i + 1) < len(A) {
			AR = float64(A[i+1])
		} else {
			AR = math.Inf(1)
		}

		if j >= 0 {
			BL = float64(B[j])
		} else {
			BL = math.Inf(-1)
		}

		if (j + 1) < len(B) {
			BR = float64(B[j+1])
		} else {
			BR = math.Inf(1)
		}

		if AL <= BR && BL <= AR {
			if total%2 != 0 {
				return math.Min(AR, BR)
			}
			return (math.Max(AL, BL) + math.Min(AR, BR)) / 2
		} else if AL > BR {
			r = i - 1
		} else {
			l = i + 1
		}
	}
}
