package metrics

import "github.com/prometheus/client_golang/prometheus"

type Metrics struct {
	NumberOfSubscribers *prometheus.GaugeVec   // per topic gauge
	NumberOfPublishers  *prometheus.GaugeVec   // per topic gauge
	NumberOfPublish     *prometheus.CounterVec // per topic counter
	NumberOfConsume     *prometheus.CounterVec // per topic counter
	NumberOfFailures    *prometheus.CounterVec // per topic counter
	FailedConnections   prometheus.Counter     // counter
	RetryPerConnection  prometheus.Counter     // counter
	Latency             prometheus.Histogram   // per topic histogram
}

func (m Metrics) UpdateNumberOfSubscribers(topic string, value int) {
}

func (m Metrics) UpdateNumberOfPublishers(topic string, value int) {
}

func (m Metrics) IncreaseNumberOfPublish(topic string) {

}

func (m Metrics) IncreaseNumberOfConsume(topic string) {

}

func (m Metrics) IncreaseNumberOfFailures(topic string) {

}

func (m Metrics) IncreaseFailedConnections() {

}

func (m Metrics) IncreaseRetryPerConnection() {

}

func (m Metrics) AddLatency(topic string, value float64) {

}
