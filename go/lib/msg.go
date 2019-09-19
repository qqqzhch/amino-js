package lib

import (
	sdk "github.com/cosmos/amino-js/go/lib/cosmos/cosmos-sdk/types"
	
)

var (
	
	_ sdk.Msg = &MsgDelegate{}
	_ sdk.Msg = &MsgUndelegate{}
	_ sdk.Msg = &MsgBeginRedelegate{}
)





type MsgDelegate struct {
	DelegatorAddress sdk.AccAddress `json:"delegator_address"`
	ValidatorAddress sdk.ValAddress `json:"validator_address"`
	Amount           sdk.Coin       `json:"amount"`
	ValidatorType    ValidatorType   `json:"validator_type"`
}

type MsgBeginRedelegate struct {
	DelegatorAddress    sdk.AccAddress `json:"delegator_address"`
	ValidatorSrcAddress sdk.ValAddress `json:"validator_src_address"`
	ValidatorDstAddress sdk.ValAddress `json:"validator_dst_address"`
	Amount              sdk.Coin       `json:"amount"`
}

type MsgUndelegate struct {
	DelegatorAddress sdk.AccAddress `json:"delegator_address"`
	ValidatorAddress sdk.ValAddress `json:"validator_address"`
	Amount           sdk.Coin       `json:"amount"`
}
