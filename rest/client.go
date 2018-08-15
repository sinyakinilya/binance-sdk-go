package rest

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	sdk "github.com/sinyakinilya/binance-sdk-go"
	"net/http"
)

type Client struct {
	URL    string
	APIKey string
	Signer sdk.Signer
}

func (c *Client) request(method string, endpoint string, params map[string]string,
	apiKey bool, sign bool) (*http.Response, error) {
	transport := &http.Transport{}
	client := &http.Client{
		Transport: transport,
	}

	url := fmt.Sprintf("%s/%s", c.URL, endpoint)
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create request")
	}

	q := req.URL.Query()
	for key, val := range params {
		q.Add(key, val)
	}
	if apiKey {
		req.Header.Add("X-MBX-APIKEY", c.APIKey)
	}
	if sign {
		q.Add("signature", c.Signer.Sign([]byte(q.Encode())))
	}
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type BinanceErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

func (e BinanceErrorResponse) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Message)
}

func (c *Client) handleError(textRes []byte) error {
	err := &BinanceErrorResponse{}

	if err := json.Unmarshal(textRes, err); err != nil {
		return errors.Wrap(err, "error unmarshal failed")
	}
	return err
}
