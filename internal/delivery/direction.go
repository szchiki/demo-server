package delivery

type Direction interface {
	GetDirection(src Coordinate, dst Coordinate) ([]Distance, error)
}
