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
			Name:      "number_of_subscribers",
			Help:      "total number of subscribers per topic",
		}, []string{"topic"}),
		numberOfPublishers: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: cfg.Namespace,
			Subsystem: cfg.Subsystem,
			Name:      "number_of_publishers",
			Help:      "total number of publishers per topic",
		}, []string{"topic"}),
		numberOfPublish: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: cfg.Namespace,
			Subsystem: cfg.Subsystem,
			Name:      "number_of_publish",
			Help:      "total number of publish events per topic",
		}, []string{"topic"}),
		numberOfConsume: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: cfg.Namespace,
			Subsystem: cfg.Subsystem,
			Name:      "number_of_consume",
			Help:      "total number of consume events per topic",
		}, []string{"topic"}),
		numberOfFailures: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: cfg.Namespace,
			Subsystem: cfg.Subsystem,
			Name:      "number_of_failures",
			Help:      "total number of failed publish events per topic",
		}, []string{"topic"}),
		failedConnections: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: cfg.Namespace,
			Subsystem: cfg.Subsystem,
			Name:      "failed_connections",
			Help:      "total number of failed connections",
		}),
		retryPerConnection: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: cfg.Namespace,
			Subsystem: cfg.Subsystem,
			Name:      "retry_connections",
			Help:      "total number of retry per connection",
		}),
		latency: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: cfg.Namespace,
			Subsystem: cfg.Subsystem,
			Name:      "latency",
			Help:      "service latency per topic",
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
