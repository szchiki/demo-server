package OSRM

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"ingrid/internal/direction"
)

// Service lets us interact with the OSRM API. Creating this http client lets us store configurations specific to this integration.
type Service struct {
	c http.Client
}

func NewClient() *Service {
	return &Service{c: http.Client{}}
}

type ResponseRoute struct {
	Code   string `json:"code"`
	Routes []struct {
		Legs []struct {
			Duration float64 `json:"duration"`
			Distance float64 `json:"distance"`
		} `json:"legs"`
	} `json:"routes"`
}

func (s *Service) GetDirection(src direction.Coordinate, dst ...direction.Coordinate) ([]direction.DirectionLeg, error) {
	baseUrl := `http://router.project-osrm.org/route/v1/driving/%s?overview=false`
	coords := []string{}
	// While we could implement a default string method on the coordinate, I'd rather keep this very specific format to this service, the format may differ slightly in other integrations.
	coords = append(coords, fmt.Sprintf("%s,%s", src.Lat, src.Lng))
	for _, d := range dst {
		coords = append(coords, fmt.Sprintf("%s,%s", d.Lat, d.Lng))
	}
	url := fmt.Sprintf(baseUrl, strings.Join(coords, ";"))
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.c.Do(req)
	if err != nil {
		return nil, err
	}
	routes := ResponseRoute{}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&routes); err != nil {
		return nil, err
	}
	legs := []direction.DirectionLeg{}
	for _, leg := range routes.Routes[0].Legs {
		legs = append(legs, direction.DirectionLeg{
			Distance: leg.Distance,
			Duration: leg.Duration,
		})
	}
	return legs, nil
}
