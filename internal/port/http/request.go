package http

// Request is being sent by publisher and subscriber
// in order to update metrics.
type Request struct {
	Field int     `json:"field"`
	Label string  `json:"label"`
	Param string  `json:"param"`
	Value float64 `json:"value"`
}
