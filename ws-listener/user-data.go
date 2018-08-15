package ws_listener

import (
	"fmt"
	"time"

	"encoding/json"
	"github.com/sinyakinilya/binance-sdk-go/rest"
)

func UserData(listenKey string, client *rest.Client) (ch chan *json.RawMessage) {

	ch = make(chan *json.RawMessage)
	go func() {
		ticker := time.NewTicker(time.Minute * 59)

		defer client.DeleteListenKeyForUserDataStream(listenKey)
		defer ticker.Stop()
		for {
			select {
			case t := <-ticker.C:
				client.UpdateListenKeyForUserDataStream(listenKey)
				fmt.Printf("[%s][PUT on listenKey]: validity is extended\n", t.String())
			case a := <-ch:
				fmt.Println(a)
				/*				i := 0
								for _, asset := range a.Balances {
									switch asset.Asset {
									case "BTC":
										b.Balance["BTC"] = asset
										i++
									case "USDT":
										b.Balance["USDT"] = asset
										i++
									}
									if i == 2 {
										msg := fmt.Sprintf("%s] Баланс BTC %f  USDT %f", time.Now().Format("2006-01-02 15:04:05"), b.Balance["BTC"].Free, b.Balance["USDT"].Free)
										fmt.Println(msg)
										b.SendToTelegram(msg)
										break
									}
								}
				*/
			}
		}
	}()

	return ch
}
