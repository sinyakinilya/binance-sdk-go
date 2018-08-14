package main

import (
	"fmt"
	sdk "github.com/sinyakinilya/binance-sdk-go"
	listener "github.com/sinyakinilya/binance-sdk-go/ws-listener"
	"os"
	"os/signal"
	"time"
)

func main() {

	symbol := sdk.NewSymbol("BTC", "USDT")

	depth := fmt.Sprintf(sdk.DiffDepthStream, symbol.ToLower())
	trade := fmt.Sprintf(sdk.TradeStream, symbol.ToLower())
	channels := make(sdk.Channels)

	channels[depth] = listener.Depth()
	channels[trade] = listener.Trades()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	done, err := sdk.OpenWebSocketStreams(channels)
	if err != nil {
		panic(err)
	}

Loop:
	for {
		select {
		case <-interrupt:
			done <- struct{}{}
			break Loop
		case <-done:
			break Loop
		}
	}

	time.Sleep(time.Second)
}
