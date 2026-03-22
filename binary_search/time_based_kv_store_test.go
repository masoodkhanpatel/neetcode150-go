package binary_search

import "testing"

func TestTimeMap(t *testing.T) {
	timeMap := Constructor()
	timeMap.Set("foo", "bar", 1)
	if got := timeMap.Get("foo", 1); got != "bar" {
		t.Errorf("Get(foo, 1) = %s; want bar", got)
	}
	if got := timeMap.Get("foo", 3); got != "bar" {
		t.Errorf("Get(foo, 3) = %s; want bar", got)
	}
	timeMap.Set("foo", "bar2", 4)
	if got := timeMap.Get("foo", 4); got != "bar2" {
		t.Errorf("Get(foo, 4) = %s; want bar2", got)
	}
	if got := timeMap.Get("foo", 5); got != "bar2" {
		t.Errorf("Get(foo, 5) = %s; want bar2", got)
	}
}
