package delivery

import "sort"

type PickupOptions struct {
	Src          string                     `json:"source"`
	Destinations []PickupOptionsDestination `json:"routes"`
}

type PickupOptionsDestination struct {
	Duration float64 `json:"duration"`
	Distance float64 `json:"distance"`
	Dst      string  `json:"destination"`
}

func SortPickupDestinations(dst []PickupOptionsDestination) {
	sort.Slice(dst, func(i, j int) bool {
		if dst[i].Duration != dst[j].Duration {
			return dst[i].Duration < dst[j].Duration
		}
		return dst[i].Distance < dst[j].Distance
	})
}
