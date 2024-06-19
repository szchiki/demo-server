package delivery

// Coordinate holds coordinates for directions. While small float errors likely won't affect the absolute distance a lot, it's possible it'll create some mismatches in comparisons
type Coordinate struct {
	Lat string
	Lng string
}

func (c Coordinate) String() string {
	return c.Lat + "," + c.Lng
}
