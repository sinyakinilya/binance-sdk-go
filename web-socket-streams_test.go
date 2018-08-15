package binance_sdk_go

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestCreateStreamsParams(t *testing.T) {

	channels := make(Events)

	channels["sn1"] = make(chan *json.RawMessage)
	channels["sn2"] = make(chan *json.RawMessage)
	channels["sn3"] = make(chan *json.RawMessage)
	params := createStreamsParams(&channels)

	assert.NotEqual(t, -1, strings.Index(params, "sn1"))
	assert.NotEqual(t, -1, strings.Index(params, "sn2"))
	assert.NotEqual(t, -1, strings.Index(params, "sn3"))
}
