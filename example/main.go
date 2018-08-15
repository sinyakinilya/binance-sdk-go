package main

import (
	"fmt"
	sdk "github.com/sinyakinilya/binance-sdk-go"
	"github.com/sinyakinilya/binance-sdk-go/rest"
	listener "github.com/sinyakinilya/binance-sdk-go/ws-listener"
	"os"
	"os/signal"
	"time"
)

func main() {

	hmacSigner := &sdk.HmacSigner{
		Key: []byte(os.Getenv("B_SECRET")),
	}

	client := rest.Client{
		URL:    "https://www.binance.com",
		APIKey: os.Getenv("B_APIKEY"),
		Signer: hmacSigner,
	}

	symbol := sdk.NewSymbol("BTC", "USDT")

	depthEventName := fmt.Sprintf(sdk.DiffDepthStream, symbol.ToLower())
	tradeEventName := fmt.Sprintf(sdk.TradeStream, symbol.ToLower())

	userDataKeyEventName, err := client.CreateListenKeyForUserDataStream()
	if err != nil {
		panic(err)
	}

	channel := make(sdk.Events)

	channel[depthEventName] = listener.Depth()
	channel[tradeEventName] = listener.Trades()
	channel[userDataKeyEventName] = listener.UserData(userDataKeyEventName, &client)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	done, err := sdk.OpenWebSocketStreams(channel)
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
