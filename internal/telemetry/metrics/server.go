package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func NewHandler(srv *http.ServeMux) {
	srv.Handle("/metrics", promhttp.Handler())
}
