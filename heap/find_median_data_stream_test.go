package heap

import "testing"

func TestMedianFinder(t *testing.T) {
	mf := ConstructorMedian()
	mf.AddNum(1)
	mf.AddNum(2)
	if mf.FindMedian() != 1.5 {
		t.Errorf("Median for [1, 2] should be 1.5, got %v", mf.FindMedian())
	}
	mf.AddNum(3)
	if mf.FindMedian() != 2.0 {
		t.Errorf("Median for [1, 2, 3] should be 2.0, got %v", mf.FindMedian())
	}
}
