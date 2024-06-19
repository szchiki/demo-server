package delivery

import (
	"errors"
	"ingrid/internal/delivery"
)

type Service struct {
	directionProvider delivery.Direction
}

func NewService(d delivery.Direction) *Service {
	return &Service{directionProvider: d}
}

func (s *Service) GetDirection(src delivery.Coordinate, destinations ...delivery.Coordinate) (*delivery.PickupOptions, error) {
	if len(destinations) < 1 {
		return nil, errors.New("destinations_missing_min_1")
	}
	pickup := &delivery.PickupOptions{Src: src.String()}
	for _, dst := range destinations {
		legs, err := s.directionProvider.GetDirection(src, dst)
		if err != nil {
			return nil, err
		}
		dstOpt := delivery.PickupOptionsDestination{Dst: dst.String()}
		for _, leg := range legs {
			dstOpt.Distance += leg.Distance
			dstOpt.Duration += leg.Duration
		}
		pickup.Destinations = append(pickup.Destinations, dstOpt)
	}
	delivery.SortPickupDestinations(pickup.Destinations)
	return pickup, nil
}
