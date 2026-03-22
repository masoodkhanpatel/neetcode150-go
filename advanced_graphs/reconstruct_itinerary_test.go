package advanced_graphs

import (
	"reflect"
	"testing"
)

func TestFindItinerary(t *testing.T) {
	tests := []struct {
		name     string
		tickets  [][]string
		expected []string
	}{
		{
			"Example 1",
			[][]string{{"MUC", "LHR"}, {"JFK", "MUC"}, {"SFO", "SJC"}, {"LHR", "SFO"}},
			[]string{"JFK", "MUC", "LHR", "SFO", "SJC"},
		},
		{
			"Example 2",
			[][]string{{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "ATL"}, {"ATL", "JFK"}, {"ATL", "SFO"}},
			[]string{"JFK", "ATL", "JFK", "SFO", "ATL", "SFO"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FindItinerary(tt.tickets)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("FindItinerary() = %v; want %v", result, tt.expected)
			}
		})
	}
}
