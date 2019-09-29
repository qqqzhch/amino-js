package lib

import (
	sdk "github.com/cosmos/amino-js/go/lib/cosmos/cosmos-sdk/types"
	
)

// ensure Msg interface compliance at compile time
var (
	_ sdk.Msg = &MsgAssetPledge{}
	_ sdk.Msg = &MsgAssetDrop{}
	_ sdk.Msg = &MsgCreateAsset{}
	_ sdk.Msg = &MsgMintAsset{}
	_ sdk.Msg = &MsgLockAsset{}
	_ sdk.Msg = &MsgUnLockAsset{}
	_ sdk.Msg = &MsgDestroyAsset{}
	_ sdk.Msg = &MsgRuinAsset{}
)

type MsgAssetPledge struct {
	Address sdk.AccAddress `json:"address"`
	Asset   sdk.Coin       `json:"asset"`
	Token   sdk.Coin       `json:"token"`
}

func NewMsgAssetPledge(address sdk.AccAddress, asset sdk.Coin, token sdk.Coin) MsgAssetPledge {
	return MsgAssetPledge{
		Address: address,
		Asset:   asset,
		Token:   token,
	}
}

//nolint
func (msg MsgAssetPledge) Route() string { return RouterKey }
func (msg MsgAssetPledge) Type() string  { return "assetPledge" }

// Return address(es) that must sign over msg.GetSignBytes()
func (msg MsgAssetPledge) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Address}
}

// GetSignBytes returns the message bytes to sign over.
func (msg MsgAssetPledge) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgAssetPledge) ValidateBasic() sdk.Error {
	if msg.Address.Empty() {
		return sdk.NewError(DefaultCodespace, CodeInvalidAddress, "nil validator address")
	}
	if msg.Asset.Amount.LTE(sdk.ZeroInt()) {
		return sdk.NewError(DefaultCodespace, CodeInvalidCoinAmount, "invalid asset amount %v", msg.Asset.Amount)
	}
	if msg.Token.Amount.LTE(sdk.ZeroInt()) {
		return sdk.NewError(DefaultCodespace, CodeInvalidCoinAmount, "invalid token amount %v", msg.Asset.Amount)
	}
	return nil
}

type MsgAssetDrop struct {
	Address sdk.AccAddress `json:"address"`
	Asset   sdk.Coin       `json:"asset"`
	Token   sdk.Coin       `json:"token"`
}

func NewMsgAssetDrop(address sdk.AccAddress, asset sdk.Coin, token sdk.Coin) MsgAssetDrop {
	return MsgAssetDrop{
		Address: address,
		Asset:   asset,
		Token:   token,
	}
}

//nolint
func (msg MsgAssetDrop) Route() string { return RouterKey }
func (msg MsgAssetDrop) Type() string  { return "assetDrop" }

// Return address(es) that must sign over msg.GetSignBytes()
func (msg MsgAssetDrop) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Address}
}

// GetSignBytes returns the message bytes to sign over.
func (msg MsgAssetDrop) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgAssetDrop) ValidateBasic() sdk.Error {
	if msg.Address.Empty() {
		return sdk.NewError(DefaultCodespace, CodeInvalidAddress, "nil validator address")
	}
	if msg.Asset.Amount.LTE(sdk.ZeroInt()) {
		return sdk.NewError(DefaultCodespace, CodeInvalidCoinAmount, "invalid asset amount %v", msg.Asset.Amount)
	}
	if msg.Token.Amount.LTE(sdk.ZeroInt()) {
		return sdk.NewError(DefaultCodespace, CodeInvalidCoinAmount, "invalid token amount %v", msg.Token.Amount)
	}
	return nil
}

type MsgCreateAsset struct {
	Address  sdk.AccAddress `json:"address"`
	Asset    sdk.Coin       `json:"asset"`
	Token    sdk.Coin       `json:"token"`
	Name     string         `json:"name"`
	Mintable bool           `json:"mintable"`
}

func NewMsgCreateAsset(address sdk.AccAddress, asset sdk.Coin, token sdk.Coin,
	name string, mintable bool) MsgCreateAsset {
	return MsgCreateAsset{
		Address:  address,
		Asset:    asset,
		Token:    token,
		Name:     name,
		Mintable: mintable,
	}
}

func (msg MsgCreateAsset) Route() string { return RouterKey }
func (msg MsgCreateAsset) Type() string  { return "createAsset" }
func (msg MsgCreateAsset) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Address}
}
func (msg MsgCreateAsset) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}
func (msg MsgCreateAsset) ValidateBasic() sdk.Error {
	if msg.Address.Empty() {
		return sdk.NewError(DefaultCodespace, CodeInvalidAddress, "nil address")
	}

	if msg.Asset.Amount.LTE(sdk.ZeroInt()) {
		return sdk.NewError(DefaultCodespace, CodeInvalidCoinAmount, "invalid asset amount %v", msg.Asset.Amount)
	}

	if msg.Token.Amount.LTE(sdk.ZeroInt()) {
		return sdk.NewError(DefaultCodespace, CodeInvalidCoinAmount, "invalid token amount %v", msg.Token.Amount)
	}

	return nil
}

type MsgMintAsset struct {
	Address sdk.AccAddress `json:"address"`
	Asset   sdk.Coin       `json:"asset"`
}

func NewMsgMintAsset(address sdk.AccAddress, asset sdk.Coin) MsgMintAsset {
	return MsgMintAsset{
		Address: address,
		Asset:   asset,
	}
}

func (msg MsgMintAsset) Route() string { return RouterKey }
func (msg MsgMintAsset) Type() string  { return "mintAsset" }
func (msg MsgMintAsset) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Address}
}
func (msg MsgMintAsset) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}
func (msg MsgMintAsset) ValidateBasic() sdk.Error {
	if msg.Address.Empty() {
		return sdk.NewError(DefaultCodespace, CodeInvalidAddress, "nil address")
	}

	if msg.Asset.Amount.LTE(sdk.ZeroInt()) {
		return sdk.NewError(DefaultCodespace, CodeInvalidCoinAmount, "invalid asset amount %v", msg.Asset.Amount)
	}
	return nil
}

type MsgLockAsset struct {
	Address      sdk.AccAddress `json:"address"`
	Asset        sdk.Coin       `json:"asset"`
	LockDuration time.Duration  `json:"lock_duration"`
}

func NewMsgLockAsset(address sdk.AccAddress, asset sdk.Coin, duration time.Duration) MsgLockAsset {
	return MsgLockAsset{
		Address:      address,
		Asset:        asset,
		LockDuration: duration,
	}
}

func (msg MsgLockAsset) Route() string { return RouterKey }
func (msg MsgLockAsset) Type() string  { return "lockAsset" }
func (msg MsgLockAsset) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Address}
}
func (msg MsgLockAsset) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}
func (msg MsgLockAsset) ValidateBasic() sdk.Error {
	if msg.Address.Empty() {
		return sdk.NewError(DefaultCodespace, CodeInvalidAddress, "nil address")
	}

	if msg.Asset.Amount.LTE(sdk.ZeroInt()) {
		return sdk.NewError(DefaultCodespace, CodeInvalidCoinAmount, "invalid asset amount %v", msg.Asset.Amount)
	}

	if int64(msg.LockDuration) < int64(time.Second) {
		return sdk.NewError(DefaultParamspace, CodeInvalidLockTime, "invalid lock time %v", msg.LockDuration)
	}

	return nil
}

type MsgUnLockAsset struct {
	Address sdk.AccAddress `json:"address"`
	Symbol  string         `json:"symbol"`
}

func NewMsgUnLockAsset(address sdk.AccAddress, symbol string) MsgUnLockAsset {
	return MsgUnLockAsset{
		Address: address,
		Symbol:  symbol,
	}
}

func (msg MsgUnLockAsset) Route() string { return RouterKey }
func (msg MsgUnLockAsset) Type() string  { return "unLockAsset" }
func (msg MsgUnLockAsset) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Address}
}
func (msg MsgUnLockAsset) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}
func (msg MsgUnLockAsset) ValidateBasic() sdk.Error {
	if msg.Address.Empty() {
		return sdk.NewError(DefaultCodespace, CodeInvalidAddress, "nil address")
	}

	return nil
}

type MsgDestroyAsset struct {
	Address sdk.AccAddress `json:"address"`
	Asset   sdk.Coin       `json:"asset"`
}

func NewMsgDestroyAsset(address sdk.AccAddress, asset sdk.Coin) MsgDestroyAsset {
	return MsgDestroyAsset{
		Address: address,
		Asset:   asset,
	}
}

func (msg MsgDestroyAsset) Route() string { return RouterKey }
func (msg MsgDestroyAsset) Type() string  { return "destroyAsset" }
func (msg MsgDestroyAsset) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Address}
}
func (msg MsgDestroyAsset) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}
func (msg MsgDestroyAsset) ValidateBasic() sdk.Error {
	if msg.Address.Empty() {
		return sdk.NewError(DefaultCodespace, CodeInvalidAddress, "nil address")
	}

	if msg.Asset.Amount.LTE(sdk.ZeroInt()) {
		return sdk.NewError(DefaultCodespace, CodeInvalidCoinAmount, "invalid asset amount %v", msg.Asset.Amount)
	}
	return nil
}

type MsgRuinAsset struct {
	Address sdk.AccAddress `json:"address"`
	Symbol  string         `json:"symbol"`
}

func NewMsgRuinAsset(address sdk.AccAddress, symbol string) MsgRuinAsset {
	return MsgRuinAsset{
		Address: address,
		Symbol:  symbol,
	}
}

func (msg MsgRuinAsset) Route() string { return RouterKey }
func (msg MsgRuinAsset) Type() string  { return "ruinAsset" }
func (msg MsgRuinAsset) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Address}
}
func (msg MsgRuinAsset) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}
func (msg MsgRuinAsset) ValidateBasic() sdk.Error {
	if msg.Address.Empty() {
		return sdk.NewError(DefaultCodespace, CodeInvalidAddress, "nil address")
	}
	return nil
}