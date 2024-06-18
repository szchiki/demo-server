package direction

type Direction interface {
	GetDirection(src Coordinate, dst ...Coordinate) ([]DirectionLeg, error)
}
