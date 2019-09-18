package lib

import (
	
	cst "github.com/cosmos/cosmos-sdk/x/staking/types"
)


// MsgDelegate - struct for bonding transactions
type MsgDelegate struct {
	cst.MsgDelegate

	ValidatorType ValidatorType `json:"validator_type"`
}
