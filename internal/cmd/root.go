package cmd

import (
	"log"

	"github.com/amirhnajafiz/Stan-Gee/internal/cmd/server"
)

func Execute() {
	s, l := server.NewServer(":8080")

	err := s.Serve(l)
	if err != nil {
		log.Fatalf("failed to create serer : %v\n", err)
	}
}
