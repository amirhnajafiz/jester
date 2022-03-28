package cmd

import (
	"log"

	"github.com/amirhnajafiz/Stan-Gee/internal/cmd/server"
	"github.com/amirhnajafiz/Stan-Gee/internal/config"
	"github.com/amirhnajafiz/Stan-Gee/internal/http/stan"
	"github.com/amirhnajafiz/Stan-Gee/internal/telemetry"
)

func Execute() {
	// load configurations
	cfg := config.Load()
	// tracer and metric init
	tracer := telemetry.New(cfg.Telemetry.Trace)
	metric := stan.NewMetrics()

	// start jaeger server
	telemetry.NewServer(cfg.Telemetry.Metric).Start()

	// get gRPC server
	s, l := server.NewServer(cfg.Server, tracer, metric)

	// start gRPC server
	err := s.Serve(l)
	if err != nil {
		log.Fatalf("failed to create serer : %v\n", err)
	}
}
