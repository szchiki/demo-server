package direction

import "ingrid/internal/direction"

type Service struct {
	directionProvider direction.Direction
}

func NewService(d direction.Direction) *Service {
	return &Service{directionProvider: d}
}

func (s *Service) GetDirection(src direction.Coordinate, dst ...direction.Coordinate) ([]direction.DirectionLeg, error) {
	legs, err := s.directionProvider.GetDirection(src, dst...)
	if err != nil {
		return nil, err
	}
	return legs, nil
}
