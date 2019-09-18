package lib

import (
	"bytes"
	"fmt"
	"time"

	"github.com/LambdaIM/lambda/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	cst "github.com/cosmos/cosmos-sdk/x/staking/types"
	"strconv"
)

type ValidatorType int

func (vt ValidatorType) Marshal() ([]byte, error) {
	return []byte(strconv.Itoa(int(vt))), nil
}

func (vt *ValidatorType) Unmarshal(data []byte) error {
	t, err := strconv.Atoi(string(data))
	*vt = ValidatorType(t)
	return err
}

func (vt ValidatorType) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Itoa(int(vt))), nil
}

func (vt *ValidatorType) UnmarshalJSON(data []byte) error {
	t, err := strconv.Atoi(string(data))
	*vt = ValidatorType(t)
	return err
}

const (
	PartnerValidatorType ValidatorType = iota
	ConsensusValidatorType
)

var (
	DefaultConsensusValMinSelfDelegation = sdk.NewInt(666666666)
	DefaultPartnerValMinSelfDelegation   = sdk.NewInt(33333333)
	DefaultConsensusValMinDelegation     = sdk.NewInt(1e6)
	DefaultPartnerValMinDelegation       = sdk.NewInt(1e5)

	DefaultConsensusValSelfDelegation = sdk.NewCoin(types.DefaultBondDenom, DefaultConsensusValMinSelfDelegation)
	DefaultPartnerValSelfDelegation   = sdk.NewCoin(types.DefaultBondDenom, DefaultPartnerValMinSelfDelegation)

	DefaultConsensusValidatorFixedCommissionRate = sdk.NewDecWithPrec(25, 2)
	DefaultPartnerValidatorFixedCommissionRate   = sdk.NewDecWithPrec(25, 2)
	DefaultCommissionMaxChangeRate               = sdk.NewDecWithPrec(1, 2)
)

// nolint - Keys for parameter access
var (
	KeyConsensusValidatorMinSelfDelegation   = []byte("ConsensusValidatorMinSelfDelegation")
	KeyPartnerValidatorMinSelfDelegation     = []byte("PartnerValidatorMinSelfDelegation")
	KeyConsensusValidatorMinDelegation       = []byte("ConsensusValidatorMinDelegation")
	KeyPartnerValidatorMinDelegation         = []byte("PartnerValidatorMinDelegation")
	KeyConsensusValidatorFixedCommissionRate = []byte("ConsensusValidatorFixedCommissionRate")
	KeyPartnerValidatorFixedCommissionRate   = []byte("PartnerValidatorFixedCommissionRate")

	KeySupplyDenom = []byte("SupplyDenom")
)

var _ params.ParamSet = (*Params)(nil)

//ParamsOutPut print out all parameters of staking module
type ParamsOutPut struct {
	UnbondingTime                         time.Duration `json:"unbonding_time"`
	MaxValidators                         uint16        `json:"max_validators"`
	MaxEntries                            uint16        `json:"max_entries"`
	BondDenom                             string        `json:"bond_denom"`
	SupplyDenom                           string        `json:"supply_denom"`
	ConsensusValidatorMinSelfDelegation   sdk.Int       `json:"consensus_validator_min_self_delegation"`
	PartnerValidatorMinSelfDelegation     sdk.Int       `json:"partner_validator_min_self_delegation"`
	ConsensusValidatorMinDelegation       sdk.Int       `json:"consensus_validator_min_delegation"`
	PartnerValidatorMinDelegation         sdk.Int       `json:"partner_validator_min_delegation"`
	ConsensusValidatorFixedCommissionRate sdk.Dec       `json:"consensus_validator_fixed_commission_rate"`
	PartnerValidatorFixedCommissionRate   sdk.Dec       `json:"partner_validator_fixed_commission_rate"`
}

// Params defines the high level settings for lambda staking
type Params struct {
	SupplyDenom                           string  `json:"supply_denom"`
	ConsensusValidatorMinSelfDelegation   sdk.Int `json:"consensus_validator_min_self_delegation"`
	PartnerValidatorMinSelfDelegation     sdk.Int `json:"partner_validator_min_self_delegation"`
	ConsensusValidatorMinDelegation       sdk.Int `json:"consensus_validator_min_delegation"`
	PartnerValidatorMinDelegation         sdk.Int `json:"partner_validator_min_delegation"`
	ConsensusValidatorFixedCommissionRate sdk.Dec `json:"consensus_validator_fixed_commission_rate"`
	PartnerValidatorFixedCommissionRate   sdk.Dec `json:"partner_validator_fixed_commission_rate"`
}

func NewParams(consensusValidatorMinSelfDelegation, partnerValidatorMinSelfDelegation,
	consensusValidatorMinDelegation, partnerValidatorMinDelegation sdk.Int,
	consensusValidatorFixedCommissionRate, partnerValidatorFixedCommissionRate sdk.Dec, supplyDemon string,
) Params {

	return Params{
		SupplyDenom:                           supplyDemon,
		ConsensusValidatorMinSelfDelegation:   consensusValidatorMinSelfDelegation,
		PartnerValidatorMinSelfDelegation:     partnerValidatorMinSelfDelegation,
		ConsensusValidatorMinDelegation:       consensusValidatorMinDelegation,
		PartnerValidatorMinDelegation:         partnerValidatorMinDelegation,
		ConsensusValidatorFixedCommissionRate: consensusValidatorFixedCommissionRate,
		PartnerValidatorFixedCommissionRate:   partnerValidatorFixedCommissionRate,
	}
}

// Implements params.ParamSet
func (p *Params) ParamSetPairs() params.ParamSetPairs {
	return params.ParamSetPairs{
		{KeySupplyDenom, &p.SupplyDenom},
		{KeyConsensusValidatorMinSelfDelegation, &p.ConsensusValidatorMinSelfDelegation},
		{KeyPartnerValidatorMinSelfDelegation, &p.PartnerValidatorMinSelfDelegation},
		{KeyConsensusValidatorMinDelegation, &p.ConsensusValidatorMinDelegation},
		{KeyPartnerValidatorMinDelegation, &p.PartnerValidatorMinDelegation},
		{KeyConsensusValidatorFixedCommissionRate, &p.ConsensusValidatorFixedCommissionRate},
		{KeyPartnerValidatorFixedCommissionRate, &p.PartnerValidatorFixedCommissionRate},
	}
}

// Equal returns a boolean determining if two Param types are identical.
// TODO: This is slower than comparing struct fields directly
func (p Params) Equal(p2 Params) bool {
	bz1 := cst.MsgCdc.MustMarshalBinaryLengthPrefixed(&p)
	bz2 := cst.MsgCdc.MustMarshalBinaryLengthPrefixed(&p2)
	return bytes.Equal(bz1, bz2)
}

// String returns a human readable string representation of the parameters.
func (p Params) String() string {
	return fmt.Sprintf(`Params:
  Partner validator min-self-delegation:   %s
  Consensus validator min-self-delegation: %s
  Partner validator min-delegation:   	   %s
  Consensus validator min-delegation:      %s
  Consensus validator commission rate:     %s
  Partner validator commission rate:       %s
  Supply Coin Denom:                       %s`,
		p.PartnerValidatorMinSelfDelegation, p.ConsensusValidatorMinSelfDelegation,
		p.PartnerValidatorMinDelegation, p.ConsensusValidatorMinDelegation,
		p.ConsensusValidatorFixedCommissionRate, p.PartnerValidatorFixedCommissionRate, p.SupplyDenom)
}

// validate a set of params
func (p Params) Validate() error {
	if p.SupplyDenom == "" {
		return fmt.Errorf("staking parameter SupplyDenom can't be an empty string")
	}
	return nil
}

// DefaultParams returns a default set of parameters.
func DefaultParams() Params {
	return NewParams(DefaultConsensusValMinSelfDelegation, DefaultPartnerValMinSelfDelegation,
		DefaultConsensusValMinDelegation, DefaultPartnerValMinDelegation,
		DefaultConsensusValidatorFixedCommissionRate, DefaultPartnerValidatorFixedCommissionRate, types.DefaultSupplyDenom)
}
