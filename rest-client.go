package binance_sdk_go

type BinanceRESTInterface interface {
	CreateListenKeyForUserDataStream() (listenKey string, err error)
	UpdateListenKeyForUserDataStream(listenKey string) error
	DeleteListenKeyForUserDataStream(listenKey string) error
}
