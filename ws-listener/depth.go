package ws_listener

import (
	"encoding/json"
	"fmt"
)

type DepthEvent struct {
	Type          string          `json:"e"`
	Time          float64         `json:"E"`
	Symbol        string          `json:"s"`
	UpdateID      int             `json:"u"`
	BidDepthDelta [][]interface{} `json:"b"`
	AskDepthDelta [][]interface{} `json:"a"`
}

func Depth() chan *json.RawMessage {

	ch := make(chan *json.RawMessage)

	go func() {
		for {
			select {
			case data := <-ch:
				var depth DepthEvent
				if err := json.Unmarshal(*data, &depth); err != nil {
					fmt.Println("Listener for DepthEvent", err)
					return
				}
				fmt.Printf("%#v\n", depth)
			}
		}
	}()

	return ch
}
