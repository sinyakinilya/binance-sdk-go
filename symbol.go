package binance_sdk_go

import "strings"

type Symbol struct {
	BaseAsset  string
	QuoteAsset string
}

func (s *Symbol) ToString() string {
	return s.BaseAsset + s.QuoteAsset
}

func (s *Symbol) ToLower() string {
	return strings.ToLower(s.ToString())
}

func NewSymbol(base, quote string) *Symbol {
	return &Symbol{
		BaseAsset:  base,
		QuoteAsset: quote,
	}
}
