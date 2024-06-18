package main

import (
	"ingrid/cmd/transport"
	"ingrid/internal/direction/OSRM"
	"ingrid/services/direction"
	"log"
)

func main() {
	maps := OSRM.NewClient()
	s := direction.NewService(maps)

	srv := transport.NewServer(s)
	if err := srv.Start(); err != nil {
		log.Fatalln(err)
	}
}
