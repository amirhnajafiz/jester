package metrics

import (
	"errors"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Server struct {
	srv     *http.ServeMux
	address string
}

func NewServer(cfg Config) Server {
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

func (s Server) Start() {
	go func() {
		if err := http.ListenAndServe(s.address, s.srv); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("metric server initiation failed: %v\n", err)
		}
	}()
}
