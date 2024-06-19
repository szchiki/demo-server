package main

import (
	"ingrid/cmd/transport"
	"ingrid/internal/delivery/direction/OSRM"
	"ingrid/services/delivery"
	"log"
)

func main() {
	maps := OSRM.NewClient()
	s := delivery.NewService(maps)

	srv := transport.NewServer(s)
	if err := srv.Start(); err != nil {
		log.Fatalln(err)
	}
}
