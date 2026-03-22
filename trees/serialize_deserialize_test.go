package trees

import (
	"reflect"
	"testing"
)

func TestSerializeDeserialize(t *testing.T) {
	tests := []struct {
		root []int
	}{
		{[]int{1, 2, 3, -1, -1, 4, 5}},
		{[]int{}},
	}
	for _, tt := range tests {
		root := SliceToTree(tt.root)
		ser := Constructor()
		deser := Constructor()
		data := ser.serialize(root)
		ans := deser.deserialize(data)
		got := TreeToSlice(ans)
		if !reflect.DeepEqual(got, tt.root) {
			t.Errorf("serialize/deserialize failed, expected %v, got %v", tt.root, got)
		}
	}
}
