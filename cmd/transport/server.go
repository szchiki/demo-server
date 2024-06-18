package transport

import (
	"github.com/labstack/echo/v4"
	"ingrid/services/direction"
)

type Server struct {
	e *echo.Echo
}

func NewServer(ds direction.Service) *Server {
	e := echo.New()

	e.GET("/direction", func(c echo.Context) error { return nil })
	return &Server{e: echo.New()}
}

func endpointDirectionsGet(c echo.Context) error {
	//TODO: Get start and stop (1 start, multiple stops)
	//TODO: Get streets
	//TODO: Calculate shortest distance
	return nil
}
