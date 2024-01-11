package http

import (
	"log"
	"time"

	"github.com/amirhnajafiz/jester/pkg"
)

func (h Handler) processMetrics(req Request) {
	switch req.Field {
	case pkg.FieldAddPublisher:
		h.Metrics.UpdateNumberOfPublishers(req.Label, 1)
	case pkg.FieldRemovePublisher:
		h.Metrics.UpdateNumberOfPublishers(req.Label, -1)
	case pkg.FieldAddSubscriber:
		h.Metrics.UpdateNumberOfSubscribers(req.Label, 1)
	case pkg.FieldRemoveSubscriber:
		h.Metrics.UpdateNumberOfSubscribers(req.Label, -1)
	case pkg.FieldPublish:
		h.Metrics.IncreaseNumberOfPublish(req.Label)

		if err := h.ETCD.Put(req.Label, time.Now().Format(time.DateTime)); err != nil {
			log.Println(err)
		}
	case pkg.FieldConsume:
		h.Metrics.IncreaseNumberOfConsume(req.Label)

		tmp, err := h.ETCD.Get(req.Label)
		if err != nil {
			log.Println(err)

			return
		}

		date, err := time.Parse(time.DateTime, tmp)
		if err != nil {
			log.Println(err)

			return
		}

		h.Metrics.AddLatency(req.Label, float64(time.Now().Sub(date).Milliseconds()))
	case pkg.FieldFailures:
		h.Metrics.IncreaseNumberOfFailures(req.Label)
	case pkg.FieldFailedConnections:
		h.Metrics.IncreaseFailedConnections()
	case pkg.FieldRetryPerConnection:
		h.Metrics.IncreaseRetryPerConnection(req.Value)
	case pkg.FieldLatency:
		h.Metrics.AddLatency(req.Label, req.Value)
	}
}
