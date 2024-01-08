package http

import "github.com/amirhnajafiz/jester/pkg"

func (h Handler) processMetrics(req Request) {
	switch req.Field {
	case pkg.FieldAddPublisher:
	case pkg.FieldRemovePublisher:
	case pkg.FieldAddSubscriber:
	case pkg.FieldRemoveSubscriber:
	case pkg.FieldPublish:
	case pkg.FieldConsume:
	case pkg.FieldFailures:
	case pkg.FieldFailedConnections:
	case pkg.FieldRetryPerConnection:
	case pkg.FieldLatency:
	}
}
