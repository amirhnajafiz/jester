package metrics

import "github.com/prometheus/client_golang/prometheus"

type Metrics struct {
	NumberOfSubscribers *prometheus.GaugeVec     // per topic gauge
	NumberOfPublishers  *prometheus.GaugeVec     // per topic gauge
	NumberOfPublish     *prometheus.CounterVec   // per topic counter
	NumberOfConsume     *prometheus.CounterVec   // per topic counter
	NumberOfFailures    *prometheus.CounterVec   // per topic counter
	FailedConnections   prometheus.Counter       // counter
	RetryPerConnection  prometheus.Counter       // counter
	Latency             *prometheus.HistogramVec // per topic histogram
}

func (m Metrics) UpdateNumberOfSubscribers(topic string, value int) {
	m.NumberOfSubscribers.With(prometheus.Labels{"topic": topic}).Add(float64(value))
}

func (m Metrics) UpdateNumberOfPublishers(topic string, value int) {
	m.NumberOfPublishers.With(prometheus.Labels{"topic": topic}).Add(float64(value))
}

func (m Metrics) IncreaseNumberOfPublish(topic string) {
	m.NumberOfPublish.With(prometheus.Labels{"topic": topic}).Inc()
}

func (m Metrics) IncreaseNumberOfConsume(topic string) {
	m.NumberOfConsume.With(prometheus.Labels{"topic": topic}).Inc()
}

func (m Metrics) IncreaseNumberOfFailures(topic string) {
	m.NumberOfFailures.With(prometheus.Labels{"topic": topic}).Inc()
}

func (m Metrics) IncreaseFailedConnections() {
	m.FailedConnections.Inc()
}

func (m Metrics) IncreaseRetryPerConnection(value float64) {
	m.RetryPerConnection.Add(value)
}

func (m Metrics) AddLatency(topic string, value float64) {
	m.Latency.With(prometheus.Labels{"topic": topic}).Observe(value)
}
