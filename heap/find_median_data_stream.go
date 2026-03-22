package heap

import (
	"container/heap"
)

type MedianFinder struct {
	small *MaxIntHeap // Max-heap for the smaller half
	large *IntHeap    // Min-heap for the larger half
}

func ConstructorMedian() MedianFinder {
	mf := MedianFinder{
		small: &MaxIntHeap{},
		large: &IntHeap{},
	}
	heap.Init(mf.small)
	heap.Init(mf.large)
	return mf
}

func (this *MedianFinder) AddNum(num int) {
	heap.Push(this.small, num)

	// small.max <= large.min
	if this.small.Len() > 0 && this.large.Len() > 0 && (*this.small)[0] > (*this.large)[0] {
		val := heap.Pop(this.small).(int)
		heap.Push(this.large, val)
	}

	// Balance size
	if this.small.Len() > this.large.Len()+1 {
		val := heap.Pop(this.small).(int)
		heap.Push(this.large, val)
	} else if this.large.Len() > this.small.Len()+1 {
		val := heap.Pop(this.large).(int)
		heap.Push(this.small, val)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.small.Len() > this.large.Len() {
		return float64((*this.small)[0])
	} else if this.large.Len() > this.small.Len() {
		return float64((*this.large)[0])
	} else {
		return float64((*this.small)[0]+(*this.large)[0]) / 2.0
	}
}
