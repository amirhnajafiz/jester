package metrics

type Metrics struct {
	NumberOfSubscribers int // per topic gauge
	NumberOfPublishers  int // per topic gauge
	NumberOfPublish     int // per topic counter
	NumberOfConsume     int // per topic counter
	NumberOfFailures    int // per topic counter
	FailedConnections   int // counter
	RetryPerConnection  int // counter
	Latency             int // per topic histogram
}

func (m Metrics) IncreaseNumberOfSubscribers(topic string) {
	m.NumberOfSubscribers++
}

func (m Metrics) IncreaseNumberOfPublishers(topic string) {
	m.NumberOfPublishers++
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

func (m Metrics) AddLatency(topic string) {

}
