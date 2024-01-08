package http

import "github.com/amirhnajafiz/jester/pkg"

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
	case pkg.FieldConsume:
		h.Metrics.IncreaseNumberOfConsume(req.Label)
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
