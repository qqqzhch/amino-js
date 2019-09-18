package lib

import (
	"encoding/json"

	"github.com/tendermint/tendermint/crypto"

	sdk "github.com/cosmos/cosmos-sdk/types"
	cst "github.com/cosmos/cosmos-sdk/x/staking/types"
)

var (
	_ sdk.Msg = &MsgCreateValidator{}
	_ sdk.Msg = &MsgEditValidator{}
	_ sdk.Msg = &MsgDelegate{}
	_ sdk.Msg = &MsgUndelegate{}
	_ sdk.Msg = &MsgBeginRedelegate{}
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

// Default way to create validator. Delegator address and validator address are the same
func NewMsgCreateValidator(
	valAddr sdk.ValAddress, pubKey crypto.PubKey, selfDelegation sdk.Coin,
	description cst.Description, commission cst.CommissionMsg, minSelfDelegation sdk.Int, validatorType ValidatorType,
) MsgCreateValidator {

	return MsgCreateValidator{
		MsgCreateValidator: cst.NewMsgCreateValidator(valAddr, pubKey, selfDelegation, description, commission, minSelfDelegation),
		ValidatorType:      validatorType,
	}
}

//nolint
func (msg MsgCreateValidator) Route() string { return msg.MsgCreateValidator.Route() }
func (msg MsgCreateValidator) Type() string  { return msg.MsgCreateValidator.Type() }

// Return address(es) that must sign over msg.GetSignBytes()
func (msg MsgCreateValidator) GetSigners() []sdk.AccAddress {
	return msg.MsgCreateValidator.GetSigners()
}

// MarshalJSON implements the json.Marshaler interface to provide custom JSON
// serialization of the MsgCreateValidator type.
func (msg MsgCreateValidator) MarshalJSON() ([]byte, error) {
	return json.Marshal(msgCreateValidatorJSON{
		Description:       msg.Description,
		Commission:        msg.Commission,
		DelegatorAddress:  msg.DelegatorAddress,
		ValidatorAddress:  msg.ValidatorAddress,
		PubKey:            sdk.MustBech32ifyConsPub(msg.PubKey),
		Value:             msg.Value,
		MinSelfDelegation: msg.MinSelfDelegation,
		ValidatorType:     msg.ValidatorType,
	})
}

// UnmarshalJSON implements the json.Unmarshaler interface to provide custom
// JSON deserialization of the MsgCreateValidator type.
func (msg *MsgCreateValidator) UnmarshalJSON(bz []byte) error {
	var msgCreateValJSON msgCreateValidatorJSON
	if err := json.Unmarshal(bz, &msgCreateValJSON); err != nil {
		return err
	}

	msg.Description = msgCreateValJSON.Description
	msg.Commission = msgCreateValJSON.Commission
	msg.DelegatorAddress = msgCreateValJSON.DelegatorAddress
	msg.ValidatorAddress = msgCreateValJSON.ValidatorAddress
	var err error
	msg.PubKey, err = sdk.GetConsPubKeyBech32(msgCreateValJSON.PubKey)
	if err != nil {
		return err
	}
	msg.Value = msgCreateValJSON.Value
	msg.MinSelfDelegation = msgCreateValJSON.MinSelfDelegation
	msg.ValidatorType = msgCreateValJSON.ValidatorType

	return nil
}

// GetSignBytes returns the message bytes to sign over.
func (msg MsgCreateValidator) GetSignBytes() []byte {
	bz := MsgCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// quick validity check
func (msg MsgCreateValidator) ValidateBasic() sdk.Error {
	if err := msg.MsgCreateValidator.ValidateBasic(); err != nil {
		return err
	}

	if msg.ValidatorType != PartnerValidatorType && msg.ValidatorType != ConsensusValidatorType {
		return ErrBadValidatorType(cst.DefaultCodespace)
	}

	return nil
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

func NewMsgEditValidator(valAddr sdk.ValAddress, description cst.Description,
	newRate *sdk.Dec, newMinSelfDelegation *sdk.Int, validatorType ValidatorType) MsgEditValidator {
	return MsgEditValidator{
		MsgEditValidator: cst.NewMsgEditValidator(valAddr, description, newRate, newMinSelfDelegation),
		ValidatorType:    validatorType,
	}
}

//nolint
func (msg MsgEditValidator) Route() string { return msg.MsgEditValidator.Route() }
func (msg MsgEditValidator) Type() string  { return msg.MsgEditValidator.Type() }
func (msg MsgEditValidator) GetSigners() []sdk.AccAddress {
	return msg.MsgEditValidator.GetSigners()
}

// get the bytes for the message signer to sign on
func (msg MsgEditValidator) GetSignBytes() []byte {
	bz := MsgCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgEditValidator) MarshalJSON() ([]byte, error) {
	return json.Marshal(msgEditValidatorJSON{
		Description:       msg.Description,
		ValidatorAddress:  msg.ValidatorAddress,
		CommissionRate:    msg.CommissionRate,
		MinSelfDelegation: msg.MinSelfDelegation,
		ValidatorType:     msg.ValidatorType,
	})
}

func (msg *MsgEditValidator) UnmarshalJSON(bz []byte) error {
	var msgEditValidatorJSON msgEditValidatorJSON
	if err := json.Unmarshal(bz, &msgEditValidatorJSON); err != nil {
		return err
	}

	msg.Description = msgEditValidatorJSON.Description
	msg.ValidatorAddress = msgEditValidatorJSON.ValidatorAddress
	msg.CommissionRate = msgEditValidatorJSON.CommissionRate
	msg.MinSelfDelegation = msgEditValidatorJSON.MinSelfDelegation
	msg.ValidatorType = msgEditValidatorJSON.ValidatorType
	return nil
}

// quick validity check
func (msg MsgEditValidator) ValidateBasic() sdk.Error {
	return msg.MsgEditValidator.ValidateBasic()
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

func NewMsgDelegate(delAddr sdk.AccAddress, valAddr sdk.ValAddress, amount sdk.Coin, validatorType ValidatorType) MsgDelegate {
	return MsgDelegate{
		MsgDelegate:   cst.NewMsgDelegate(delAddr, valAddr, amount),
		ValidatorType: validatorType,
	}
}

//nolint
func (msg MsgDelegate) Route() string { return msg.MsgDelegate.Route() }
func (msg MsgDelegate) Type() string  { return msg.MsgDelegate.Type() }
func (msg MsgDelegate) GetSigners() []sdk.AccAddress {
	return msg.MsgDelegate.GetSigners()
}

// get the bytes for the message signer to sign on
func (msg MsgDelegate) GetSignBytes() []byte {
	bz := MsgCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgDelegate) MarshalJSON() ([]byte, error) {
	return json.Marshal(msgDelegateJSON{
		DelegatorAddress: msg.DelegatorAddress,
		ValidatorAddress: msg.ValidatorAddress,
		Amount:           msg.Amount,
		ValidatorType:    msg.ValidatorType,
	})
}

func (msg *MsgDelegate) UnmarshalJSON(bz []byte) error {
	var msgDelegateJSON msgDelegateJSON
	if err := json.Unmarshal(bz, &msgDelegateJSON); err != nil {
		return err
	}

	msg.DelegatorAddress = msgDelegateJSON.DelegatorAddress
	msg.ValidatorAddress = msgDelegateJSON.ValidatorAddress
	msg.Amount = msgDelegateJSON.Amount
	msg.ValidatorType = msgDelegateJSON.ValidatorType
	return nil
}

// quick validity check
func (msg MsgDelegate) ValidateBasic() sdk.Error {
	return msg.MsgDelegate.ValidateBasic()
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

func NewMsgUndelegate(delAddr sdk.AccAddress, valAddr sdk.ValAddress, amount sdk.Coin, validatorType ValidatorType) MsgUndelegate {
	return MsgUndelegate{
		MsgUndelegate: cst.NewMsgUndelegate(delAddr, valAddr, amount),
		ValidatorType: validatorType,
	}
}

//nolint
func (msg MsgUndelegate) Route() string                { return msg.MsgUndelegate.Route() }
func (msg MsgUndelegate) Type() string                 { return msg.MsgUndelegate.Type() }
func (msg MsgUndelegate) GetSigners() []sdk.AccAddress { return msg.MsgUndelegate.GetSigners() }

// get the bytes for the message signer to sign on
func (msg MsgUndelegate) GetSignBytes() []byte {
	bz := MsgCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgUndelegate) MarshalJSON() ([]byte, error) {
	return json.Marshal(msgUndelegateJSON{
		DelegatorAddress: msg.DelegatorAddress,
		ValidatorAddress: msg.ValidatorAddress,
		Amount:           msg.Amount,
		ValidatorType:    msg.ValidatorType,
	})
}

func (msg *MsgUndelegate) UnmarshalJSON(bz []byte) error {
	var msgUndelegateJSON msgUndelegateJSON
	if err := json.Unmarshal(bz, &msgUndelegateJSON); err != nil {
		return err
	}

	msg.DelegatorAddress = msgUndelegateJSON.DelegatorAddress
	msg.ValidatorAddress = msgUndelegateJSON.ValidatorAddress
	msg.Amount = msgUndelegateJSON.Amount
	msg.ValidatorType = msgUndelegateJSON.ValidatorType
	return nil
}

// quick validity check
func (msg MsgUndelegate) ValidateBasic() sdk.Error {
	return msg.MsgUndelegate.ValidateBasic()
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

func NewMsgBeginRedelegate(delAddr sdk.AccAddress, valSrcAddr,
	valDstAddr sdk.ValAddress, amount sdk.Coin, validatorType ValidatorType) MsgBeginRedelegate {

	return MsgBeginRedelegate{
		MsgBeginRedelegate: cst.NewMsgBeginRedelegate(delAddr, valSrcAddr, valDstAddr, amount),
		ValidatorType:      validatorType,
	}
}

//nolint
func (msg MsgBeginRedelegate) Route() string { return msg.MsgBeginRedelegate.Route() }
func (msg MsgBeginRedelegate) Type() string  { return msg.MsgBeginRedelegate.Type() }
func (msg MsgBeginRedelegate) GetSigners() []sdk.AccAddress {
	return msg.MsgBeginRedelegate.GetSigners()
}

// get the bytes for the message signer to sign on
func (msg MsgBeginRedelegate) GetSignBytes() []byte {
	bz := MsgCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgBeginRedelegate) MarshalJSON() ([]byte, error) {
	return json.Marshal(msgBeginRedelegateJSON{
		DelegatorAddress:    msg.DelegatorAddress,
		ValidatorSrcAddress: msg.ValidatorSrcAddress,
		ValidatorDstAddress: msg.ValidatorDstAddress,
		Amount:              msg.Amount,
		ValidatorType:       msg.ValidatorType,
	})
}

func (msg *MsgBeginRedelegate) UnmarshalJSON(bz []byte) error {
	var msgBeginRedelegateJSON msgBeginRedelegateJSON
	if err := json.Unmarshal(bz, &msgBeginRedelegateJSON); err != nil {
		return err
	}

	msg.DelegatorAddress = msgBeginRedelegateJSON.DelegatorAddress
	msg.ValidatorSrcAddress = msgBeginRedelegateJSON.ValidatorSrcAddress
	msg.ValidatorDstAddress = msgBeginRedelegateJSON.ValidatorDstAddress
	msg.Amount = msgBeginRedelegateJSON.Amount
	msg.ValidatorType = msgBeginRedelegateJSON.ValidatorType
	return nil
}

// quick validity check
func (msg MsgBeginRedelegate) ValidateBasic() sdk.Error {
	return msg.MsgBeginRedelegate.ValidateBasic()
}
