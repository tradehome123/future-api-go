package private

import (
	"encoding/json"
	"fmt"
	"zb_sdk_go/types"

	. "zb_sdk_go/log"
)

func BalanceHandler(msg []byte) bool {
	var data []types.FundBalanceResult
	err := json.Unmarshal(msg, &data)
	if err != nil {
		Log.Error().Str("data", string(msg)).Msg(err.Error())
		return true
	}
	Log.Info().Str("data", fmt.Sprint(data)).Msg("BalanceHandler")
	return false
}
