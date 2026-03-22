package binary_search

type valueStruct struct {
	val       string
	timestamp int
}

type TimeMap struct {
	store map[string][]valueStruct
}

func Constructor() TimeMap {
	return TimeMap{
		store: make(map[string][]valueStruct),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.store[key] = append(this.store[key], valueStruct{value, timestamp})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	values, ok := this.store[key]
	if !ok {
		return ""
	}

	l, r := 0, len(values)-1
	res := ""

	for l <= r {
		m := l + (r-l)/2
		if values[m].timestamp <= timestamp {
			res = values[m].val
			l = m + 1
		} else {
			r = m - 1
		}
	}

	return res
}
