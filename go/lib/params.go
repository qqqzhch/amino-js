package lib

import (
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
