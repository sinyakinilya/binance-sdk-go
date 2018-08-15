package binance_sdk_go

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"strings"
	"time"
)

/*

 - Aggregate Trade Streams
 - Trade Streams
 - Kline/Candlestick Streams
 - Individual Symbol Mini Ticker Stream
 - All Market Mini Tickers Stream
 - Individual Symbol Ticker Streams
 - All Market Tickers Stream
 - Partial Book Depth Streams
 - Diff. Depth Stream

*/

const (
	//	AggregateTradeStream string = "%s@aggtrade"
	TradeStream string = "%s@trade"
	//	PartialDepthStream   string = "%s@depth%d"
	DiffDepthStream string = "%s@depth"
)

type Events map[string]chan *json.RawMessage

type WsEvent struct {
	Stream string          `json:"stream"`
	Data   json.RawMessage `json:"data"`
}

func createStreamsParams(channels *Events) string {
	streamNames := make([]string, len(*channels))
	i := 0
	for streamName := range *channels {
		streamNames[i] = streamName
		i++
	}

	return strings.Join(streamNames, "/")
}

func OpenWebSocketStreams(eventsName Events) (chan struct{}, error) {

	url := fmt.Sprintf("wss://stream.binance.com:9443/stream?streams=%s", createStreamsParams(&eventsName))
	c, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return nil, err
	}

	done := make(chan struct{})

	go func() {
		defer close(done)

		for {
			select {
			default:
				_, message, err := c.ReadMessage()
				if err != nil {
					fmt.Println("wsRead", err)
					return
				}
				var eventData WsEvent
				if err := json.Unmarshal(message, &eventData); err != nil {
					fmt.Println("wsRead", err)
					return
				}
				eventsName[eventData.Stream] <- &eventData.Data
			}
		}
	}()

	go checkerConnection(c, done)

	return done, nil
}

func checkerConnection(c *websocket.Conn, done chan struct{}) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	defer c.Close()

	for {
		select {
		case <-ticker.C:
			err := c.WriteMessage(websocket.PingMessage, []byte{})
			if err != nil {
				fmt.Println("wsWrite", err)
				return
			}
		case <-done:
			fmt.Println("closing connection")
			return
		}
	}
}
