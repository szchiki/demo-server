package main

import (
	"fmt"
	"ingrid/cmd/transport"
	"ingrid/internal/delivery/direction/OSRM"
	"ingrid/services/delivery"
	"log"
	"os"
)

func main() {
	port := ":8080"
	if p := os.Getenv("PORT"); p != "" {
		port = fmt.Sprintf(":%s", os.Getenv(p))
	}
	maps := OSRM.NewClient()
	s := delivery.NewService(maps)

	srv := transport.NewServer(s)
	if err := srv.Start(port); err != nil {
		log.Fatalln(err)
	}
}
