package transport

import (
	"errors"
	"github.com/labstack/echo/v4"
	"ingrid/internal/delivery"
	directionService "ingrid/services/delivery"
	"strings"
)

type Server struct {
	e *echo.Echo
	d *directionService.Service
}

func NewServer(ds *directionService.Service) *Server {
	s := Server{
		e: echo.New(),
		d: ds,
	}

	s.e.GET("/delivery", s.endpointDirectionsGet)

	return &s
}

func (s *Server) Start() error {
	return s.e.Start(":8080")
}

// http://localhost:8080/direction?src=13.388860,52.517037&dst=13.397634,52.529407&dst=13.428555,52.523219
type DirectionRequest struct {
	Src string   `query:"src"`
	Dst []string `query:"dst"`
}

func parseCoordinate(s string) (delivery.Coordinate, error) {
	parts := strings.Split(s, ",")
	if len(parts) != 2 {
		return delivery.Coordinate{}, errors.New("invalid coordinate")
	}
	return delivery.Coordinate{
		Lat: parts[0],
		Lng: parts[1],
	}, nil
}

// src 59.336451,18.061630
// dst 59.349957,18.001322;59.310650,18.112240;59.337336,17.933603
func (s *Server) endpointDirectionsGet(c echo.Context) error {
	//TODO: Get start and stop (1 start, multiple stops)
	req := &DirectionRequest{}
	if err := c.Bind(req); err != nil {
		return err
	}
	src, err := parseCoordinate(req.Src)
	if err != nil {
		return err
	}
	destinations := []delivery.Coordinate{}
	for _, dst := range req.Dst {
		d, err := parseCoordinate(dst)
		if err != nil {
			return err
		}
		destinations = append(destinations, d)
	}
	if len(destinations) == 0 {
		return errors.New("no destinations")
	}
	//TODO: Get streets
	legs, err := s.d.GetDirection(src, destinations...)
	if err != nil {
		return err
	}
	_ = legs
	//TODO: Calculate shortest distance
	return c.JSON(200, legs)
}
