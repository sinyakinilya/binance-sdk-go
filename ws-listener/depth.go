package ws_listener

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"strconv"
)

type OpenOffer struct {
	Price    float64
	Quantity float64
}

func createOpenOffer(data []interface{}) (*OpenOffer, error) {
	p, e := strconv.ParseFloat(data[0].(string), 64)
	if e != nil {
		return &OpenOffer{}, errors.New("data[0] - not float64")
	}
	q, e := strconv.ParseFloat(data[1].(string), 64)
	if e != nil {
		return &OpenOffer{}, errors.New("data[0] - not float64")
	}
	return &OpenOffer{
		Price:    p,
		Quantity: q,
	}, nil
}

type depthEventResponse struct {
	Type      string `json:"e"`
	EventTime int64  `json:"E"`
	Symbol    string `json:"s"`

	UpdateID      int             `json:"u"`
	BidDepthDelta [][]interface{} `json:"b"`
	AskDepthDelta [][]interface{} `json:"a"`
}

type DepthEvent struct {
	Type      string
	EventTime int64
	Symbol    string

	UpdateID      int
	BidDepthDelta []*OpenOffer
	AskDepthDelta []*OpenOffer
}

func Depth() chan *json.RawMessage {

	ch := make(chan *json.RawMessage)

	go func() {
		for {
			select {
			case data := <-ch:
				var depthResponse depthEventResponse
				if err := json.Unmarshal(*data, &depthResponse); err != nil {
					fmt.Println("Listener for DepthEvent", err)
					return
				}
				fmt.Printf("%#v\n", depthResponse)
				depth := DepthEvent{
					Type:      depthResponse.Type,
					EventTime: depthResponse.EventTime,
					Symbol:    depthResponse.Symbol,

					UpdateID: depthResponse.UpdateID,
				}
				depth.AskDepthDelta = make([]*OpenOffer, len(depthResponse.AskDepthDelta))
				for k, v := range depthResponse.AskDepthDelta {
					offer, err := createOpenOffer(v)
					if err != nil {
						fmt.Println(err)
						return
					}
					depth.AskDepthDelta[k] = offer
				}
				depth.BidDepthDelta = make([]*OpenOffer, len(depthResponse.BidDepthDelta))
				for k, v := range depthResponse.BidDepthDelta {
					offer, err := createOpenOffer(v)
					if err != nil {
						fmt.Println(err)
						return
					}
					depth.BidDepthDelta[k] = offer
				}
				fmt.Printf("%#v\n", depth)
			}
		}
	}()

	return ch
}
