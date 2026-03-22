package arrays_hashing

import (
	"reflect"
	"testing"
)

func TestEncodeDecode(t *testing.T) {
	tests := []struct {
		strs []string
	}{
		{[]string{"hello", "world"}},
		{[]string{""}},
		{[]string{"", ""}},
		{[]string{"long_string_with_many_characters", "short"}},
		{[]string{"#", "##", "###"}},
		{[]string{"123#abc", "456#def"}},
	}

	for _, tt := range tests {
		encoded := Encode(tt.strs)
		decoded := Decode(encoded)
		if !reflect.DeepEqual(tt.strs, decoded) {
			t.Errorf("Encode/Decode failed: expected %v, got %v (encoded: %s)", tt.strs, decoded, encoded)
		}
	}
}
