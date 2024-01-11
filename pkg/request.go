package pkg

import "encoding/json"

// Request is being sent by publisher and subscriber
// in order to update metrics.
type Request struct {
	Field int     `json:"field"`
	Label string  `json:"label"`
	Param string  `json:"param"`
	Value float64 `json:"value"`
}

func (r Request) WithParam(param string) Request {
	r.Param = param

	return r
}

func (r Request) WithLabel(topic string) Request {
	r.Label = topic

	return r
}

func (r Request) WithValue(value float64) Request {
	r.Value = value

	return r
}

func (r Request) ToBytes() []byte {
	bytes, _ := json.Marshal(r)

	return bytes
}

func NewRequest(field int) Request {
	return Request{
		Field: field,
	}
}
