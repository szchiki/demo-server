package delivery

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSortPickupDestinations(t *testing.T) {
	cases := []struct {
		name     string
		input    []PickupOptionsDestination
		expected []PickupOptionsDestination
	}{
		{
			name: "sort by duration",
			input: []PickupOptionsDestination{
				{Duration: 10, Distance: 10},
				{Duration: 8, Distance: 10},
				{Duration: 12, Distance: 10},
			},
			expected: []PickupOptionsDestination{
				{Duration: 8, Distance: 10},
				{Duration: 10, Distance: 10},
				{Duration: 12, Distance: 10},
			},
		},
		{
			name: "sort by distance",
			input: []PickupOptionsDestination{
				{Duration: 10, Distance: 10},
				{Duration: 10, Distance: 12},
				{Duration: 10, Distance: 8},
			},
			expected: []PickupOptionsDestination{
				{Duration: 10, Distance: 8},
				{Duration: 10, Distance: 10},
				{Duration: 10, Distance: 12},
			},
		},
		{
			name: "mix and match",
			input: []PickupOptionsDestination{
				{Duration: 10, Distance: 10},
				{Duration: 8, Distance: 12},
				{Duration: 10, Distance: 6},
				{Duration: 8, Distance: 8},
			},
			expected: []PickupOptionsDestination{
				{Duration: 8, Distance: 8},
				{Duration: 8, Distance: 12},
				{Duration: 10, Distance: 6},
				{Duration: 10, Distance: 10},
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			// Loop variables can no longer be safely modified inside the scope
			pickupOpts := make([]PickupOptionsDestination, len(tc.input))
			copy(pickupOpts, tc.input)
			SortPickupDestinations(pickupOpts)
			if diff := cmp.Diff(tc.expected, pickupOpts); diff != "" {
				t.Errorf("unexpected result (-want +got):\n%s", diff)
			}
		})
	}
}
