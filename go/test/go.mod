module github.com/qqqzhch/amino-js/go/test

go 1.12

require (
	github.com/cosmos/amino-js/go/src v0.0.0
	github.com/cosmos/cosmos-sdk v0.35.0
	github.com/tendermint/go-amino v0.15.0
	github.com/tendermint/tendermint v0.31.7
)

replace github.com/cosmos/amino-js/go/src => ../src

replace github.com/cosmos/amino-js/go/lib => ../lib
