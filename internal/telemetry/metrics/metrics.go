package metrics

import "github.com/prometheus/client_golang/prometheus"

type Metrics struct {
	numberOfSubscribers *prometheus.GaugeVec     // per topic gauge
	numberOfPublishers  *prometheus.GaugeVec     // per topic gauge
	numberOfPublish     *prometheus.CounterVec   // per topic counter
	numberOfConsume     *prometheus.CounterVec   // per topic counter
	numberOfFailures    *prometheus.CounterVec   // per topic counter
	failedConnections   prometheus.Counter       // counter
	retryPerConnection  prometheus.Counter       // counter
	latency             *prometheus.HistogramVec // per topic histogram
}

func New(cfg Config) *Metrics {
	m := &Metrics{
		numberOfSubscribers: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: cfg.Namespace,
			Subsystem: cfg.Subsystem,
		}, []string{"topic"}),
		numberOfPublishers: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: cfg.Namespace,
			Subsystem: cfg.Subsystem,
		}, []string{"topic"}),
		numberOfPublish: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: cfg.Namespace,
			Subsystem: cfg.Subsystem,
		}, []string{"topic"}),
		numberOfConsume: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: cfg.Namespace,
			Subsystem: cfg.Subsystem,
		}, []string{"topic"}),
		numberOfFailures: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: cfg.Namespace,
			Subsystem: cfg.Subsystem,
		}, []string{"topic"}),
		failedConnections: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: cfg.Namespace,
			Subsystem: cfg.Subsystem,
		}),
		retryPerConnection: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: cfg.Namespace,
			Subsystem: cfg.Subsystem,
		}),
		latency: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: cfg.Namespace,
			Subsystem: cfg.Subsystem,
		}, []string{"topic"}),
	}

	return m
}

func (m Metrics) UpdateNumberOfSubscribers(topic string, value int) {
	m.numberOfSubscribers.With(prometheus.Labels{"topic": topic}).Add(float64(value))
}

func (m Metrics) UpdateNumberOfPublishers(topic string, value int) {
	m.numberOfPublishers.With(prometheus.Labels{"topic": topic}).Add(float64(value))
}

func (m Metrics) IncreaseNumberOfPublish(topic string) {
	m.numberOfPublish.With(prometheus.Labels{"topic": topic}).Inc()
}

func (m Metrics) IncreaseNumberOfConsume(topic string) {
	m.numberOfConsume.With(prometheus.Labels{"topic": topic}).Inc()
}

func (m Metrics) IncreaseNumberOfFailures(topic string) {
	m.numberOfFailures.With(prometheus.Labels{"topic": topic}).Inc()
}

func (m Metrics) IncreaseFailedConnections() {
	m.failedConnections.Inc()
}

func (m Metrics) IncreaseRetryPerConnection(value float64) {
	m.retryPerConnection.Add(value)
}

func (m Metrics) AddLatency(topic string, value float64) {
	m.latency.With(prometheus.Labels{"topic": topic}).Observe(value)
}
