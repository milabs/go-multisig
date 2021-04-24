module github.com/milabs/go-multisig

go 1.15

require (
	github.com/btcsuite/btcd v0.21.0-beta
	github.com/btcsuite/btcutil v1.0.2
	github.com/milabs/go-multisig/address v0.0.0-00010101000000-000000000000
)

replace github.com/milabs/go-multisig/address => ./address
