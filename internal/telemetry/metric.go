package telemetry

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Server struct {
	srv     *http.ServeMux
	address string
}

func NewServer(cfg Metric) Server {
	var srv *http.ServeMux

	if cfg.Enabled {
		srv = http.NewServeMux()
		srv.Handle("/metrics", promhttp.Handler())
	}

	return Server{
		address: cfg.Address,
		srv:     srv,
	}
}
