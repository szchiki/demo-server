package delivery

type Direction interface {
	GetDirection(src Coordinate, dst Coordinate) ([]Distance, error)
}

type Distance struct {
	Duration float64 `json:"duration"`
	Distance float64 `json:"distance"`
}
