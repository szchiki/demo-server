package delivery

// Coordinate holds coordinates for directions. While small float errors likely won't affect the absolute distance a lot, it's possible it'll create some mismatches in comparisons
type Coordinate struct {
	Lat string
	Lng string
}

func (c Coordinate) String() string {
	return c.Lat + "," + c.Lng
}

type Distance struct {
	Duration float64 `json:"duration"`
	Distance float64 `json:"distance"`
}

type PickupOptions struct {
	Src          string                     `json:"source"`
	Destinations []PickupOptionsDestination `json:"routes"`
}

type PickupOptionsDestination struct {
	Distance
	Dst string `json:"destination"`
}
