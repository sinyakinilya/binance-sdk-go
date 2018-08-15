package rest

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
)

type CreateListenKeyResponse struct {
	ListenKey string
}

func (c *Client) CreateListenKeyForUserDataStream() (listenKey string, err error) {
	params := make(map[string]string)

	res, err := c.request("POST", "api/v1/userDataStream", params, true, false)
	if err != nil {
		return listenKey, err
	}
	textRes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return listenKey, errors.Wrap(err, "unable to read response from userDataStream.post")
	}
	defer res.Body.Close()

	fmt.Println(string(textRes))
	if res.StatusCode != 200 {
		return listenKey, c.handleError(textRes)
	}

	var s CreateListenKeyResponse
	if err := json.Unmarshal(textRes, &s); err != nil {
		return listenKey, errors.Wrap(err, "stream unmarshal failed")
	}
	return s.ListenKey, nil
}

func (c *Client) UpdateListenKeyForUserDataStream(listenKey string) error {
	params := make(map[string]string)
	params["listenKey"] = listenKey

	res, err := c.request("PUT", "api/v1/userDataStream", params, true, false)
	if err != nil {
		return err
	}
	textRes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return errors.Wrap(err, "unable to read response from userDataStream.put")
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return c.handleError(textRes)
	}
	return nil
}
func (c *Client) DeleteListenKeyForUserDataStream(listenKey string) error {
	params := make(map[string]string)
	params["listenKey"] = listenKey

	res, err := c.request("DELETE", "api/v1/userDataStream", params, true, false)
	if err != nil {
		return err
	}
	textRes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return errors.Wrap(err, "unable to read response from userDataStream.delete")
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return c.handleError(textRes)
	}
	return nil
}
