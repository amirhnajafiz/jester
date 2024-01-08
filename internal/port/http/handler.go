package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Handler struct{}

func (h Handler) healthy(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("ok"))
}

func (h Handler) cover(w http.ResponseWriter, r *http.Request) {
	var req Request

	// get request body
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	h.processMetrics(req)

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("ok"))
}

func (h Handler) Register(port int) error {
	mux := http.NewServeMux()

	mux.HandleFunc("/healthz", h.healthy)
	mux.HandleFunc("/readyz", h.healthy)
	mux.HandleFunc("/cover", h.cover)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}

	log.Println(fmt.Sprintf("http server running on %d...", port))

	return server.ListenAndServe()
}
