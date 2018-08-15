package ws_listener

import (
	"encoding/json"
	"fmt"
)

type TradeEvent struct {
	Type      string `json:"e"`
	EventTime int64  `json:"E"`
	Symbol    string `json:"s"`

	TradeID       uint64  `json:"t"`
	Price         float64 `json:"p,string"`
	Quantity      float64 `json:"q,string"`
	BuyerId       uint64  `json:"b"`
	SellerId      uint64  `json:"a"`
	TradeTime     int64   `json:"T"`
	IsMarketMaker bool    `json:"m"`
}

func Trades() chan *json.RawMessage {
	ch := make(chan *json.RawMessage)

	go func() {
		defer close(ch)

		for {
			select {
			case data := <-ch:
				var trade TradeEvent
				if err := json.Unmarshal(*data, &trade); err != nil {
					fmt.Println("Listener for TradeEvent", err)
					return
				}
				fmt.Println(trade)
			}
		}
	}()

	return ch
}
