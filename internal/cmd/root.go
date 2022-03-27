package cmd

import (
	"log"

	"github.com/amirhnajafiz/Stan-Gee/internal/cmd/server"
	"github.com/amirhnajafiz/Stan-Gee/internal/config"
)

func Execute() {
	cfg := config.Load()
	s, l := server.NewServer(cfg.Server)

	err := s.Serve(l)
	if err != nil {
		log.Fatalf("failed to create serer : %v\n", err)
	}
}
