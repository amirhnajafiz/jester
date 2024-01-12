package http

import (
	"bytes"
	"errors"
	"net/http"

	"github.com/amirhnajafiz/jester/pkg"
)

type Client struct {
	Host string
}

func (c Client) SendPost(body []byte) error {
	req, err := http.NewRequest(http.MethodPost, c.Host, bytes.NewReader(body))
	if err != nil {
		return pkg.WrapError("http-client", "failed to create post request", err)
	}

	client := &http.Client{}
	if resp, er := client.Do(req); er != nil {
		return pkg.WrapError("http-client", "failed to send post request", er)
	} else {
		if resp.StatusCode != http.StatusOK {
			return pkg.WrapError("http-client", "server response", errors.New("server response was not OK"))
		}
	}

	return nil
}
