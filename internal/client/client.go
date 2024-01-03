package client

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
)

type Client struct {
	Host string
}

func (c Client) SendPost(body []byte) error {
	req, err := http.NewRequest(http.MethodPost, c.Host, bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("failed to create post request: %w", err)
	}

	client := &http.Client{}
	if resp, er := client.Do(req); er != nil {
		return fmt.Errorf("failed to send http request: %w", err)
	} else {
		if resp.StatusCode != http.StatusOK {
			return errors.New("server response was not OK")
		}
	}

	return nil
}
