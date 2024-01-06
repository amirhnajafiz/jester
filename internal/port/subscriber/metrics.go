package subscriber

type Metrics interface {
	IncreaseConsumers(topic string)
	DecreaseConsumers(topic string)
	AddConsume(topic string, msg string)
	AddFailure(topic string)
	AddConnectionFailure()
	AddRetry(number int)
}
