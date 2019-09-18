package lib

import (
	"encoding/json"

	"github.com/tendermint/tendermint/crypto"

	sdk "github.com/cosmos/cosmos-sdk/types"
	cst "github.com/cosmos/cosmos-sdk/x/staking/types"
)


// MsgCreateValidator - struct for bonding transactions
type MsgCreateValidator struct {
	cst.MsgCreateValidator
	ValidatorType ValidatorType `json:"validator_type"`
}

type msgCreateValidatorJSON struct {
	Description       cst.Description   `json:"description"`
	Commission        cst.CommissionMsg `json:"commission"`
	MinSelfDelegation sdk.Int           `json:"min_self_delegation"`
	DelegatorAddress  sdk.AccAddress    `json:"delegator_address"`
	ValidatorAddress  sdk.ValAddress    `json:"validator_address"`
	PubKey            string            `json:"pubkey"`
	Value             sdk.Coin          `json:"value"`
	ValidatorType     ValidatorType     `json:"validator_type"`
}



// MsgEditValidator - struct for editing a validator
type MsgEditValidator struct {
	cst.MsgEditValidator
	ValidatorType ValidatorType `json:"validator_type"`
}

type msgEditValidatorJSON struct {
	cst.Description
	ValidatorAddress  sdk.ValAddress `json:"address"`
	CommissionRate    *sdk.Dec       `json:"commission_rate"`
	MinSelfDelegation *sdk.Int       `json:"min_self_delegation"`
	ValidatorType     ValidatorType  `json:"validator_type"`
}

// MsgDelegate - struct for bonding transactions
type MsgDelegate struct {
	cst.MsgDelegate

	ValidatorType ValidatorType `json:"validator_type"`
}

type msgDelegateJSON struct {
	DelegatorAddress sdk.AccAddress `json:"delegator_address"`
	ValidatorAddress sdk.ValAddress `json:"validator_address"`
	Amount           sdk.Coin       `json:"amount"`
	ValidatorType    ValidatorType  `json:"validator_type"`
}

// MsgUndelegate - struct for unbonding transactions
type MsgUndelegate struct {
	cst.MsgUndelegate

	ValidatorType ValidatorType `json:"validator_type"`
}

type msgUndelegateJSON struct {
	DelegatorAddress sdk.AccAddress `json:"delegator_address"`
	ValidatorAddress sdk.ValAddress `json:"validator_address"`
	Amount           sdk.Coin       `json:"amount"`
	ValidatorType    ValidatorType  `json:"validator_type"`
}


// MsgDelegate - struct for bonding transactions
type MsgBeginRedelegate struct {
	cst.MsgBeginRedelegate

	ValidatorType ValidatorType `jason:"validator_type"`
}

type msgBeginRedelegateJSON struct {
	DelegatorAddress    sdk.AccAddress `json:"delegator_address"`
	ValidatorSrcAddress sdk.ValAddress `json:"validator_src_address"`
	ValidatorDstAddress sdk.ValAddress `json:"validator_dst_address"`
	Amount              sdk.Coin       `json:"amount"`
	ValidatorType       ValidatorType  `json:"validator_type"`
}
