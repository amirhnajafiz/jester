package http

import (
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

}

func (h Handler) Register(port int) error {
	http.HandleFunc("/healthz", h.healthy)
	http.HandleFunc("/cover", h.cover)

	log.Println(fmt.Sprintf("http server running on %d...", port))

	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
