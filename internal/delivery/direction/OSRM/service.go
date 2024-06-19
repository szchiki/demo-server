package OSRM

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"ingrid/internal/delivery"
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

func (s *Service) GetDirection(src delivery.Coordinate, dst delivery.Coordinate) ([]delivery.Distance, error) {
	baseUrl := `http://router.project-osrm.org/route/v1/driving/%s?overview=false`
	coords := []string{src.String(), dst.String()}
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
	legs := []delivery.Distance{}
	for _, leg := range routes.Routes[0].Legs {
		legs = append(legs, delivery.Distance{
			Distance: leg.Distance,
			Duration: leg.Duration,
		})
	}
	return legs, nil
}
