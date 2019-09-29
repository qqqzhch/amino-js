package lib

import (
	sdk "github.com/cosmos/amino-js/go/lib/cosmos/cosmos-sdk/types"
	"github.com/cosmos/amino-js/go/lib/cosmos/cosmos-sdk/x/params"
	
)


// Default parameter namespace
const (
	DefaultParamspace = ModuleName
)

// nolint - Keys for parameter access
var (
	DefaultAssetPrefix    = "u"
	DefaultPledgeCost     = sdk.NewInt(1000000000000)
	DefaultMaxNameLength  = sdk.NewInt(32)
	DefaultMaxSupply      = sdk.NewInt(90000000000000000)
	DefaultIssueCostRate  = sdk.NewDecWithPrec(10, 2)
	DefaultExcludeSymbols = []string{"lamb", "tbb", "1amb"}
)

var (
	KeyAssetPrefix    = []byte("AssetPrefix")
	KeyPledgeCost     = []byte("PledgeCost")
	KeyMaxNameLength  = []byte("MaxNameLength")
	KeyMaxSupply      = []byte("MaxSupply")
	KeyIssueCostRate  = []byte("IssueCostRate")
	KeyExcludeSymbols = []byte("ExcludeSymbols")
)

var _ params.ParamSet = (*Params)(nil)

// Params defines the parameters for the asset module.
type Params struct {
	AssetPrefix    string   `json:"asset_prefix"`
	PledgeCost     sdk.Int  `json:"pledge_cost"`
	MaxNameLength  sdk.Int  `json:"max_name_length"`
	MaxSupply      sdk.Int  `json:"max_supply"`
	IssueCostRate  sdk.Dec  `json:"issue_cost_rate"`
	ExcludeSymbols []string `json:"exclude_symbols"`
}

func NewParams(assetPrefix string, pledgeCost sdk.Int,
	maxNameLen sdk.Int, maxSupply sdk.Int, issueCostRate sdk.Dec, excludeSymbols []string) Params {
	return Params{
		AssetPrefix:    assetPrefix,
		PledgeCost:     pledgeCost,
		MaxNameLength:  maxNameLen,
		MaxSupply:      maxSupply,
		IssueCostRate:  issueCostRate,
		ExcludeSymbols: excludeSymbols,
	}
}

// Implements params.ParamSet
func (p *Params) ParamSetPairs() params.ParamSetPairs {
	return params.ParamSetPairs{
		{KeyAssetPrefix, &p.AssetPrefix},
		{KeyPledgeCost, &p.PledgeCost},
		{KeyMaxNameLength, &p.MaxNameLength},
		{KeyMaxSupply, &p.MaxSupply},
		{KeyIssueCostRate, &p.IssueCostRate},
		{KeyExcludeSymbols, &p.ExcludeSymbols},
	}
}

func DefaultParams() Params {
	return NewParams(DefaultAssetPrefix, DefaultPledgeCost,
		DefaultMaxNameLength, DefaultMaxSupply, DefaultIssueCostRate, DefaultExcludeSymbols)
}

func (p Params) String() string {
	return fmt.Sprintf(`Params:
  asset_prefix:    %s
  pledge_cost:     %s
  max_name_length: %s
  max_supply:      %s
  issue_cost_rate: %s
  exclude_symbols: %s`,
		p.AssetPrefix, p.PledgeCost,
		p.MaxNameLength, p.MaxSupply, p.IssueCostRate, p.ExcludeSymbols)
}