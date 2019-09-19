package lib

import (
	"bytes"
	"fmt"
	"time"

	
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
