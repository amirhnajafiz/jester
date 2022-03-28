package stan

import (
	"errors"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	Namespace = "stan_gee"
	Subsystem = "stan"
)

type Metrics struct {
	ConnectionErrors prometheus.Counter
	SuccessfulSub    prometheus.Counter
	FailedSub        prometheus.Counter
	SuccessfulPut    prometheus.Counter
	FailedPut        prometheus.Counter
}

func newCounter(counterOpts prometheus.CounterOpts) prometheus.Counter {
	ev := prometheus.NewCounter(counterOpts)

	if err := prometheus.Register(ev); err != nil {
		var are prometheus.AlreadyRegisteredError
		if ok := errors.As(err, &are); ok {
			ev, ok = are.ExistingCollector.(prometheus.Counter)
			if !ok {
				panic("different metric type registration")
			}
		} else {
			panic(err)
		}
	}

	return ev
}

func NewMetrics() Metrics {
	return Metrics{
		ConnectionErrors: newCounter(prometheus.CounterOpts{
			Namespace:   Namespace,
			Subsystem:   Subsystem,
			Name:        "connection_errors_total",
			ConstLabels: nil,
		}),
		SuccessfulSub: newCounter(prometheus.CounterOpts{
			Namespace:   Namespace,
			Subsystem:   Subsystem,
			Name:        "successful_subscribes",
			ConstLabels: nil,
		}),
		FailedSub: newCounter(prometheus.CounterOpts{
			Namespace:   Namespace,
			Subsystem:   Subsystem,
			Name:        "failed_subscribes",
			ConstLabels: nil,
		}),
		SuccessfulPut: newCounter(prometheus.CounterOpts{
			Namespace:   Namespace,
			Subsystem:   Subsystem,
			Name:        "successful_publications",
			ConstLabels: nil,
		}),
		FailedPut: newCounter(prometheus.CounterOpts{
			Namespace:   Namespace,
			Subsystem:   Subsystem,
			Name:        "failed_publications",
			ConstLabels: nil,
		}),
	}
}
