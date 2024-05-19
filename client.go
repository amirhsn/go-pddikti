package gopddikti

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	clientConfig ClientConfig
}

// InitClient is used to init the PDDikti SDK Client
func InitClient(cfg *ClientConfig) (*Client, error) {
	var (
		config ClientConfig
		err    error
	)

	if cfg == nil {
		config = DefaultConfig()
	} else {
		config, err = CustomConfig(cfg)
	}

	go func() {
		initRegex()
	}()

	return &Client{
		clientConfig: config,
	}, err
}

func (c *Client) createRequest(method, path string) (*http.Request, error) {
	urlPath := c.clientConfig.baseURL + path

	req, err := http.NewRequest(method, urlPath, nil)
	if err != nil {
		return req, err
	}

	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	return req, err
}

func (c *Client) doRequest(req *http.Request, v any) error {
	res, err := c.clientConfig.httpClient.Do(req)
	if err != nil {
		return err
	}
	if res != nil {
		defer res.Body.Close()
	} else {
		return errors.New("response is nil")
	}

	if res.StatusCode != http.StatusOK {
		return handleSpecialStatusCodeResponse(res.StatusCode)
	}

	return doDecode(res, v)
}

func doDecode(res *http.Response, v any) error {
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil
	}

	return json.Unmarshal(b, v)
}

func handleSpecialStatusCodeResponse(code int) error {
	// 403 error code is the most common error code returned by PDDikti
	// it is caused by expired certificate, but sometimes will be fixed in several days
	if code == http.StatusForbidden {
		return fmt.Errorf("%d: forbidden, could be caused by expired certificate", code)
	}
	return fmt.Errorf("%d: error happened", code)
}
