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
