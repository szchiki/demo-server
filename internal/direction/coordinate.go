package direction

// Coordinate holds coordinates for directions. While small float errors likely won't affect the absolute distance a lot, it's possible it'll create some mismatches in comparisons
type Coordinate struct {
	Lat string
	Lng string
}

// DirectionLeg is a utitlity struct to hold lists of movement for us. It assumes the only mode is driving.
type DirectionLeg struct {
	Src      Coordinate
	Dst      Coordinate
	Distance float64
	Duration float64
}
