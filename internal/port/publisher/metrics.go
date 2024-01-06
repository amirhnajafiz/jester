package publisher

type Metrics interface {
	IncreasePublishers(topic string)
	DecreasePublishers(topic string)
	AddPublish(topic string, msg string)
	AddFailure(topic string)
	AddConnectionFailure()
	AddRetry(number int)
}
