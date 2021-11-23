package public

import (
	"encoding/json"
	"strconv"
	. "zb_sdk_go/log"
)

func MarkIndexPriceWSHandler(msg []byte) bool {
	var data string
	err := json.Unmarshal(msg, &data)
	if err != nil {
		Log.Error().Str("data", string(msg)).Msg(err.Error())
		return true
	}

	price, err := strconv.ParseFloat(data, 10)
	if err != nil {
		Log.Error().Str("data", string(msg)).Msg(err.Error())
		return true
	}
	Log.Info().Float64("price", price).Msg("MarkIndexPriceWSHandler")

	return false
}
