package cmd

import (
	"github.com/amirhnajafiz/Stan-Gee/internal/http/stan"
	"github.com/amirhnajafiz/Stan-Gee/internal/telemetry"
	"log"

	"github.com/amirhnajafiz/Stan-Gee/internal/cmd/server"
	"github.com/amirhnajafiz/Stan-Gee/internal/config"
)

func Execute() {
	cfg := config.Load()
	tracer := telemetry.New(cfg.Telemetry.Trace)
	telemetry.NewServer(cfg.Telemetry.Metric).Start()

	metric := stan.NewMetrics()

	s, l := server.NewServer(cfg.Server, tracer, metric)

	err := s.Serve(l)
	if err != nil {
		log.Fatalf("failed to create serer : %v\n", err)
	}
}
